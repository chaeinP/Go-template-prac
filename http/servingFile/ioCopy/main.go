package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/yc.png", dogPic)
	http.ListenAndServe(":8080",nil)
}

func dogPic(w http.ResponseWriter, r *http.Request){
	f, err := os.Open("yc.png")
	if err != nil {
		fmt.Fprintln(w, err)
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
