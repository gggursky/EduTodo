package store

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/Konstantin299/EduTodo.git/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

//go:embed migrations
var migrations embed.FS

type PGStore struct {
	log *logrus.Entry
	dsn string
	db  *pgxpool.Pool
}

func New(ctx context.Context, log *logrus.Logger, dsn string) (*PGStore, error) {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New(ctx, dsn): %w", err)
	}
	return &PGStore{
		log: log.WithField("module", "store"),
		dsn: dsn,
		db:  conn,
	}, nil

}

func (db *PGStore) Migrate(direction migrate.MigrationDirection) error {
	conn, err := sql.Open("pgx", db.dsn)
	if err != nil {
		return fmt.Errorf("sql.Open(\"pgx\", db.dsn): %w", err)
	}

	defer func() {
		if err = conn.Close(); err != nil {
			db.log.Error("err closing migration connection")
		}
	}()

	assetDir := func() func(string) ([]string, error) {
		return func(path string) ([]string, error) {
			dirEntry, err := migrations.ReadDir(path)
			if err != nil {
				return nil, fmt.Errorf("migrations.ReadDir(path): %w", err)
			}

			entries := make([]string, 0)

			for _, e := range dirEntry {
				entries = append(entries, e.Name())
			}

			return entries, nil
		}
	}()

	asset := migrate.AssetMigrationSource{
		Asset:    migrations.ReadFile,
		AssetDir: assetDir,
		Dir:      "migrations",
	}

	_, err = migrate.Exec(conn, "postgres", asset, direction)
	if err != nil {
		return fmt.Errorf("migrate.Exec(conn, \"postgres\", asset, direction): %w", err)
	}

	return nil
}

func (db *PGStore) GetPool() *pgxpool.Pool {
	return db.db
}

func (db *PGStore) GetThemasByCourse(code string) ([]models.Thema, error) {
	return nil, nil
}
