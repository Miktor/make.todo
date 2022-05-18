package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Miktor/make.todo/back/cmd/auth/database"
	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/gorilla/sessions"
)

type SessionStore interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
}

type auth struct {
	db    database.AuthDB
	store SessionStore
}

type Auth interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	RefreshHandler(w http.ResponseWriter, r *http.Request)
}

func AuthController(db database.AuthDB, store SessionStore) (Auth, error) {
	auth := &auth{db: db, store: store}

	return auth, nil
}

func (auth *auth) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo := models.UserInfo{EmailHash: request.EmailHash, PasswordHash: request.PasswordHash}

	//TODO: add email validation
	if len(userInfo.EmailHash) < 3 {
		http.Error(w, "Invalid Email", http.StatusBadRequest)
		return
	}

	//TODO: add password validation
	if len(userInfo.PasswordHash) < 3 {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	err = auth.db.RegisterUser(context.Background(), &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	session, _ := auth.store.Get(r, "session")
	session.Values["loggedin"] = "true"
	session.Values["email_hash"] = userInfo.EmailHash
	session.Save(r, w)

}

func (auth *auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo := models.UserInfo{EmailHash: request.EmailHash, PasswordHash: request.PasswordHash}
	err = auth.db.LoginUser(context.Background(), &userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, _ := auth.store.Get(r, "session")
	session.Values["loggedin"] = "true"
	session.Values["email_hash"] = userInfo.EmailHash
	session.Save(r, w)
}

func (auth *auth) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
