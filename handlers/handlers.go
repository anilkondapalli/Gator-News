package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SaiKumarMalve/Gator-News/middlew"
	"github.com/SaiKumarMalve/Gator-News/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers set the PORT for the servers to listen*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registration", middlew.CheckBD(routers.Registration)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", middlew.CheckBD(middlew.ValidJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckBD(middlew.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/post", middlew.CheckBD(middlew.ValidJWT(routers.RecordPost))).Methods("POST")
	router.HandleFunc("/readPosts", middlew.CheckBD(middlew.ValidJWT(routers.ReadPosts))).Methods("GET")
	router.HandleFunc("/removePost", middlew.CheckBD(middlew.ValidJWT(routers.RemovePost))).Methods("DELETE")
	router.HandleFunc("/uploadAvatar", middlew.CheckBD(middlew.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckBD(middlew.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckBD(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/highRelationship", middlew.CheckBD(middlew.ValidJWT(routers.HighRelationship))).Methods("POST")
	router.HandleFunc("/lowRelationship", middlew.CheckBD(middlew.ValidJWT(routers.LowRelationship))).Methods("DELETE")
	router.HandleFunc("/consultRelation", middlew.CheckBD(middlew.ValidJWT(routers.ConsultRelation))).Methods("GET")
	router.HandleFunc("/userList", middlew.CheckBD(middlew.ValidJWT(routers.UserList))).Methods("GET")
	router.HandleFunc("/readFollowersPosts", middlew.CheckBD(middlew.ValidJWT(routers.ReadFollowersPosts))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
