package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"time"
)

type Header struct {
	Coordinate string `json:"c"`
	Text string `json:"n"`
	Width int `json:"w"`
}

type Cell struct {
	Coordinate string `json:"s"`
	Value string `json:"v"`
}

type Title struct {
	Name string `json:"name"`
}

type Json struct {
	Headers []Header `json:"h"`
	Companies   map[string][]Cell `json:"d"`
	Title   `json:"t"`
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
	ts := time.Now()

	body, _ := ioutil.ReadAll(r.Body)
	j := Json{}
	err := json.Unmarshal(body, &j)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := j.Title

	f := excelize.NewFile()
	index := f.NewSheet(title.Name)
	tp := time.Now()
	fmt.Println(tp.Sub(ts).String())
	for _, header := range j.Headers {
		err := f.SetCellValue(title.Name, header.Coordinate, header.Text)
		if err != nil {
			fmt.Println(err, 1)
			return
		}
	}

	for _, company := range j.Companies {
		for _, data := range company {
			err := f.SetCellValue(title.Name, data.Coordinate, data.Value)
			if err != nil {
				fmt.Println(err, 1)
				return
			}

		}
	}

	f.SetActiveSheet(index)

	var formattedTime = time.Now().Format("2-01-06 15-04")
	var filename  = "runtime/" + formattedTime + ".xlsx"
	err = f.SaveAs(filename)
	if err != nil {
		fmt.Println(err, 1)
		return
	}
	te := time.Now()
	fmt.Fprintf(w, "<h1>" + te.Sub(ts).String() + "</h1>")

}