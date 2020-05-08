package main

import (
	"fmt"
	"net/http"
	"strings"
)

// GET KV PAIR
func handleGetKV(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("called the get method")
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	var kvPair [2]string
	index := 0

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		//if its the actual form data:
		kvPair[index] = strings.Join(v, "")
		index++
	}
	//index = 0 should be the key
	getKVPair(kvPair[0])

	//convert response to proper response
	/*resOut := ResponseInt{int(result["statusCode"].(float64)), result["message"].(string), ConvertMapInterfaceToMapInt(result["data"])}
	jOut, _ := resOut.JSON()

	//return response to client
	w.WriteHeader(int(result["statusCode"].(float64)))
	fmt.Fprintf(w, jOut)
	*/
}

//POST A KV PAIR
func handlePutKV(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called the put method")

	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	var kvPair [2]string
	index := 0

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		//if its the actual form data:
		kvPair[index] = strings.Join(v, "")
		index++
	}
	//assuming index = 0 is the key, index = 1 is the value,
	putKVPair(kvPair[0], kvPair[1])
}
