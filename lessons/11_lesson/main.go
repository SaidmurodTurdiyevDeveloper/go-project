package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("welcome to my file examples")

	var currentTime = time.Now()
	var content = "This my file content text, this writed on:" + currentTime.Format("2006:01:01 Mon 15:04:05")

	file, err := os.Create("./myExampleFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Succcessful print text:", length)

	defer file.Close()
	readFile("./myExampleFile.txt")
}

func readFile(fileName string) {

	databyte, err := ioutil.ReadFile(fileName)

	checkNilError(err)
	fmt.Println("Text data inside the file is \ndatabyte: ", databyte)
	fmt.Println("text:", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
