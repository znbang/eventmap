package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/znbang/eventmap/gen/auth/v1/authv1connect"
	"github.com/znbang/eventmap/gen/book/v1/bookv1connect"
	"github.com/znbang/eventmap/gen/event/v1/eventv1connect"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/internal/env"
	"github.com/znbang/eventmap/internal/gorm/repository"
	"github.com/znbang/eventmap/pkg/auth"
	"github.com/znbang/eventmap/pkg/bookservice"
	"github.com/znbang/eventmap/pkg/eventservice"
	"github.com/znbang/eventmap/pkg/login"
	"github.com/znbang/eventmap/pkg/modfs"
	"github.com/znbang/eventmap/pkg/userservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//go:embed frontend/dist/spa
var distFS embed.FS

// @title EventMap API
// @version 1.0
// @license.name Apache 2.0
func main() {
	env.Verify()

	// db, err := dbx.Open(env.Get(env.DatabaseURL))
	db, err := dbx.Open("eventmap.db")
	if err != nil {
		log.Fatal(err)
	}

	bookRepository := repository.NewBookRepository(db)
	bookJobRepository := repository.NewBookJobRepository(db)
	chapterRepository := repository.NewChapterRepository(db)
	eventRepository := repository.NewEventRepository(db)
	userRepository := repository.NewUserRepository(db)
	userSessionRepository := repository.NewUserSessionRepository(db)

	userService := userservice.NewUserService(userRepository)
	bookService := bookservice.NewBookService(bookRepository, bookJobRepository, chapterRepository)
	eventService := eventservice.NewEventService(eventRepository)
	loginService := login.NewService(userSessionRepository)
	authService := auth.NewService(userService)

	authServer := auth.NewAuthServer(loginService, userService)
	bookServer := bookservice.NewBookServer(bookService)
	eventServer := eventservice.NewEventServer(eventService)

	api := http.NewServeMux()
	api.Handle(bookv1connect.NewBookServiceHandler(bookServer))
	api.Handle(eventv1connect.NewEventServiceHandler(eventServer))
	api.Handle(authv1connect.NewAuthServiceHandler(authServer))

	mux := http.NewServeMux()
	mux.Handle("/api/", auth.WithUser(userService, loginService, http.StripPrefix("/api", api)))
	mux.Handle("/api/sync-book-status", bookservice.CreateSyncStatusHandler())
	mux.Handle("/api/download-book", bookservice.CreateDownloadBookHandler(bookService))
	mux.Handle("/login/oauth2/code/", auth.CreateOauthHandleFunc(authService, loginService))

	if spaFS, err := fs.Sub(distFS, "frontend/dist/spa"); err != nil {
		log.Fatal(err)
	} else {
		mux.Handle("/", http.FileServer(modfs.New(http.FS(spaFS), time.Now())))
	}

	go bookService.HandleBookJob()

	addr := ":" + env.Get(env.Port)
	if err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatal(err)
	}
}
