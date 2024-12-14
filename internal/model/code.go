package model

import "fmt"

type Code struct {
	phoneNumber int
	message     string
}

func test(m Code) {
	fmt.Println(m.phoneNumber)
	fmt.Println(m.message)
}

func main() {

	test(Code{phoneNumber: 1, message: "hello"})
}
