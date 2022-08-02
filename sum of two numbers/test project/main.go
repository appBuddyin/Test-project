package main

import"fmt"

func main(){
	var int s1,s2,s3,s4,s5
	fmt.Println("Enter marks on five subject")
	fmt.Scanf("%d",&s1)
	fmt.Scanf("%d,&s2")
	fmt.Scanf("%d,&s3")
	fmt.Scanf("&d",s4)
	fmt.Scanf("&d",s5)
	avg:=(s1+s2+s3+s4+s5)/5
	if avg>=90{
		print("grade: A")
	}else{
		print("grade:B")
	}
}