package main

import (
	"io"
	"net/http"
)

type hotdog int
func (a hotdog)ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "walf")
}

type hotcat int
func (a hotcat)ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "meow")
}

func main(){
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	/*
	멀티플렉서(multiplexer) 또는 mux는 여러 아날로그 또는 디지털 입력 신호 중 하나를 선택하여 선택된 입력을 하나의 라인에 전달하는 장치이다.
	*/
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)
	/*
	새로 정의한 먹스는 Handle 메소드를 이용해 url 경로를 연결할 수 있다.
	*/
	http.ListenAndServe(":8080", mux)
}
