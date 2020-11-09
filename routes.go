package main

import (
	"net/http"
	"log"
)


//hangleRequests
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getArticle)
	http.HandleFunc("/post", PostArticle)
	http.HandleFunc("/articles/search", searchArticle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}