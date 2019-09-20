package multitenancy

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func TenantResolver(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := mux.Vars(r)["tenant"]
		ctx := context.WithValue(r.Context(), "tenant", tenant)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
