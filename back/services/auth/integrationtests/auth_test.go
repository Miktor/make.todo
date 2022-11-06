package dockertestpsql_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"github.com/Miktor/make.todo/back/cmd/auth/models"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	log "github.com/sirupsen/logrus"
)

type Db struct {
	connectionString string
	db               *sql.DB
	resource         *dockertest.Resource
}

// Error implements error.
func (*Db) Error() string {
	panic("unimplemented")
}

func GetAbsolutePath(relPath string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path.Join(pwd, relPath), nil
}

func startPsql() (*Db, error) {
	hostAndPort := "localhost:5432"
	databaseUrl := fmt.Sprintf("postgres://admin:admin@%s/todo?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	hostAndPort = "db:5432"
	databaseUrl = fmt.Sprintf("postgres://admin:admin@%s/todo?sslmode=disable", hostAndPort)
	database := &Db{
		connectionString: databaseUrl,
		db:               db,
	}

	return database, database.db.Ping()
}

func startAuthService() error {
	resp, err := http.Get("http://localhost:8000/health")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid return code")
	}

	return err
}

func TestMain(m *testing.M) {
	_, err := startPsql()
	if err != nil {
		log.Fatalf("Could not start psql: %s", err)
	}

	err = startAuthService()
	if err != nil {
		log.Fatalf("Could not start psql: %s", err)
	}

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func TestCreatePhoneNumber(t *testing.T) {
	cases := []struct {
		error       bool
		description string
	}{
		{
			description: "Should succeed with valid creation of a phone number",
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			request := &models.RegisterRequest{
				EmailHash:    randomString(10),
				PasswordHash: randomString(10),
			}

			requestBody, err := json.Marshal(request)
			if err != nil {
				t.Fatalf("Failed to marshal json '%s'", err)
			}

			resp, err := http.Post("http://localhost:8000/register", "application/json", bytes.NewReader(requestBody))
			if err != nil {
				t.Fatalf("Got error '%s'", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Got invalid status code '%d'", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				t.Fatalf("Got error '%s'", err)
			}

			t.Log(string(body))
		})
	}
}
