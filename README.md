### Environment variables

#### Server
```
PORT=9000
COOKIE_NAME=EVENTMAP_SID
COOKIE_SECRET=
JWT_SECRET=
DATABASE_URL=postgres://postgres@localhost/eventmap?sslmode=disable
```

#### Google
```
GOOGLE_MAPS_API_KEY=
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_REDIRECT_URI=http://localhost:8080/login/oauth2/code/google
GOOGLE_USER_INFO_URL=https://www.googleapis.com/oauth2/v3/userinfo
```

#### Facebook
```
FACEBOOK_CLIENT_ID=
FACEBOOK_CLIENT_SECRET=
FACEBOOK_REDIRECT_URI=http://localhost:8080/login/oauth2/code/facebook
FACEBOOK_USER_INFO_URL=https://graph.facebook.com/v10.0/me?fields=id,name,picture{url},email
```

#### Github
```
GITHUB_CLIENT_ID=
GITHUB_CLIENT_SECRET=
GITHUB_REDIRECT_URI=http://localhost:8080/login/oauth2/code/github
GITHUB_USER_INFO_URL=https://api.github.com/user
```

#### Line
```
LINE_CLIENT_ID=
LINE_CLIENT_SECRET=
LINE_REDIRECT_URI=http://localhost:8080/login/oauth2/code/line
LINE_USER_INFO_URL=https://api.line.me/v2/profile
```

### Build
```
go build -trimpath -ldflags="-s -w"
```

### Test Crawlers
```
go test -v ./pkg/crawlerservice
```