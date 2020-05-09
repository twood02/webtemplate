package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/hashicorp/consul/api"
)

var kv *api.KV

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Sreya!") // send data to client side
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		formName := strings.Join(r.Form["store-submit"], "")
		formName2 := strings.Join(r.Form["getval-submit"], "")
		formName3 := strings.Join(r.Form["getall-submit"], "")
		fmt.Println(formName)
		fmt.Println(strings.Compare(formName, "Store KV pair"))
		fmt.Println(strings.Compare(formName2, "Get corresponding value stored"))
		fmt.Println(strings.Compare(formName3, "Get all KV pairs stored"))

		if strings.Compare(formName, "Store KV pair") == 0 {
			key := strings.Join(r.Form["key"], "")
			value := strings.Join(r.Form["value"], "")
			fmt.Println("key:", key)
			fmt.Println("value:", value)
			response := putKVPair(key, value)
			if strings.Compare(response, "KV store success") == 0 {
				fmt.Fprintf(w, response+"\n"+"Key: "+key+"  Value: "+value)
			} else {
				fmt.Fprintf(w, "Unable to store KV pair. Key/Value pair is empty")
			}

		} else if strings.Compare(formName2, "Get corresponding value stored") == 0 {
			key := strings.Join(r.Form["key"], "")
			fmt.Println("key:", key)
			response := getKVPair(key)
			fmt.Fprintf(w, response)

		} else if strings.Compare(formName3, "Get all KV pairs stored") == 0 {
			response := getAllKVPair()
			fmt.Fprintf(w, response)
		}

	}
}

func main() {
	kv = consulInit()
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/index", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func consulInit() *api.KV {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return client.KV()
}

func putKVPair(key string, value string) string {
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return (string)("KV store failed")
		// panic(err)
	}
	fmt.Println("put KV called")
	response := "KV store success"
	return (string)(response)
}

func getKVPair(key string) string {
	fmt.Println("key: ", key)

	pairs, _, err := kv.List(key, nil)
	if pairs != nil || err != nil {
		pair, _, err := kv.Get(key, nil)
		if err != nil {
			// panic(err)
			response := "No KV pair found for this key"
			fmt.Println(response)
			return response
		}
		return (string)(pair.Value)
	}
	response := "NO KV pair found for this key"
	return response

}

func getAllKVPair() string {
	fmt.Println("getting all")
	key := ""
	pairs, _, err := kv.List(key, nil)
	if err != nil {
		response := "Bad"
		return response
	}
	response := ""
	index := 0
	for i := range pairs {
		p := pairs[i]
		val := (string)(p.Value)
		key := (string)(p.Key)
		addition := "Key: " + key + ", Value: " + val + "\n"
		fmt.Print(addition)
		response += addition
		index++
	}
	fmt.Println(response)
	return response
}
