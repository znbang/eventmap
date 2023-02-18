package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/znbang/eventmap/internal/env"
	"github.com/znbang/eventmap/pkg/login"
	"github.com/znbang/eventmap/pkg/userservice"
)

func WithUser(userService *userservice.UserService, loginService *login.Service, next http.Handler) http.HandlerFunc {
	jwtSecret := []byte(env.Get("JWT_SECRET"))

	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if bearerToken == "" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)

		var (
			userSession login.UserSession
			user        userservice.User
		)

		if err = loginService.FindSessionById(&userSession, claims.Subject); err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if err = userService.FindById(&user, userSession.UserID); err != nil {
			next.ServeHTTP(w, r)
			return
		} else {
			ctx := context.WithValue(r.Context(), sessionIdKey, userSession.ID)
			ctx = context.WithValue(ctx, currentUserKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
	}
}

func CurrentUser(ctx context.Context, user *userservice.User) error {
	if obj := ctx.Value(currentUserKey); obj != nil {
		*user = obj.(userservice.User)
		return nil
	}
	return errors.New("currentUser is not available")
}
