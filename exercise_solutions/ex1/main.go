package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed english_rights.txt
var englishRights string

//go:embed french_rights.txt
var frenchRights string

//go:embed spanish_rights.txt
var spanishRights string

func main() {
	process(os.Args)
}

func process(args []string) {
	if len(args) == 1 {
		fmt.Println("no language provided")
		return
	}
	switch strings.ToLower(args[1]) {
	case "english":
		fmt.Println(englishRights)
	case "spanish":
		fmt.Println(spanishRights)
	case "french":
		fmt.Println(frenchRights)
	default:
		fmt.Println("unknown language:", os.Args[1])
	}
}
