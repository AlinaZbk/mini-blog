package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlinaZbk/mini-blog.git/handler"
)

func main() {
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/posts", handler.PostsHandler)
	http.HandleFunc("/posts/", handler.PostByIDHandler)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
