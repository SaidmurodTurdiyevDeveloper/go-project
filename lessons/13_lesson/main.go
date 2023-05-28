package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// const myurl string = "https://www.youtube.com/watch?v=xYKzdx-B1w0"
const geturl string = "https://rickandmortyapi.com/api/character"
const posturl string = "https://rickandmortyapi.com/api/character"

func main() {
	fmt.Println("This my first network request")
	/*
		defer func() {
			var result = recover()
			fmt.Print("Finish")
			if result != nil {
				fmt.Println(result)
			}
		}()
		result, err := url.Parse(myurl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success")
		fmt.Println(result)
	*/
	/*
		fmt.Print("Scheme:")
		fmt.Println(result.Scheme)
		fmt.Print("Host:")
		fmt.Println(result.Host)
		fmt.Print("Path:")
		fmt.Println(result.Path)
		fmt.Print("Port:")
		fmt.Println(result.Port())
		fmt.Print("RawQuery:")
		fmt.Println(result.RawQuery)
	*/

	/*
		anotherUrl := &url.URL{
			Scheme:   "https",
			Host:     "www.youtube.com",
			Path:     "/watch",
			RawQuery: "v=xYKzdx-B1w0",
		}
		fmt.Println("Another")
		fmt.Print("Scheme:")
		fmt.Println(anotherUrl.Scheme)
		fmt.Print("Host:")
		fmt.Println(anotherUrl.Host)
		fmt.Print("Path:")
		fmt.Println(anotherUrl.Path)
		fmt.Print("Port:")
		fmt.Println(anotherUrl.Port())
		fmt.Print("RawQuery:")
		fmt.Println(anotherUrl.RawQuery)
	*/
	/*
		qparams := result.Query()
		fmt.Printf("The type of query params are:%T\n", qparams)

		for _, v := range qparams {
			fmt.Println("Param is:", v)
		}
	*/
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostSendDataRequest()
}

func PerformGetRequest() {
	response, err := http.Get(geturl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	writeFile("./httpsGetRequestResult.txt", string(content))
}

func writeFile(fileName string, text string) {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success-", length)
}

func PerformPostRequest() {
	requestBody := strings.NewReader(`
	"id":"1"
	`)
	response, err := http.Post(posturl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	writeFile("./httpsPostRequestResult.txt", string(content))
}
func PerformPostSendDataRequest() {

	data := url.Values{}
	data.Add("id", "0")
	data.Add("firstName", "Saidmurod")
	data.Add("lastName", "Turdiyev")

	response, err := http.PostForm(posturl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	writeFile("./httpsPostSenDataRequestResult.txt", string(content))
}
