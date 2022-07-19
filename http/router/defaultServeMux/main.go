package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "walf walf")
}

func cat(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "meow")
}

func main(){
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)

	/*
	위 식을 아래와 같이 표현할 수도 있다.
	http.HandlerFunc()은 메소드가 아니라 타입 변환식이다.

	http.Handle("/dog/", http.HandlerFunc(dog))
	*/

	/*
	ListenAndServe 메소드에 두번째 인자로 nil을 넣으면
	DefaultServeMux를 쓰겠다는 뜻이다.
	DefaultServeMux는 따로 정의할 필요 없이 바로 사용할 수 있다. 이때 hanlder로는 HandleFunc 메소드를 사용해야한다. HandleFunc 메소드는 Handle 메소드와는 달리 핸들러 인터페이스 타입이 아닌 함수는 받는다.
	*/
	http.ListenAndServe(":8080",nil)

}
