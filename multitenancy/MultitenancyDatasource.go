package multitenancy

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"go-multitenancy/util"
	"log"
	"os"
	"strings"
)

var datasources map[string]*sql.DB
var sourceUrl string

func init() {
	_ = gotenv.Load()
	datasources = map[string]*sql.DB{}

	tenants := strings.Split(os.Getenv("TENANTS"), ",")

	for _, t := range tenants {
		createDatasource(t)
	}
}

func loadProperties(tenant string) string {
	t := strings.ToUpper(tenant)
	dbhost := os.Getenv(strings.Join([]string{t, "_DB_HOST"}, ""))
	dbport := os.Getenv(strings.Join([]string{t, "_DB_PORT"}, ""))
	dbname := os.Getenv(strings.Join([]string{t, "_DB_NAME"}, ""))
	dbusername := os.Getenv(strings.Join([]string{t, "_DB_USERNAME"}, ""))
	dbpassword := os.Getenv(strings.Join([]string{t, "_DB_PASSWORD"}, ""))

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbusername, dbpassword, dbname)
}

func createDatasource(tenant string) *sql.DB {
	log.Printf("Creating datasource for tenant [%s]", tenant)
	config := loadProperties(tenant)
	ds, err := sql.Open("postgres", config)
	util.LogError(err)
	err = ds.Ping()
	util.LogError(err)

	sourceUrl = os.Getenv("MIGRATION_SOURCE_URL")
	migration(ds, sourceUrl)

	datasources[tenant] = ds
	return ds
}

func GetDatasource(tenant string) *sql.DB {
	ds := datasources[tenant]
	if ds == nil {
		util.LogError(errors.New(fmt.Sprintf("Tenant not found [%s]", tenant)))
	}
	return ds
}
