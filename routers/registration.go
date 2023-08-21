package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

// Registartion handles the user registartion and calls the bd API to insert the User profile to DB
func Registration(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in Data Recieved "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User Email is Required "+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password should be atleast 6 characters"+err.Error(), 400)
		return
	}
	_, found, _ := bd.CheckUserAlreadyExists(t.Email)
	if found == true {
		http.Error(w, "Already user exists with this username ", 400)
		return
	}

	_, status, err := bd.InsertRegistration(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to perform user registration"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed to insert user record ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
