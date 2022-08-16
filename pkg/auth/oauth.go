package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/znbang/eventmap/internal/env"
	"github.com/znbang/eventmap/pkg/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

func getProviders() map[string]oauth2.Config {
	return map[string]oauth2.Config{
		"google": {
			ClientID:     env.Get(env.GoogleClientID),
			ClientSecret: env.Get(env.GoogleClientSecret),
			Endpoint:     google.Endpoint,
			RedirectURL:  env.Get(env.GoogleRedirectURI),
			Scopes:       []string{"openid", "profile", "email"},
		},
		"facebook": {
			ClientID:     env.Get(env.FacebookClientID),
			ClientSecret: env.Get(env.FacebookClientSecret),
			Endpoint:     facebook.Endpoint,
			RedirectURL:  env.Get(env.FacebookRedirectURI),
			Scopes:       []string{"public_profile", "email"},
		},
		"github": {
			ClientID:     env.Get(env.GithubClientID),
			ClientSecret: env.Get(env.GithubClientSecret),
			Endpoint:     github.Endpoint,
			RedirectURL:  env.Get(env.GithubRedirectURI),
			Scopes:       []string{"read:user"},
		},
		"line": {
			ClientID:     env.Get(env.LineClientID),
			ClientSecret: env.Get(env.LineClientSecret),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
				TokenURL: "https://api.line.me/oauth2/v2.1/token",
			},
			RedirectURL: env.Get(env.LineRedirectURI),
			Scopes:      []string{"profile"},
		},
	}
}

func getUserInfoEndpoints(provider string) string {
	endpoints := map[string]string{
		"google":   env.Get(env.GoogleUserInfoURL),
		"facebook": env.Get(env.FacebookUserInfoURL),
		"github":   env.Get(env.GithubUserInfoURL),
		"line":     env.Get(env.LineUserInfoURL),
	}
	return endpoints[provider]
}

func GetAuthUrls() map[string]string {
	state := uuid.RandomID()
	model := map[string]string{}
	for provider, config := range getProviders() {
		model[provider] = config.AuthCodeURL(state)
	}
	return model
}

func getUserInfo(userInfo *UserInfo, provider string, code string) error {
	config := getProviders()[provider]

	ctx := context.Background()
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	client := config.Client(ctx, token)
	resp, err := client.Get(getUserInfoEndpoints(provider))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch provider {
	case "google":
		*userInfo = &GoogleUserInfo{}
	case "github":
		*userInfo = &GithubUserInfo{}
	case "facebook":
		*userInfo = &FacebookUserInfo{}
	case "line":
		*userInfo = &LineUserInfo{}
	default:
		return fmt.Errorf("auth: unsupported provider: %v", provider)
	}

	return json.Unmarshal(data, userInfo)
}
