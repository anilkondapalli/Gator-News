package routers

import (
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/* HighRelationship registers the relationship between users*/
func HighRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "the ID parameter is mandatory", http.StatusBadRequest)
		return
	}
	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	status, err := bd.InsertRelationship(t)
	if err != nil {
		http.Error(w, "An error occured while inserting Relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Failed to insert the relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
