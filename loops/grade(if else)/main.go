package main

import "fmt"

func main() {
	marks := 76

	if marks > 90 {
		fmt.Println("grade A")
	} else if marks > 80 {
		fmt.Println("grade B")
	} else if marks > 70 {
		fmt.Println("grade C")
	} else if marks > 60 {
		fmt.Println("grade D")
	} else {
		fmt.Println("fail")
	}
}
