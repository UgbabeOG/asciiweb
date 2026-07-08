package main

import (
	"github.com/UgbabeOG/asciiweb/internals/handlers"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/", handlers.HandleHome)
	router.HandleFunc("/ascii-art", handlers.HandleASCII)
	router.HandleFunc("/download", handlers.HandleDownload)
	fmt.Println("starting server...")
	fmt.Println("server started... listening on port:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(fmt.Errorf("unable to start server on port 8080. %w", err))
		return
	}
}
