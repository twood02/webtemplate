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

func setup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments entered in form
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	// get inputs entered by user
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}
func processForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// determine which action to perform
		formName := strings.Join(r.Form["store-submit"], "")
		formName2 := strings.Join(r.Form["getval-submit"], "")
		formName3 := strings.Join(r.Form["getall-submit"], "")
		strings.Compare(formName, "Store KV pair")
		strings.Compare(formName2, "Get corresponding value stored")
		strings.Compare(formName3, "Get all KV pairs stored")

		// store KV pair
		if strings.Compare(formName, "Store KV pair") == 0 {
			key := strings.Join(r.Form["key"], "")
			value := strings.Join(r.Form["value"], "")
			response := putKVPair(key, value)
			if strings.Compare(response, "KV store success") == 0 {
				fmt.Fprintf(w, response+"\n"+"Key: "+key+"  Value: "+value)
			} else {
				fmt.Fprintf(w, "Unable to store KV pair. Key/Value pair is empty")
			}

		} else if strings.Compare(formName2, "Get corresponding value stored") == 0 { // get value given key
			key := strings.Join(r.Form["key"], "")
			response := getKVPair(key)
			if strings.Compare(response, "No KV pair found for this key") == 0 {
				fmt.Fprintf(w, "No pair found. Emapty key entered in search.")
			} else if strings.Compare(response, "No pair found for this key") == 0 {
				fmt.Fprintf(w, "No pair found for key '"+key+"'.")
			} else {
				fmt.Fprintf(w, "Value found key '"+key+"': "+response)
			}

		} else if strings.Compare(formName3, "Get all KV pairs stored") == 0 { //get all values store by node
			response := getAllKVPair()
			if strings.Compare(response, "Bad") == 0 {
				fmt.Fprintf(w, "Unable to get all KV pairs stored")
			} else {
				fmt.Fprintf(w, "Displaying all KV pairs stored. \n\n"+response)
			}
		}

	}
}

func main() {
	kv = consulInit()
	http.HandleFunc("/", setup)              // setting router rule
	http.HandleFunc("/index", processForm)   // connection to web interface/form
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// initialize KV struct
func consulInit() *api.KV {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return client.KV()
}

// add KV pair
func putKVPair(key string, value string) string {
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return (string)("KV store failed")
	}
	response := "KV store success"
	return (string)(response)
}

// GET value given pair
func getKVPair(key string) string {

	pairs, _, err := kv.List(key, nil)
	if pairs != nil || err != nil {
		pair, _, err := kv.Get(key, nil)
		if err != nil {
			response := "No KV pair found for this key"
			return response
		}
		return (string)(pair.Value)
	}
	response := "No pair found for this key"
	return response

}

// GET all pairs
func getAllKVPair() string {
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
		response += addition
		index++
	}
	return response
}
