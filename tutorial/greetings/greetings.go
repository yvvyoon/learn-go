package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v.  Welcome!", name)
	// := 연산자는 선언과 할당을 동시에 함
	// var message string
	// message = fmt.Sprintf("Hi, %v.  Welcome!", name)

	return message, nil
}