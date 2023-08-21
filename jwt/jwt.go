package jwt

import (
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Generates a JSON Web Token which is valid through out the login session
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MastersofDeveleopment_groupofFacebook")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthdate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
