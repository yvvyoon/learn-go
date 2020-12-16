package main

import (
	"fmt"
)

type Printable interface {
	String() string
}

type StructA struct {
	value string
}

type StructB struct {
	value int
}

func (s *StructA) String() string {
	return "Value: " + s.value
}

// main 함수에서 StructB의 결과를 호출하고 있지만 포인터 타입의 StructB의 String()을 구현하지 않았으므로 컴파일 에러 발생
// 	main 함수에서 호출하고 있는 Printable() 메소드는 Printable이라는 인터페이스를 인자로 받음
// 	Printable 인터페이스는 String() 메소드를 포함한다고 정의하고 있는데,
// 	Println() 메소드와 StructB 구조체의 관계를 설명하고 있는 Printable 인터페이스의 메소드를 정의하지 않았기 때문

// func (s *StructB) String() string {
// 	s.value += 10

// 	fmt.Println("StructB String()")

// 	return "StructB: " + strconv.Itoa(s.value)
// }

func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	// structA := &StructA{value: "AAA다아아ㅏㅏ"}
	structB := &StructB{value: 10000}

	Println(structB)
	// # command-line-arguments
	// ./interfaceEx.go:39:9: cannot use structB (type *StructB) as type Printable in argument to Println:
	//     *StructB does not implement Printable (missing String method)
}
