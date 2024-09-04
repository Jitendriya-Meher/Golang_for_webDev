package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev";

func main() {

	fmt.Println("main")

	response,err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// fmt.Println("response =",response)
	fmt.Printf(" type of response = %T \n",response)

	// close the connection
	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("response body =",string(databytes))

}