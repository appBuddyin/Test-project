package main

import"fmt"

func main() {
	fmt.Println("enter three numbers:")
	var first int
	fmt.Scanln(&first)
	var second int
	fmt.Scanln(&second)
	var third int
	fmt.Scanln(&third)
	if(first>=second && first>= third){
		fmt.Println(first,"is largest among three numbers.")
	}
	if(second>=first && second>=third){
		fmt.Println(second,"is largest among three numbers.")
	}
	if(third>=first && third>=second){
		fmt.Println( third,"is largest  among three numbers.")
	}
}