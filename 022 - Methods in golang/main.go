package main

import "fmt"



func main() {

	fmt.Println("main")

	user1 := User{"jiten", "jiten@go.com", true,22}
	fmt.Println("User 1 =",user1)
	fmt.Printf("User 1 = %+v\n", user1)


	user2 := User{
		Name:   "Jitendriya Meher",
		Email:  "jiten@go.com",
		Status: false,
		Age:    22,
	}
	fmt.Printf("User 2 = %+v\n", user2)

	fmt.Println(user1.GetStatus())
	fmt.Println(user2.GetStatus())

	// pass by value
	user2.changeStatus()
	fmt.Println(user2.GetStatus())

}

type User struct {
	Name string
	Email string
	Status bool
	Age int


}

func ( u User) GetStatus () string {
	if u.Status {
		return u.Name+" is Active Now"
	}else{
		return u.Name+" is not Active Now"
	}
}

func ( u User) changeStatus () {
	
	u.Status = !u.Status

}