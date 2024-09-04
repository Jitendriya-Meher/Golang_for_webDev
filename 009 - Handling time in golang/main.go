package main

import (
	"fmt"
	"time"
)

func main(){

	presentTime := time.Now()

	fmt.Println("time : ",presentTime)
	fmt.Println("time : ",presentTime.Format("01-01-2001"))
	fmt.Println("time : ",presentTime.Format("01-01-2001 15:04:05 Monday"))

	createDate := time.Date(2001,time.November,20,10,30,00,00,time.UTC)
	fmt.Println("Create Time : ",createDate)
	fmt.Println("Create Time : ",createDate.Format("01-01-2001 15:04:05 Monday"))

}