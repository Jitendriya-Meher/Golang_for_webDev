package main

import "fmt"

func add(  a int,  b int) int{
	return a+b
}

func proAdder( values ...int) (int,string) {

	total := 0;

	for _,val := range values {
		total += val
	}

	return total,"Calculation done"

}

func main() {

	fmt.Println("hii from main")

	greet()

	// func greetInsideMain()  {
	// 	fmt.Println("Greet inside main")
	// }

	// greetInsideMain()

	var res = add( 2,3 )
	fmt.Println("result :",res)

	res = add( 4,5)
	fmt.Println("result :",res)

	res,msg := proAdder( 1,2,3,4,5,6,7,8,9,10)
	fmt.Println("result :",res)
	fmt.Println(msg)


}

func greet(){
	fmt.Println("Namestey from greet")
}