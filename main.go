package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sava "ascii-art-web/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	port := ":8080"
	// http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/styles.CSS")
	// })
	// http.HandleFunc("/400/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/400.CSS")
	// })
	// http.HandleFunc("/404/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/404.CSS")
	// })
	// http.HandleFunc("/405/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/405.CSS")
	// })
	// http.HandleFunc("/500/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/500.CSS")
	// })

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", sava.Handler)
	http.HandleFunc("/ascii-art", sava.HandleasciiArt)
	log.Printf("Server is running on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
