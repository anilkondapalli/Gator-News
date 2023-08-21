package routers

import(
	"encoding/json"
	"net/http"
	"time"
	"github.com/SaiKumarMalve/Gator-News/bd"
	"github.com/SaiKumarMalve/Gator-News/models"
)
/* RecordPost allows you to record the post in the database*/
func RecordPost(w http.ResponseWriter, r *http.Request){
	var message models.Post
	err := json.NewDecoder(r.Body).Decode(&message)

	registration := models.RecordPost{
		UserID: IDUser,
		Message: message.Message,
		Date: time.Now(),
	}
	_, status, err := bd.InsertPost(registration)
	if err != nil{
		http.Error(w, "An error occured while trying to insert the record, try again"+err.Error(), 400)
		return
	}
	if status== false{
		http.Error(w, "The post could not be inserted", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
	
}