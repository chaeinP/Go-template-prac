package main

import (
	"net/http"
	"log"
)

func main(){
	/*
	http.ListenAndServe 는 에러를 반환한다.
	log.Fatal은 어떤 에러든 발생시 에러 로그를 찍고 프로그램을 종료시킨다.
	*/
	log.Fatal(http.ListenAndServe(":8080",nil))
}

/*
http.Error(w, "에러 메세지", 에러코드(상수))
위에는 에러 발생시 response값으로 개발자가 직접 에러 메세지를 내려주어야할 때 사용한다.

에러 코드는 숫자로도 줄 수 있고 http 패키지에 정의되어있는 상수로도 표현할 수 있다.
*/
