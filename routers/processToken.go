package routers

import (
	"errors"
	"strings"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email value used in all EndPoints*/
var Email string

/*IDUser is the returned ID of the model, which will be used in all EndPoints*/
var IDUser string

/*ProcessToken token process to extract its values*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPassword := []byte("MastersofDevelopmentFacebookGroup")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Format of Token is Invalid!")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})
	if err != nil {
		_, found, _ := bd.CheckUserAlreadyExists(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Invalid!")
	}
	return claims, false, string(""), err
}
