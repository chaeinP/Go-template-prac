package main

import (
	"net/http"
)


func main(){
	/*
	http.FileServer는 단일 파일이 아닌 디렉토리 단위로 리소스를 제공한다.
	FileServer는 handler를 반환한다.
	HandleFunc가 아닌 Handle 메소드를 사용하는 이유가 바로 이 때문이다.
	*/
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080",nil)
}

