package main

import (
	"net/http"
	"io"
)

/*
http.StripPrefix는 url 패스에서 첫번째 파라미터값을 제거하고
그 다음 패스 경로와 파일서버에 등록된 asset 경로를 비교하여
경로가 동일한 리소스를 서빙한다.
*/

func main(){
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets")) ))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}
