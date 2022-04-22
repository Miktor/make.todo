package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Miktor/make.todo/back/cmd/auth/database"
	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Env struct {
	pool  *pgxpool.Pool
	store *sessions.CookieStore
}

func main() {
	pool, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	env := &Env{pool: pool, store: sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SALT")))}
	env.store.Options.Secure = true
	env.store.Options.HttpOnly = true

	http.HandleFunc("/register", env.registerHandler)
	http.HandleFunc("/login", env.loginHandler)
	http.HandleFunc("/refresh-token", env.handler)
	log.Print("Starting...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (env *Env) registerHandler(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo := models.UserInfo{EmailHash: request.EmailHash, PasswordHash: request.PasswordHash}
	err = database.RegisterUser(context.Background(), env.pool, &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	session, _ := env.store.Get(r, "session")
	session.Values["loggedin"] = "true"
	session.Values["email_hash"] = userInfo.EmailHash
	session.Save(r, w)

}

func (env *Env) loginHandler(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo := models.UserInfo{EmailHash: request.EmailHash, PasswordHash: request.PasswordHash}
	err = database.LoginUser(context.Background(), env.pool, &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, _ := env.store.Get(r, "session")
	session.Values["loggedin"] = "true"
	session.Values["email_hash"] = userInfo.EmailHash
	session.Save(r, w)
}

func (env *Env) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
