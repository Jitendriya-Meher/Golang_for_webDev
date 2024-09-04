package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	jiten := User("Jiten","jiten@go.com",true,22)

	fmt.Println(jiten)
	fmt.Printf("jiten details are : +%+v \n",jiten)
	fmt.Println("name :",jiten.Name)

}
