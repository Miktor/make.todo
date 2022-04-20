package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Miktor/make.todo/back/cmd/auth/database"
	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Env struct {
	pool *pgxpool.Pool
}

func main() {
	pool, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	env := &Env{pool: pool}

	http.HandleFunc("/register", env.registerHandler)
	http.HandleFunc("/login", env.handler)
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
	token, err := database.RegisterUser(context.Background(), env.pool, &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(token)
}

func (env *Env) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
