package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// 타임스탬프 또는 소스 파일 정보 미포함
	log.SetFlags(0)
	// flag 1: greetings: 2020/12/14 empty name
	// flag 2: greetings: 15:38:28 empty name
	// flag 3: greetings: 2020/12/14 15:38:28 empty name
	// flag 4: greetings: 2020/12/14 15:38:28.564935 empty name

	names := []string{"Yoon", "Park", "Min"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		// error 로그 출력과 함께 프로그램 종료
		log.Fatal(err)
	}

	fmt.Println(messages)
}