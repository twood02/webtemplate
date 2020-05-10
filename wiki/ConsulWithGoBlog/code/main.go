package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
)

var kv *api.KV

func main() {

	kv = consulInit()

	r := mux.NewRouter()
	r.HandleFunc("/put-kv", handlePutKV).Methods("POST")
	r.HandleFunc("/get-kv/{key}", handleGetKV).Methods("GET")

	//	log.Println("Listening on: ", viper.GetString("services."+ServiceName))
	//	log.Fatal(http.ListenAndServe(":"+viper.GetString("services"+ServiceName), r))

	putKVPair("hi!", "1")
	getKVPair("hi!")
}

func consulInit() *api.KV {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return client.KV()
}

func putKVPair(key string, value string) {
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
}

func getKVPair(key string) {
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}
