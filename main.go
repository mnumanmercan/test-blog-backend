package main

import (
	"log"
	"net/http"

	"blog-backend/database"
	"blog-backend/handlers"
	"blog-backend/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectMongo()

	r := mux.NewRouter()

	r.HandleFunc("/blogs", handlers.GetBlogs).Methods("GET")
	r.HandleFunc("/blogs", handlers.CreateBlog).Methods("POST")

	handler := middlewares.CORS(r)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", handler)
}
