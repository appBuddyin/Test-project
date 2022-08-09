package main

import"fmt"

func main(){
	var value int = 2

	switch{
	case value ==1 :
		fmt.Println("hello")
	case value == 2:
		fmt.Println("hii")
	case value == 3:
		fmt.Println("namastey")
	default:
		fmt.Println("invalid")
	}
}