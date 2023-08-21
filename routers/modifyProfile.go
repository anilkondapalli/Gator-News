package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/*modifies the user profile*/
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data Incorrection "+err.Error(), 400)
		return
	}
	var status bool

	status, err = bd.ModifyRecord(t, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry, try again"+err.Error(), 400)
		return
	}
	if status == false{
		http.Error(w, "Failed to modify user record", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
