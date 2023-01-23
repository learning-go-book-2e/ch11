package main

import "fmt"

func main() {
	x := 10 // this assignment isn't read!
	x = 20
	fmt.Println(x)
	x = 30 // this assignment isn't read!
}
