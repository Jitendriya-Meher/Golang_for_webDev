package main

import "fmt"

func main(){

	fmt.Println("Pointers")

	var one int = 101

	fmt.Println(&one)

	var ptr *int = &one

	fmt.Println(ptr)
	fmt.Println(*ptr)

	myNum := 23

	var myPtr = &myNum

	fmt.Println(myPtr)

	*ptr = *ptr + 1
	fmt.Println(one)

}