package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Header struct {
	Coordinate string `json:"c"`
	Text string `json:"v"`
	Width int `json:"w"`
}

type Json struct {
	Headers []Header `json:"h"`
}

func main() {
	router := mux.NewRouter()

	/* Homepage */
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage. Try /excel and set .json body\n</h1>")
	})

	/* Post request test */
	router.HandleFunc("/excelp", excelPostTest).Methods("POST")

	/* Get request test */
	router.HandleFunc("/excelg", excelGetTest).Methods("GET")

	/* Get request test */
	router.HandleFunc("/excel", excelPost).Methods("POST")

	/* Specifying ports */
	http.ListenAndServe(":8002", router)
}

func excelGetTest(w http.ResponseWriter, r *http.Request) {
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

func excelPostTest(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := req["id"].(string)
	title := req["title"].(string)

	//response
	var res = map[string]string{
		"id" : id,
		"title" : title,
	}
	fmt.Println("Post id: " ,id)

	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func excelPost(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	j := Json{}
	err := json.Unmarshal(body, &j)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range j.Headers {
		fmt.Printf("%v\n", item.Coordinate)
	}
}