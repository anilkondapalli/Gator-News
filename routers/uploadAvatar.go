package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/*UploadAvatar is a service to upload avatar to the database*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading the image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying the image! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = bd.ModifyRecord(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error inserting the image to the Database! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
