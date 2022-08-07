package testingdemo

import "fmt"

func Hello() {
	fmt.Println("hello")
}

func Salutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
}

type T struct{}

func (t *T) M() {
	fmt.Println("t.m()")
}
