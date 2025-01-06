package unittesting

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/GogoManager/utils/migration"
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDatabase(migrationsPath string, t *testing.T) *sqlx.DB {
	ctx := context.Background()

	// 1. Start the postgres container and run any migrations on it
	container, err := postgres.RunContainer(
		ctx,
		testcontainers.WithImage("postgis/postgis:16-master"),
		postgres.WithDatabase("testing"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	dbURL, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	if err := migration.MigrateUp(dbURL, migrationsPath); err != nil {
		t.Fatalf("failed to migrate up db: %s", err)
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	statementutil.SetUp(db)

	t.Cleanup(statementutil.CleanUp)

	return db
}
