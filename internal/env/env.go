package env

import (
	"log"
	"os"
)

func IsSet(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

func Get(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	log.Fatalln("env not set:", key)
	return ""
}

func check(key string) {
	if !IsSet(key) {
		log.Fatalln("env not set: ", key)
	}
}

func Verify() {
	check(Port)
	check(CookieName)
	check(CookieSecret)
	check(JWTSecret)
	check(DatabaseURL)

	check(GoogleClientID)
	check(GoogleClientSecret)
	check(GoogleRedirectURI)
	check(GoogleUserInfoURL)

	check(FacebookClientID)
	check(FacebookClientSecret)
	check(FacebookRedirectURI)
	check(FacebookUserInfoURL)

	check(GithubClientID)
	check(GithubClientSecret)
	check(GithubRedirectURI)
	check(GithubUserInfoURL)

	check(LineClientID)
	check(LineClientSecret)
	check(LineRedirectURI)
	check(LineUserInfoURL)
}
