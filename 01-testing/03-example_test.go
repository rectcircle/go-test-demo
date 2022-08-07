package testingdemo

import "fmt"

// go test -run ^Example$ ./01-testing

func Example() {
	fmt.Println("This is a package example")
	// Output: This is a package example
}

func Example_a01() {
	fmt.Println("This is a package example")
	// Output: This is a package example
}

func ExampleHello() {
	Hello()
	// Output: hello
}

func ExampleHello_a01() {
	Hello()
	// Output: hello
}

func ExampleSalutations() {
	Salutations()
	// Output:
	// hello, and
	// goodbye
}

func ExampleT_M() {
	t := T{}
	t.M()
	// Output: t.m()
}

func ExampleT_M_a01() {
	t := T{}
	t.M()
	// Output: t.m()
}
