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

func RegisterUser(ctx context.Context, pool *pgxpool.Pool, user *models.UserInfo) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		return nil
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

	return err
}

func LoginUser(ctx context.Context, pool *pgxpool.Pool, user *models.UserInfo) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		return err
	}
	defer conn.Release()

	var found bool
	err = conn.QueryRow(ctx, "SELECT TRUE FROM users where email_hash = $1 and pwd_hash = $2", user.EmailHash, user.PasswordHash).Scan(&found)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
		return err
	}

	return nil
}
