package main

import "fmt"

func main() {
	switch class := 4; class {
	case 2:
		fmt.Println("maths class")
	case 4:
		fmt.Println("chemistry class")
	case 5:
		fmt.Println("computer class")
		fallthrough
	case 7:
		fmt.Println("physics class")
	default:
		fmt.Println("no information")
	}
}
