package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// http 패키지의 HandleFunc이 아닌
	// router의 HandleFunc을 호출한다는 것이 차이점
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// 요청 URL에 포함되어 있는 slug들의 map 형태
		vars := mux.Vars(r)
		// map[page:10 title:aaa]
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// ListenAndServe 메서드의 두 번째 인자로 항상 nil을 전달했는데,
	// nil을 전달하면 router로 net/http 패키지의 기본 router를 사용한다는 의미이다.
	// 여기서는 gorilla/mux의 router를 사용하므로 변경하여 전달한다.
	http.ListenAndServe(":80", router)
}

/*
1) HTTP 요청 메서드 분기
	router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

2) 호스트명과 서브도메인 설정
	router.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

3) http/https 스킴 분기
	router.HandleFunc("/secure", SecureHandler).Schemes("https")
	router.HandleFunc("/insecure", InsecureHandler).Schemes("http")

4) 경로 접두사와 서브라우터
	bookrouter := router.PathPrefix("/books").subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)
*/
