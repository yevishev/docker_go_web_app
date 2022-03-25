package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/Sammy\n</h1>")
	})

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker!\n</h1>")
	})

	router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["name"]

		fmt.Fprintf(w, "<h1>Hey hey, %s!\n</h1>", title)
	})

	router.HandleFunc("/excel", MyExcel).Methods("POST")


	http.ListenAndServe(":80", router)
}

func MyExcel(w http.ResponseWriter, r *http.Request) {
	//request handling
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	id := req["id"].(string)
	title := req["title"].(string)

	//response
	var res = map[string]string{
		"id" : strconv.Itoa(rand.Intn(10000000)),
		"title" : title,
	}
	fmt.Println("Post id: " ,id)

	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}