package router

import (
	"database/sql"
	"log"
	"net/http"

	"URL-Shortener-API/auth"
	"URL-Shortener-API/handler"
	"URL-Shortener-API/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(port string, DB *sql.DB) {
	r := mux.NewRouter()

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.RegisterHandler(w, r, DB)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handler.LoginHandler(w, r, DB)
	}).Methods("POST")

	r.HandleFunc("/shortener", func(w http.ResponseWriter, r *http.Request) {
		handler.ShortenHandler(w, r, DB)
	}).Methods("POST", "GET")

	r.HandleFunc("/shortener/{code}", func(w http.ResponseWriter, r *http.Request) {
		handler.ShortenHandlerId(w, r, DB)
	}).Methods("GET")

	r.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(auth.MeHandler))).Methods("GET")

	log.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
