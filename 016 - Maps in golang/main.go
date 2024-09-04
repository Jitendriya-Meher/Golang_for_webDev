package main

import "fmt"

func main() {

	languages := make(map[string]string)

	// adding

	languages["Js"] = "javascript"
	languages["Py"] = "python"
	languages["Rb"] = "Ruby"
	languages["html"] = "hyper text markup language"


	fmt.Println("languages",languages)
	fmt.Println("languages js :",languages["Js"])

	// delete
	delete(languages,"Rb")
	fmt.Println("languages",languages)

	// traversal
	for key,value := range languages{
		fmt.Printf("for key %v : value is %v \n", key, value)
	}

}