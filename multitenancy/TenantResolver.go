package multitenancy

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func TenantResolver(r *http.Request) *sql.DB {
	tenant := mux.Vars(r)["tenant"]
	return GetDatasource(tenant)
}
