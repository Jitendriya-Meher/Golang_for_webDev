package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Print("Please rate our pizza between 1 to 5 : ")

	reader := bufio.NewReader(os.Stdin)

	rate,_ := reader.ReadString('\n')

	fmt.Println("Thanks for rating : ",rate)

	newRate,err := strconv.ParseFloat(strings.TrimSpace(rate),64) 

	if( err != nil){
		// newRate = newRate + 1
		fmt.Println("New rating : ",newRate)
	}else{
		fmt.Println("error in rating ",err)
	}


}