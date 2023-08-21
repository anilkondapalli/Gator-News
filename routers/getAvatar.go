package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/SaiKumarMalve/Gator-News/bd"
)

/*GetAvatar sends the avatar to HTTP*/
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must Provide an ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.ProfileLookup(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying the image", http.StatusBadRequest)
	}
}
