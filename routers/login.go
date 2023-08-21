package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/jwt"
	"github.com/SaiKumarMalve/Gator-News/models"
)

/* Login function is defined to perform the login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Username and/or password invalid"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The user Email is required ", 400)
		return

	}
	document, exists := bd.TryLogin(t.Email, t.Password)
	if exists == false {
		http.Error(w, "Username and/or password invalid", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occured while trying to generate the corresponding token"+err.Error(), 400)
		return
	}
	resp := models.AnswerLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
