package main

import "fmt"

var area int

func main() {
	var l, b int
	fmt.Println("enter length of rectangle:")
	fmt.Scan(&l)
	fmt.Println("enter breadth of rectangle:")
	fmt.Scan(&b)
	area = l * b
	fmt.Println("area of rectangle:", area)
}
