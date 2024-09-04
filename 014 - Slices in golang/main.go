package main

import (
	"fmt"
	"sort"
)

func main() {

	var list1 = []int{1,2,3,4,5,6,7,8,9,10};

	fmt.Printf("type of list1 : %T\n",list1)

	fmt.Println("list1 :",list1)

	fmt.Println(list1[7])

	// inser at end

	list1 = append(list1, 11,12,13)
	fmt.Println("list1 :",list1)

	var list2 = append(list1[2:])
	fmt.Println("list1 :",list2)

	list2 = append(list1[2:6])
	fmt.Println("list1 :",list2)

	list2 = append(list1[:6])
	fmt.Println("list1 :",list2)

	list2 = append(list1[1:10])
	fmt.Println("list1 :",list2)

	score := make([]int,4);
	score[0]=10
	score[1]=200
	score[2]=30
	score[3]=400
	fmt.Println(score)

	score = append(score, 333,444,555)
	fmt.Println(score)

	// sorting 
	sort.Ints(score)
	fmt.Println(score)

	fmt.Println(sort.IntsAreSorted(score))

}