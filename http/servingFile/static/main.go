package main

import (
	"net/http"
)

/*
파일서버 루트에 index.html이 있으면
예외적으로 파일 구조가 보이지 않고
자동으로 index.html이 서빙된다.
*/
func main(){
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	http.ListenAndServe(":8080", nil)
}
