package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/toby.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "toby.jpeg")
}
