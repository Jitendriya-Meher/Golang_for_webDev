package main

import (
	"fmt"
)

func main() {

	var loginCount = 23;
	
	if( loginCount < 10){
		fmt.Println("Regular User")
	}else if ( loginCount < 20) {
		fmt.Println("Watch Out")
	}else{
		fmt.Println("Special user")
	}

	if( loginCount % 2 == 0){
		fmt.Println("login count is even")
	}else{
		fmt.Println("login count is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	}else{
		fmt.Println("Num is not less then 10")
	}

	var err *int= nil

	if (err != nil){
		fmt.Println("no error found")
	}else{
		fmt.Println("error found")
	}
}