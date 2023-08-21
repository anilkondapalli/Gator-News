package middlew

import (
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/routers"
)

/*ValidJWT allows to validate the JWT that comes to us in the request*/
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in the Token! "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
