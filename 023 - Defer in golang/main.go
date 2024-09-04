package main

import "fmt"

func main() {

	defer fmt.Println("World")
	defer fmt.Println("World 1")
	defer fmt.Println("World 2")

	fmt.Println("hello")

	myDefer()

	defer fmt.Println("world 3")

}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println("fun ",i)
	}
}