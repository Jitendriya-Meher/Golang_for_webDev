package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=abcd1234"

func main() {

	fmt.Println("Handling URL")
	fmt.Println("url =",myurl)

	// parsing the url
	result,_ := url.Parse(myurl)

	fmt.Println("parse url =",result)
	fmt.Println("parse =",result.Scheme)
	fmt.Println("parse =",result.Path)
	fmt.Println("parse =",result.Port())
	fmt.Println("parse =",result.RawQuery)

	qparams := result.Query()
	fmt.Println("params =",qparams)

	fmt.Println("params =",qparams["coursename"])

	for _,val := range qparams{
		fmt.Println("course =",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=jiten",
	}

	fmt.Println("url =",partsOfUrl)

	stringUrl := partsOfUrl.String()

	fmt.Println("url =",stringUrl)

}