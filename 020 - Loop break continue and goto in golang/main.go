package main

import "fmt"

func main() {

	nums := []int{2,3,4,5,6,7,8,9}

	fmt.Println("nums",nums)

	for i := 0; i < len(nums); i++ {
		fmt.Printf("num[%d] = %d\n", i, nums[i])
	}

	fmt.Println("")

	for i := range nums{
		fmt.Printf("num[%d] = %d\n", i, nums[i])
	}

	fmt.Println("")

	for i,num := range nums{
		fmt.Printf("num[%d] = %d\n", i, num)
	}

	fmt.Println("")

	var i = 0

	for i < 10{

		if( i == 5){
			goto loc
			break
		}
		if( i==2){
			i++;
			continue
		}
		
		fmt.Println("val is",i)
		i++;
	}

	print("skip this line")

	loc:
	fmt.Println("Hii i am free now")
}