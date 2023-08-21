package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/SaiKumarMalve/Gator-News/bd"
)
/* ReadPosts reads the posts*/
func ReadPosts(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w,"must send the id parameter",http.StatusBadRequest)
		return 
	}

	if len(r.URL.Query().Get("page"))<1{
		http.Error(w,"must send the page parameter",http.StatusBadRequest)
		return 
	}
	page,err:= strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil{
		http.Error(w,"must send the page parameter with a value greater than 0",http.StatusBadRequest)
		return
	}
	p:=int64(page)
	answer,right :=bd.ReadPosts(ID,p)
	if right ==false{
		http.Error(w,"Error reading the posts",http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)

}