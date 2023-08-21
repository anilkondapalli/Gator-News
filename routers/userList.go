package routers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/SaiKumarMalve/Gator-News/bd"
)

/* UserList reads the list of users */
func UserList(w http.ResponseWriter, r *http.Request){
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp,err := strconv.Atoi(page)
	if err != nil{
		http.Error(w,"Must send the page parameter as an integer greater than 0",http.StatusBadRequest)
		return
	}
	p:=int64(pageTemp)
	result , status := bd.ReadAllUsers(IDUser, p,search,typeUser)
	if status == false{
		http.Error(w,"Error Reading Users",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}