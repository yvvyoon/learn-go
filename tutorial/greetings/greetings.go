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

func Hellos(names []string) (map[string]string, error) {
	// map: key-value 형태의 자료구조
	// ex) map[key의 type]value의 type
	// make: map을 생성하는 함수
	messages := make(map[string]string)

	// range는 index와 item을 리턴
	// Python과 마찬가지로 index가 필요없으면 언더스코어로 처리
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
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