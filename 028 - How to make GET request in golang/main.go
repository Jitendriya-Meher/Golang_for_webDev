package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("start")

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