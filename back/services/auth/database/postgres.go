package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	Close()
}

type DB struct {
	repository Repository
}

func InitPG() (*DB, error) {
	repository, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &DB{repository: repository}, nil
}

func (db *DB) RegisterUser(ctx context.Context, user *models.UserInfo) error {
	_, err := db.repository.Exec(ctx, "INSERT INTO users (email_hash, pwd_hash) VALUES ($1, $2)", user.EmailHash, user.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	return err
}

func (db *DB) LoginUser(ctx context.Context, user *models.UserInfo) error {
	var found bool
	err := db.repository.QueryRow(ctx, "SELECT TRUE FROM users where email_hash = $1 and pwd_hash = $2", user.EmailHash, user.PasswordHash).Scan(&found)
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

func (db *DB) Close() {
	db.repository.Close()
}

var _ AuthDB = (*DB)(nil)
