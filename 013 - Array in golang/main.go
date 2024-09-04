package main

import "fmt"

func main(){

	var list [10]string;

	list[0] = "apple";
	list[1] = "orange";
	list[2] = "bannana";
	list[3] = "burger";
	
	fmt.Println("list =",list);

	fmt.Println("len list =",len(list));

	var list2 = []string{"a","b","c","d","e"};

	fmt.Println("list2 =",list2);
	fmt.Println("len list2 =",len(list2));

}