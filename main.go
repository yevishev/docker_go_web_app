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

	router.HandleFunc("/hell1o/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["name"]

		fmt.Fprintf(w, "<h1>Hey hey, %s!\n</h1>", title)
	})

	router.HandleFunc("/excel", excelPost).Methods("POST")
	router.HandleFunc("/excel", excelGet).Methods("GET")


	http.ListenAndServe(":80", router)
}

func excelGet(w http.ResponseWriter, r *http.Request) {
	var res = make(map[string]string)
	var status = http.StatusOK

	params := r.URL.Query()
	param, ok := params["title"]
	if !ok {
		res["result"] = "fail"
		res["error"] = "required parameter is not defined"
		status = http.StatusBadRequest
	} else {
		res["result"] = "ok"
		res["title"] = param[0]
		status = http.StatusOK
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func excelPost(w http.ResponseWriter, r *http.Request) {
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