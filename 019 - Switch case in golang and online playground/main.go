package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var dice = rand.Intn(6) + 1

	fmt.Println("dice :",dice)

	switch dice{
	case 1 :
		fmt.Println("Dice 1")
		break
	case 2 :
		fmt.Println("Dice 2")
		break
	case 3 :
		fmt.Println("Dice 3")
		fallthrough
	case 4 :
		fmt.Println("Dice 4")
		fallthrough
	case 5 :
		fmt.Println("Dice 5")
		break
	case 6 :
		fmt.Println("Dice 6")
		break
	default:
		fmt.Println("Invalid Dice")
	}
	

}