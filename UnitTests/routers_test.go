package UnitTests

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/middlew"
	"github.com/SaiKumarMalve/Gator-News/models"
	"github.com/SaiKumarMalve/Gator-News/routers"
	"github.com/stretchr/testify/assert"
)

//Registration function Unit test case
func TestRegistration(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("POST", "https://gatornews.herokuapp.com/registration", strings.NewReader("{\"name\": \"Test\",\"lastname\": \"User\",\"email\": \"testuser1@se.com\",\"password\": \"123456\"}"))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/registration", middlew.CheckBD(routers.Registration)).Methods("POST")
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
	assert.Equal(t, true, deleteTestUser("testuser1@se.com"), "Failed to delete user")

}

//Login function Unit test case
func TestLogin(t *testing.T) {
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
func deleteTestUser(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database(("gatorNews"))
	col := db.Collection("users")
	condition := bson.M{"email": email}
	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false
	} else {
		return true
	}
}

func TestViewProfile(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("GET", "https://gatornews.herokuapp.com/viewProfile?id=624764520f5c75a01e16720f", strings.NewReader(``))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/viewProfile", middlew.CheckBD(routers.ViewProfile)).Methods("GET")
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

func TestModifyProfile(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("PUT", "https://gatornews.herokuapp.com/modifyProfile", strings.NewReader("{\"name\": \"Test\",\"lastname\": \"User\",\"birthDate\": \"1995-02-23T00:00:00Z\",\"location\": \"GNV\",\"biography\": \"I am a student\",\"website\": \"https://github.com/SaiKumarMalve/Gator-News\"}"))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/modifyProfile", middlew.CheckBD(routers.ModifyProfile)).Methods("PUT")
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
func TestRecordPostRouter(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("POST", "http://gatornews.herokuapp.com/post", strings.NewReader("{\"message\": \"For Test Purpose\"}"))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/post", middlew.CheckBD(routers.RecordPost)).Methods("POST")
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
func TestReadPostsRouter(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("GET", "https://gatornews.herokuapp.com/readPosts?id=6220338c9f185341ce8a7d93&page=1", strings.NewReader(``))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/readPosts", middlew.CheckBD(routers.ReadPosts)).Methods("GET")
	router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(users) >= 1 {
		idNotNull := len(users[0].ID) > 0

		assert.Equal(t, true, idNotNull, "Non empty name expected")
	}

}
func TestProcessToken(t *testing.T) {
	tkn := "BearereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MjIwMzEwNDE2NWVmYzAyZTA0YmU2MmUiLCJiaW9ncmFwaHkiOiIiLCJiaXJ0aGRhdGUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImVtYWlsIjoidGVzdHVzZXJAc2UuY29tIiwiZXhwIjoxNjQ2MzYzMjc1LCJsYXN0bmFtZSI6IlVzZXIiLCJsb2NhdGlvbiI6IiIsIm5hbWUiOiJUZXN0Iiwid2Vic2l0ZSI6IiJ9.7N25ujcLKWSAE8M_wMAxMEFYqGySpHdiXrsDK6ae46M"
	if tkn == "" {
		t.Errorf("Given Token is not Valid")

	}

}
func TestReadFollowersPostsRouter(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("GET", "https://gatornews.herokuapp.com/readFollowersPosts?page=1", strings.NewReader(``))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/readFollowersPosts", middlew.CheckBD(routers.ReadFollowersPosts)).Methods("GET")
	router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
	if err != nil {
		fmt.Println("err is ", err)
	}


}

func TestRemovePostRouter(t *testing.T) {
	var users []models.User

	request, _ := http.NewRequest("DELETE", "https://gatornews.herokuapp.com/removePost?id=6238b1b33af5f47b275985e3", strings.NewReader(``))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/removePost", middlew.CheckBD(routers.RemovePost)).Methods("DELETE")
	router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
	if err != nil {
		fmt.Println("err is ", err)
	}

}

func TestUserListRouter(t *testing.T){
	var users []models.User
	
		request, _ := http.NewRequest("GET", "https://gatornews.herokuapp.com/userList?type=new&page=1&search=", strings.NewReader(``))
		response := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/userList", middlew.CheckBD(routers.UserList)).Methods("GET")
		router.ServeHTTP(response, request)
		assert.Equal(t, 201, response.Code, "OK response is expected")
		err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
		if err != nil {
			fmt.Println("err is ", err)
		}
	
	}

	func TestLowRelationshipRouter(t *testing.T) {
		var users []models.User
	
		request, _ := http.NewRequest("DELETE", "https://gatornews.herokuapp.com/lowRelationship?id=6220338c9f185341ce8a7d93", strings.NewReader(``))
		response := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/lowRelationship", middlew.CheckBD(routers.LowRelationship)).Methods("DELETE")
		router.ServeHTTP(response, request)
		assert.Equal(t, 201, response.Code, "OK response is expected")
		err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
		if err != nil {
			fmt.Println("err is ", err)
		}
	
	
	}
	

	func TestHighRelationshipRouter(t *testing.T) {
		var users []models.User
	
		request, _ := http.NewRequest("POST", "https://gatornews.herokuapp.com/highRelationship?id=621bf76b40beb317228291e2", strings.NewReader(``))
		response := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/highRelationship", middlew.CheckBD(routers.HighRelationship)).Methods("POST")
		router.ServeHTTP(response, request)
		assert.Equal(t, 201, response.Code, "OK response is expected")
		err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
		if err != nil {
			fmt.Println("err is ", err)
		}
	
	
	}

	func TestConsultRelationRouter(t *testing.T){
		var users []models.User
		
			request, _ := http.NewRequest("GET", "https://gatornews.herokuapp.com/consultRelation?id=621bf76b40beb317228291e2", strings.NewReader(``))
			response := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/consultRelation", middlew.CheckBD(routers.ConsultRelation)).Methods("GET")
			router.ServeHTTP(response, request)
			assert.Equal(t, 201, response.Code, "OK response is expected")
			err := json.Unmarshal([]byte(response.Body.Bytes()), &users)
			if err != nil {
				fmt.Println("err is ", err)
			}
		
		
		}