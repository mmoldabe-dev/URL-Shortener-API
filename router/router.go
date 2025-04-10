package router

import (
	"URL-Shortener-API/handler"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(port string, DB *sql.DB) {

	r := mux.NewRouter()

	r.HandleFunc("/shortener", func(w http.ResponseWriter, r *http.Request) {
		handler.ShortenHandler(w, r, DB)
	})
	r.HandleFunc("/shortener/{code}", func(w http.ResponseWriter, r *http.Request) {

		handler.ShortenHandlerId(w, r, DB)
	})

	log.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
