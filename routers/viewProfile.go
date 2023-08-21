package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
)

/*ViewProfile allows to extract the values ​​of the profile*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find the record "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
