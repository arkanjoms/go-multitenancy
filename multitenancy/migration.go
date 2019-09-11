package multitenancy

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go-multitenancy/util"
)

func migration(datasource *sql.DB, sourceUrl string) {
	driver, err := postgres.WithInstance(datasource, &postgres.Config{})
	util.LogError(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+sourceUrl,
		"postgres", driver,
	)
	_ = m.Steps(2)
}
