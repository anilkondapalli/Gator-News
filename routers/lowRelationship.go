package routers

import (
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/* LowRelationship performs the deletion of the relationship between users */
func LowRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	status, err := bd.DeleteRelationship(t)
	if err != nil {
		http.Error(w, "An error occured while deleting Relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Failed to delete the relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
