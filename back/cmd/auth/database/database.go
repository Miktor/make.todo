package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDb() (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}

func RegisterUser(ctx context.Context, pool *pgxpool.Pool, user *models.UserInfo) (*models.RegisterResponse, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "INSERT INTO users (email_hash, pwd_hash) VALUES ($1, $2)", user.EmailHash, user.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	return &models.RegisterResponse{Token: "asdasd"}, err
}
