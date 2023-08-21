package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SaiKumarMalve/Gator-News/bd"
)

/*ReadFollowersPosts reads the posts of all user followers*/
func ReadFollowersPosts(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Send the page number as parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Send the page parameter as an integer greater than 0", http.StatusBadRequest)
		return
	}

	answer, correct := bd.ReadFollowersPosts(IDUser, page)
	if correct == false {
		http.Error(w, "Error reading posts", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
