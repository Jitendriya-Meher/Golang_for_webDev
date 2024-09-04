package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("files")

	content := "This need to go in a file okk";

	file,err := os.Create("./myLcoGoFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)


	length,err := io.WriteString(file,content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("length is",length)

	file.Close()

	readFile("./myLcoGoFile.txt")

}

func readFile(filename string){

	dataByte,err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("data =",dataByte)
	fmt.Println("string data =", string(dataByte))

}

func checkNilError( err error){
	if err != nil {
		panic(err)
	}
}