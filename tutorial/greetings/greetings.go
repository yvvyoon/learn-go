package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// := 연산자는 선언과 할당을 동시에 함
	// var message string
	// message = fmt.Sprintf("Hi, %v.  Welcome!", name)

	return message, nil
}

// 전역변수들이 초기화된 이후에 자동으로 실행되는 init() 메소드
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}