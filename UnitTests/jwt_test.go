package UnitTests

import (
	"testing"
	"time"

	"github.com/SaiKumarMalve/Gator-News/jwt"
	"github.com/SaiKumarMalve/Gator-News/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGenerateJWT(t *testing.T) {

	User := models.User{primitive.NewObjectID(), "Test", "User", time.Now(), "testuser@se.com", "123456", "", "", "", "", ""}

	tokenString, error := jwt.GenerateJWT(User)

	bTokenGenerated := (tokenString == "")
	assert.Equal(t, true, !bTokenGenerated, "JWT token not generated")

	bError := (error == nil)
	assert.Equal(t, true, bError, "There was an error generating JWT token")

}
