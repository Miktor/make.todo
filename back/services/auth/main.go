package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Miktor/make.todo/back/cmd/auth/controller"
	"github.com/Miktor/make.todo/back/cmd/auth/database"
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
	log.Print("Starting...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
