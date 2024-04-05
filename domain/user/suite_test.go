package user

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/andreiac-silva/testcontainers_demo/test/integration"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite
	db        *bun.DB
	container *integration.PostgresDatabase
}

func (s *IntegrationTestSuite) SetupSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.setupDatabase(ctx)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.container.Close(s.T(), ctx)
}

func (s *IntegrationTestSuite) setupDatabase(ctx context.Context) {
	s.container = integration.NewPostgresDatabase(s.T(), ctx)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(s.container.DSN(s.T(), ctx))))
	s.db = bun.NewDB(sqldb, pgdialect.New())
	s.migrate(ctx)
}

func (s *IntegrationTestSuite) migrate(ctx context.Context) {
	err := s.initDatabase(ctx)
	require.NoError(s.T(), err)

	migrations := &migrate.Migrations{}
	err = migrations.Discover(os.DirFS("../../migrations"))
	require.NoError(s.T(), err)

	migrator := migrate.NewMigrator(s.db, migrations)
	_, err = migrator.Migrate(ctx)
	require.NoError(s.T(), err)
}

func (s *IntegrationTestSuite) initDatabase(ctx context.Context) error {
	type hack struct {
		bun.BaseModel `bun:"table:bun_migrations"`
		*migrate.Migration
	}
	_, err := s.db.NewCreateTable().Model((*hack)(nil)).Table("bun_migrations").Exec(ctx)
	require.NoError(s.T(), err)
	return nil
}
