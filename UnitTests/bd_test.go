package UnitTests

import (
	"context"
	"testing"
	"time"

	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCheckUserAlreadyExists(t *testing.T) {
	user, bExists, ID := bd.CheckUserAlreadyExists("testuser@se.com")

	assert.Equal(t, true, bExists, "User doesn't exist")
	assert.Equal(t, "Test", user.Name, "User First Name doesn't Match")
	assert.Equal(t, ID, user.ID.Hex(), "User ID doesn't match")
}

func TestConnectorBD(t *testing.T) {
	client := bd.ConnectorBD()
	if client == nil {
		t.Errorf("The Database connection failed check the DBName")
	}
}

func TestCheckConnection(t *testing.T) {
	assert.Equal(t, 1, bd.CheckConnection(), "Connection Lost")
}

func TestEncryptPassword(t *testing.T) {
	encrypt, _ := bd.EncryptPassword("123456")
	bEncrypted := (encrypt != "123456")
	assert.Equal(t, true, bEncrypted, "Encryption didn't work")
}

func TestInsertPost(t *testing.T) {
	post := models.RecordPost{"624764520f5c75a01e16720f", "Test User Message", time.Now()}

	ID, bInserted, error := bd.InsertPost(post)

	if ID == "" {
		t.Errorf("ID has not been generated for the inserted post")
	}

	assert.Equal(t, true, bInserted, "Insertion of post failed")
	assert.Equal(t, nil, error, "An error occured when inserting the post")

	assert.Equal(t, true, deleteTestPostInsert(ID), "Test user inserted post deletion failed")
}

func deleteTestPostInsert(postID string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database(("gatorNews"))
	col := db.Collection("news")
	condition := bson.M{"_id": postID}
	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false
	} else {
		return true
	}
}

func TestInsertRegistration(t *testing.T) {
	User := models.User{primitive.NewObjectID(), "Test", "User", time.Now(), "testuser1@se.com", "123456", "", "", "", "", ""}
	ObjID, bInserted, error := bd.InsertRegistration(User)

	if ObjID == "" {
		t.Errorf("User ID has not been generated for the user " + User.Name)
	}
	assert.Equal(t, true, bInserted, "Insertion of the user "+User.Name+" failed")
	assert.Equal(t, nil, error, "An error occured when inserting the user "+User.Name)

	assert.Equal(t, true, deleteTestUser("testuser1@se.com"), "Failed to delete user")
}

func TestTryLogin(t *testing.T) {
	email := "testuser@se.com"
	password := "123456"
	user, bLoginSucceeded := bd.TryLogin(email, password)

	assert.Equal(t, "Test", user.Name, "User First Name doesn't Match")
	assert.Equal(t, email, user.Email, "Email doesn't Match")
	assert.Equal(t, true, bLoginSucceeded, "Login Failed")
}

func TestSearchProfile(t *testing.T) {
	ID := "624764520f5c75a01e16720f"
	user, error := bd.SearchProfile(ID)
	assert.Equal(t, "Test", user.Name, "User First Name doesn't Match")
	assert.Equal(t, "testuser@se.com", user.Email, "Email doesn't Match")
	assert.Equal(t, nil, error, "An error occured when Searching the profile")
}

func TestModifyRecord(t *testing.T) {
	time := time.Now()
	ID := "624764520f5c75a01e16720f"
	User := models.User{primitive.NewObjectID(), "Test", "User", time, "testuser@se.com", "123456", "", "", "", "", ""}

	bSuccess, error := bd.ModifyRecord(User, ID)

	assert.Equal(t, true, bSuccess, "Profile modification failed")
	user, _ := bd.SearchProfile(ID)
	expectedYear, expectedMonth, expectedDay := time.UTC().Date()
	actualYear, actualMonth, actualDay := user.BirthDate.Date()
	assert.Equal(t, expectedYear, actualYear, "Year in Date of Birth doesn't match")
	assert.Equal(t, expectedMonth, actualMonth, "Month in Date of Birth doesn't match")
	assert.Equal(t, expectedDay, actualDay, "Day in Date of Birth doesn't match")
	assert.Equal(t, nil, error, "An error occured when Modifying the Record")
}

func TestConsultRelation(t *testing.T){
	UserID  := "6220338c9f185341ce8a7d93"
	UserRelationshipID := "621bf76b40beb317228291e2"
	st,_:= bd.ConsultRelation(models.Relationship{UserID, UserRelationshipID})
	assert.Equal(t, true, st, "Failed")
}

func TestInsertRelationship(t *testing.T){
	User := models.Relationship{"6220338c9f185341ce8a7d93","6243500b39758ce6a8d758ae"}
	bInserted, error := bd.InsertRelationship(User)

	assert.Equal(t, true, bInserted, "Insertion of the relationship "+User.UserID+" failed")
	assert.Equal(t, nil, error, "An error occured when inserting the relationship "+User.UserID)

	
}

func TestReadAllUsers(t *testing.T){
	

	_,bSuccess := bd.ReadAllUsers("6220338c9f185341ce8a7d93",1,"","follow")

	assert.Equal(t, true, bSuccess, "Reading All users failed")
	
}
func TestReadFollowersPosts(t *testing.T){
	

	_,bSuccess := bd.ReadFollowersPosts("6220338c9f185341ce8a7d93",1)

	assert.Equal(t, true, bSuccess, "Reading Followers Posts failed")
	
}

func TestDeleteRelationship(t *testing.T){
	UserID  := "6220338c9f185341ce8a7d93"
	UserRelationshipID := "621bf76b40beb317228291e2"
	st,_:= bd.ConsultRelation(models.Relationship{UserID, UserRelationshipID})
	assert.Equal(t, true, st, "Failed")
}
func TestProfileLookup(t *testing.T) {
	ID := "6243500b39758ce6a8d758ae"
	user, _ := bd.ProfileLookup(ID)
	assert.Equal(t, "Jayanth", user.Name, "User First Name doesn't Match")
	
}