package main

import "fmt"

type Direction int

const (
	_ Direction = iota
	North
	South
	East
	West
)

//go:generate stringer -type=Direction

func main() {
	fmt.Println(North.String())
}
