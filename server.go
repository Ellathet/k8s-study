package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/foods", ConfigMap)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	fmt.Fprint(w, "Hello, I'm %s.", name)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/foods/foods.txt")
	if err != nil {
		log.Fatalf("error reading file", err)
	}

	fmt.Fprint(w, "My foods are: %s.", string(data))
}
