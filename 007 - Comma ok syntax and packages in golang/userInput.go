package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	welcome := "Welcome to user Input";
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the rating for our pizza : ");

	// comma ok // err ok
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating",input)
	fmt.Printf("type : %T",input)
	fmt.Println("\nerror",err)

}