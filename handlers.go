package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getArticle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	// fmt.Println("var1 = ", reflect.TypeOf(r.URL.Query()["id"][0]))
	GetPost(id, w)
	json.NewEncoder(w).Encode(Articles)
}

func searchArticle(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()["q"][0]
	searchArticleByName(q)
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "inShots API")
}

//PostArticle posts the article
func PostArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("Title")
	content := r.FormValue("Content")
	subtitle := r.FormValue("Subtitle")

	Articles = append(Articles, Article{Title: title, Content: content})
	InsertPost(title, subtitle, content)
	fmt.Fprintf(w, title)
}

