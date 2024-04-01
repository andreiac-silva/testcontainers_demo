package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresDatabase struct {
	instance *postgres.PostgresContainer
}

func NewPostgresDatabase(t *testing.T, ctx context.Context) *PostgresDatabase {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:12"),
		postgres.WithDatabase("test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)
	return &PostgresDatabase{
		instance: pgContainer,
	}
}

func (db *PostgresDatabase) DSN(t *testing.T, ctx context.Context) string {
	dsn, err := db.instance.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)
	return dsn
}

func (db *PostgresDatabase) Close(t *testing.T, ctx context.Context) {
	require.NoError(t, db.instance.Terminate(ctx))
}
