package middlew

import (
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
)

//Checks if the database connection is lost
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Database connection is lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
