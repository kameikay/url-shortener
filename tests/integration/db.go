package integration

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var TestDB *sql.DB

func TestMainIntegrationDB(m *testing.M, migrationPath string) int {
	pool, err := dockertest.NewPool("")
	if err != nil {
		fmt.Printf("Could not create docker pool: %s", err)
		return 1
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	var cleanup func() error
	TestDB, cleanup, err = setupTestDB(pool)
	if err != nil {
		fmt.Printf("Could not setup test db: %s", err)
		return 1
	}

	defer func() {
		if err := cleanup(); err != nil {
			fmt.Printf("could not cleanup postgres container: %v\n", err)
		}
	}()

	err = runMigrations(migrationPath, TestDB)
	if err != nil {
		fmt.Printf("Could not run migrations: %s", err)
		return 1
	}

	return m.Run()
}

func setupTestDB(pool *dockertest.Pool) (*sql.DB, func() error, error) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2",
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=db",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return nil, nil, fmt.Errorf("could not create resource: %w", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://postgres:postgres@%s/db?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second

	var db *sql.DB
	err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return fmt.Errorf("could not connect to docker db: %w", err)
		}

		if err = db.Ping(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("could not connect to docker db: %w", err)
	}

	return db, func() error {
		return pool.Purge(resource)
	}, nil
}

func runMigrations(migrationPath string, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: "db",
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
