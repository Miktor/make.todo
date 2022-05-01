package database

import (
	"context"
	"testing"

	"github.com/Miktor/make.todo/back/cmd/auth/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pashagolub/pgxmock"
)

type PgxPoolMock interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
}

func TestDB_RegisterUser(t *testing.T) {
	mock_pool, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("NewPool error = %v", err)
	}

	db := &DB{
		repository: mock_pool,
	}

	email, pwd := "email", "pdw"
	mock_pool.ExpectExec("INSERT INTO users").WithArgs(email, pwd).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	if err = db.RegisterUser(context.Background(), &models.UserInfo{EmailHash: email, PasswordHash: pwd}); err != nil {
		t.Errorf("DB.RegisterUser() error = %v", err)
	}

	if err = mock_pool.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDB_LoginUser(t *testing.T) {
	mock_pool, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("NewPool error = %v", err)
	}

	db := &DB{
		repository: mock_pool,
	}

	email, pwd := "email", "pdw"
	mock_pool.ExpectQuery("SELECT TRUE FROM users").WithArgs(email, pwd).WillReturnRows(pgxmock.NewRows([]string{"true"}).AddRow(true))

	if err = db.LoginUser(context.Background(), &models.UserInfo{EmailHash: email, PasswordHash: pwd}); err != nil {
		t.Errorf("DB.RegisterUser() error = %v", err)
	}

	if err = mock_pool.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDB_Close(t *testing.T) {
	mock_pool, err := pgxmock.NewPool()
	if err != nil {
		t.Errorf("NewPool error = %v", err)
	}

	db := &DB{
		repository: mock_pool,
	}

	mock_pool.ExpectClose()

	db.Close()

	if err = mock_pool.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
