package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Miktor/make.todo/back/cmd/auth/controller"
	"github.com/Miktor/make.todo/back/cmd/auth/database"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/etherlabsio/healthcheck/v2/checkers"
	"github.com/gorilla/sessions"
)

func main() {
	db, err := database.InitPG()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	session_store := sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SALT")))
	session_store.Options.Secure = true
	session_store.Options.HttpOnly = true

	env, err := controller.AuthController(db, session_store)

	http.HandleFunc("/register", env.RegisterHandler)
	http.HandleFunc("/login", env.LoginHandler)
	http.HandleFunc("/refresh-token", env.RefreshHandler)

	http.HandleFunc("/health", healthcheck.HandlerFunc(

		// WithTimeout allows you to set a max overall timeout.
		healthcheck.WithTimeout(5*time.Second),

		healthcheck.WithChecker(
			"database", healthcheck.CheckerFunc(
				func(ctx context.Context) error {
					return db.Ping(ctx)
				},
			),
		),

		// Observers do not fail the status in case of error.
		healthcheck.WithObserver(
			"diskspace", checkers.DiskSpace("/var/log", 90),
		),
	))

	log.Print("Starting...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
