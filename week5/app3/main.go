package main

import "fmt"

func main() {
	var a int = 5
	var b int = 15
	addNumbers(a, b)
}

func addNumbers(a int, b int) {
	var c int = a + b
	fmt.Println("Sum of", a, "+", b, "is", c)

}
