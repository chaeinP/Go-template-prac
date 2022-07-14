package main

import (
	"io"
	"net/http"
)

/*
http.Handler

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

http.ListenAndServe
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
*/

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
