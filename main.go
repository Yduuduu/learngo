package main

import (
	"fmt"
)

func main() {
	a := 2
	// b := a
	// a = 10

	b := &a
	*b = 20 // 주소값을 바라보기 때문에 해당 값이 바뀌면 a도 바뀌게 됨.

	// fmt.Println(a, b)
	fmt.Println(&a, &b) // 메모리 주소 확인
	fmt.Println(*b)
	fmt.Println(a)
}
