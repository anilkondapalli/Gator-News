package bd

import (
	"github.com/SaiKumarMalve/Gator-News/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin checks if the user exists in the database . If found it will compare the given password and password found in database.*/
func TryLogin(email string, password string) (models.User, bool) {
	use, found, _ := CheckUserAlreadyExists(email)
	if found == false {
		return use, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(use.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return use, false
	}
	return use, true
}
