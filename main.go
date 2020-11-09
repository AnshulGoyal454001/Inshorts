package main

//Article is a type
type Article struct {
	Title    string `json:"Title"`
	Subtitle string `json:"Subtitle"`
	Content  string `json:"content"`
}

//Articles is a database simulation
var Articles []Article

func main() {
	handleRequests()
}
