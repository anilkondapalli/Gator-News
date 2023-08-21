package routers

import (
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
)

/* This removes the user's post */
func RemovePost(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	err := bd.DeletePost(ID, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the Post"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
