package UnitTests

import(

	"fmt"
	"strings"
	"testing"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/SaiKumarMalve/Gator-News/middlew"
	"github.com/SaiKumarMalve/Gator-News/models"
	"github.com/SaiKumarMalve/Gator-News/routers"
	"github.com/stretchr/testify/assert"
)

// checkBD function unit test case
func TestCheckBD(t *testing.T){
	var users []models.User

	request, _ := http.NewRequest("POST", "https://gatornews.herokuapp.com/login", strings.NewReader("{\"email\": \"testuser@se.com\",\"password\": \"123456\"}"))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(users) >= 1 {
		idNotNull := len(users[0].ID) > 0
		nameNotNull := len(users[0].Name) > 0
		assert.Equal(t, true, idNotNull, "Non empty name expected")
		assert.Equal(t, true, nameNotNull, "Non empty name expected")
	}


}

func TestValidJWT(t *testing.T){
	tkn := "BearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MjIwMzEwNDE2NWVmYzAyZTA0YmU2MmUiLCJiaW9ncmFwaHkiOiIiLCJiaXJ0aGRhdGUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImVtYWlsIjoidGVzdHVzZXJAc2UuY29tIiwiZXhwIjoxNjQ2MzYzMjc1LCJsYXN0bmFtZSI6IlVzZXIiLCJsb2NhdGlvbiI6IiIsIm5hbWUiOiJUZXN0Iiwid2Vic2l0ZSI6IiJ9.7N25ujcLKWSAE8M_wMAxMEFYqGySpHdiXrsDK6ae46M"
	if tkn == "" {
		t.Errorf("Given Token is not Valid")

	}

}




