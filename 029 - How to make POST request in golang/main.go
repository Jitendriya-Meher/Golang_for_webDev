package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("start")

	PerformRequest()
	PerformPostJsonRequest()
	PerformRequest()

}

func PerformRequest(){

	const myurl = "http://localhost:8000/get";

	res,err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("status code:",res.StatusCode)
	fmt.Println("content length:",res.ContentLength)

	content,err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content =", string(content))

}

func PerformPostJsonRequest (){

	const myurl = "http://localhost:8000/post";

	// fake json payload
	reqBody := strings.NewReader(`
	{
		"coursename":"let's go with golang",
		"price":0,
		"platform":"online"
	}
	`)

	res,err := http.Post(myurl,"application/json",reqBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content,err := ioutil.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}

	fmt.Println("response =",string(content))

}