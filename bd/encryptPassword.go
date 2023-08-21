package bd

import "golang.org/x/crypto/bcrypt"

//EncryptPassword encrypts the user password for storage in Database
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err

}
