package main

import "fmt"

// jwtToken := 30
var jwtToken = 30

const pi float64 = 3.1416

func main() {

	fmt.Println("variables");

	var username = "Jitendtiya Meher";
	fmt.Println(username);
	
	fmt.Printf("%s is type : %T \n", username,username);

	var isBool bool = false;
	fmt.Printf("bool is type : %T \n",isBool);

	var smallNum uint8 = 23;
	fmt.Printf("variable is type : %T \n",smallNum);

	var intNum int = 12223;
	fmt.Printf("variable is type : %T \n",intNum);

	// var floatNum float32 = 12223.344123455666;
	var floatNum float64 = 12223.344123455666;
	fmt.Printf("variable is type : %T \n",floatNum);
	fmt.Println(floatNum)


	// default values and some aliases
	var variable int;
	fmt.Printf("variable : %T \n",variable);
	fmt.Println(variable)

	// implicit type
	var website = "www.google.com";
	fmt.Println(website);
	// website = 3;

	// no var style
	numOfUsr := 30000.5
	fmt.Println(numOfUsr)

	fmt.Println(pi)

}