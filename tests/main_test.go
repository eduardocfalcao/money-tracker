package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/eduardocfalcao/money-tracker/database/queries"
	"github.com/golang-migrate/migrate/v4"
	migratePostgress "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // used by migrator used by migrator
	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var pgxConn *pgx.Conn

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbName := "test"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15"),
		//postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		//postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	ensureNil(err)

	pgxConn, err = pgx.Connect(ctx, connStr)
	ensureNil(err)
	defer pgxConn.Close(ctx)
	err = migrateDb(connStr)
	ensureNil(err)

	code := m.Run()

	os.Exit(code)
}

func ensureNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func migrateDb(connStr string) error {

	// get location of test
	_, fpath, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}
	pathToMigrationFiles := "file://" + filepath.Dir(fpath) + "/../database/migrations"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	driver, err := migratePostgress.WithInstance(db, &migratePostgress.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(pathToMigrationFiles, "test", driver)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("migration done")

	return nil
}

type testStage struct {
	Repository queries.QuerierTx
	pgxConn    *pgx.Conn
}

func createTestStage() *testStage {
	r := queries.New(pgxConn)

	return &testStage{
		Repository: r,
		pgxConn:    pgxConn,
	}
}

func (t *testStage) CleanUp(ctx context.Context) {
	tables := []string{
		"users",
		"raw_transactions",
	}

	for _, table := range tables {
		pgxConn.Exec(ctx, fmt.Sprintf("TRUNCAte %s", table))
	}
}
