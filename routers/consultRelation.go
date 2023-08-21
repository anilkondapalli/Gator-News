package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/*ConsultRelation checks the relationship between 2 users*/
func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	var response models.RequestConsultRelation

	status, err := bd.ConsultRelation(t)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
