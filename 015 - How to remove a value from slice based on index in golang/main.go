package main

import "fmt"

func main() {

	var nums = []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("nums :",nums)

	var index int = 2;

	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("nums :",nums)

}