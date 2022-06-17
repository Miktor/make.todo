package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Miktor/make.todo/back/cmd/auth/database"
	"github.com/Miktor/make.todo/back/cmd/auth/mocks"
	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/gorilla/sessions"
)

type Request struct {
	EmailHash    string
	PasswordHash string
}

func Test_auth_RegisterHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		body       Request
		want       string
		dbCalled   bool
		statusCode int
	}{
		{
			name:       "simple register",
			method:     http.MethodGet,
			body:       Request{EmailHash: "", PasswordHash: "123123"},
			want:       `Invalid email`,
			dbCalled:   false,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "simple register",
			method:     http.MethodGet,
			body:       Request{EmailHash: "some@mail", PasswordHash: ""},
			want:       `Invalid password`,
			dbCalled:   false,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "simple register",
			method:     http.MethodGet,
			body:       Request{EmailHash: "some@mail", PasswordHash: "123123"},
			want:       `Invalid password`,
			dbCalled:   true,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db := mocks.NewAuthDB(t)
			store := mocks.NewSessionStore(t)
			auth := &auth{
				db:    db,
				store: store,
			}

			request := httptest.NewRequest(tc.method, "/register", bytes.NewReader(b))
			responseRecorder := httptest.NewRecorder()

			if tc.dbCalled == true {
				db.On("RegisterUser", context.Background(), &models.UserInfo{EmailHash: tc.body.EmailHash, PasswordHash: tc.body.PasswordHash}).Return(nil)

				store.On("Get", request, "session").Return()
			}

			b, err := json.Marshal(tc.body)
			if err != nil {
				fmt.Println(err)
				return
			}

			auth.RegisterHandler(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func Test_auth_LoginHandler(t *testing.T) {
	type fields struct {
		db    database.AuthDB
		store *sessions.CookieStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := &auth{
				db:    tt.fields.db,
				store: tt.fields.store,
			}
			auth.LoginHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_auth_RefreshHandler(t *testing.T) {
	type fields struct {
		db    database.AuthDB
		store *sessions.CookieStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := &auth{
				db:    tt.fields.db,
				store: tt.fields.store,
			}
			auth.RefreshHandler(tt.args.w, tt.args.r)
		})
	}
}
