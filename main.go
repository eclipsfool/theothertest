package main

import (
	"html/template"
	"log"
	"net/http"
)

type HomeData struct {
	Title    string
	VideoSrc string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handleHome)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Title:    "The Other",
		VideoSrc: "/static/videos/trailer.mp4",
	}
	if err := templates.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
