package tm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"

	"github.com/xkamail/huberlink-platform/pkg/rand"
)

func New(t *testing.T, migrationPath string) (*TestMigration, error) {
	migratePath, err := filepath.Abs(migrationPath)
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DB_TEST_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_TEST_PASS")
	if password == "" {
		password = "postgres"
	}
	host := os.Getenv("DB_TEST_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_TEST_PORT")
	if port == "" {
		port = "5432"
	}
	// create random database name
	dbName, err := rand.String(20)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf(`user=%s password=%s host=%s port=%s dbname=%s sslmode=disable`,
		user,
		password,
		host,
		port,
		dbName,
	)

	// setup database if not exists
	masterDB, err := sql.Open("postgres", fmt.Sprintf(`user=%s password=%s host=%s port=%s dbname=%s sslmode=disable`,
		user,
		password,
		host,
		port,
		"postgres",
	))
	if err != nil {
		return nil, err
	}
	if err = createDatabase(masterDB, dbName); err != nil {
		return nil, err
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: postgres.DefaultMigrationsTable,
	})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:"+migrationPath,
		dbName,
		driver,
	)
	if err != nil {
		return nil, err
	}

	return &TestMigration{
		ctx:           ctx,
		db:            db,
		pool:          pool,
		masterDB:      masterDB,
		migrationPath: migratePath,
		m:             m,
	}, nil
}

type TestMigration struct {
	ctx           context.Context
	masterDB      *sql.DB
	db            *sql.DB
	pool          *pgxpool.Pool
	migrationPath string
	m             *migrate.Migrate
}

func (t TestMigration) CreateTable() error {
	err := t.m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}
func (t TestMigration) PgxPool() *pgxpool.Pool {
	return t.pool
}

func createDatabase(db *sql.DB, dbName string) error {
	var exists bool
	err := db.QueryRow(`
			select exists(select datname from pg_database where datname = $1)`,
		dbName,
	).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// language=TEXT
	_, err = db.Exec("create database " + pq.QuoteIdentifier(dbName))
	return err
}
