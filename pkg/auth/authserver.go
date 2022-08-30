package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt"
	v1 "github.com/znbang/eventmap/gen/auth/v1"
	"github.com/znbang/eventmap/internal/env"
	"github.com/znbang/eventmap/pkg/login"
	"github.com/znbang/eventmap/pkg/userservice"
)

const GetUserPath = "/api/auth.v1.AuthService/GetUser"

type AuthServer struct {
	loginService *login.Service
	userService *userservice.UserService
}

func NewAuthServer(loginService *login.Service, userService *userservice.UserService) *AuthServer {
	return &AuthServer{
		loginService: loginService,
		userService: userService,
	}
}

func (s *AuthServer) ListProvider(ctx context.Context, r *connect.Request[v1.ListProviderRequest]) (*connect.Response[v1.ListProviderResponse], error) {
	var providers = make([]*v1.Provider, 0)

	for k, v := range GetAuthUrls() {
		providers = append(providers, &v1.Provider{
			Name: k,
			Url:  v,
		})
	}

	return connect.NewResponse(&v1.ListProviderResponse{
		Providers: providers,
	}), nil
}

func (s *AuthServer) GetUser(ctx context.Context, r *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	var (
		sessionId   string
		userSession login.UserSession
		user        userservice.User
	)

	if err := getSessionIdFromCookie(r, &sessionId); err != nil {
		resp := connect.NewResponse(&v1.GetUserResponse{})
		deleteCookie(resp)
		return resp, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// load user session
	if err := s.loginService.FindSessionById(&userSession, sessionId); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// update user session
	if err := s.loginService.SaveSession(&userSession); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// get user
	if err := s.userService.FindById(&user, userSession.UserID); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	// generate token
	claims := jwt.StandardClaims{
		Subject:   sessionId,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(env.Get(env.JWTSecret)))
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	resp := connect.NewResponse(&v1.GetUserResponse{
		Email: user.Email,
		Name:  user.Name,
		Token: signedToken,
	})

	// update cookie
	addCookie(resp, env.Get(env.CookieName), []byte(env.Get(env.CookieSecret)), sessionId)

	return resp, nil
}

func (s *AuthServer) Logout(ctx context.Context, r *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	resp := connect.NewResponse(&v1.LogoutResponse{})
	deleteCookie(resp)
	if sessionId, ok := ctx.Value("sessionId").(string); ok {
		if err := s.loginService.DeleteSessionById(sessionId); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func unauthenticated(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = fmt.Fprintf(w, "auth: provider is required")
}

func CreateOauthHandleFunc(authService *Service, loginService *login.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/login/oauth2/code/"))
		if provider == "" {
			unauthenticated(w, fmt.Errorf("auth: provider is required"))
			return
		}

		query := r.URL.Query()
		code := strings.TrimSpace(query.Get("code"))
		if code == "" {
			unauthenticated(w, fmt.Errorf("auth: code is required"))
			return
		}

		var (
			user        userservice.User
			userSession login.UserSession
		)

		if err := authService.Login(&user, provider, code); err != nil {
			unauthenticated(w, fmt.Errorf("auth: login failed: %v", err))
			return
		}

		if err := loginService.CreateSession(&userSession, user.ID); err != nil {
			unauthenticated(w, fmt.Errorf("auth: create session failed: %v", err))
			return
		}

		cookie := http.Cookie{
			Name:     env.Get(env.CookieName),
			Value:    sign(userSession.ID, []byte(env.Get(env.CookieSecret))),
			Path:     "/api/auth.v1.AuthService/GetUser",
			MaxAge:   30 * 24 * 60 * 60,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func getSessionIdFromCookie[T any](r *connect.Request[T], sessionId *string) error {
	request := http.Request{Header: r.Header()}
	cookie, err := request.Cookie(env.Get(env.CookieName))
	if err != nil {
		return err
	}

	value := cookie.Value
	if value == "" {
		return http.ErrNoCookie
	}

	// verify value
	values := strings.Split(value, ".")
	if len(values) != 2 {
		return http.ErrNoCookie
	}

	*sessionId = values[0]
	sig := values[1]

	if !verify(*sessionId, sig, []byte(env.Get(env.CookieSecret))) {
		return http.ErrNoCookie
	}

	return nil
}

func addCookie[T any](w *connect.Response[T], cookieName string, cookieSecret []byte, sessionId string) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    sign(sessionId, cookieSecret),
		Path:     GetUserPath,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	w.Header().Set("Set-Cookie", cookie.String())
}

func deleteCookie[T any](w *connect.Response[T]) {
	cookie := http.Cookie{
		Name:   env.Get(env.CookieName),
		Path:   GetUserPath,
		MaxAge: -1,
	}
	w.Header().Set("Set-Cookie", cookie.String())
}

func sign(value string, cookieSecret []byte) string {
	h := hmac.New(sha256.New, cookieSecret)
	h.Write([]byte(value))
	sig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return value + "." + sig
}

func verify(value string, sig string, cookieSecret []byte) bool {
	h := hmac.New(sha256.New, cookieSecret)
	h.Write([]byte(value))
	return sig == base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
