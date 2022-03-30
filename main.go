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
		fmt.Fprintf(w, "<h1>Pong. Try /excel and set .json body\n</h1>")
	})

	/* Get request test */
	router.HandleFunc("/excel", excelPost).Methods("POST")

	/* Specifying ports */
	http.ListenAndServe(":8001", router)
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

	var formattedTime = time.Now().Format("2.01.06 15:04")
	var filename  = "runtime/" + formattedTime + ".xlsx"
	err = f.SaveAs(filename)
	if err != nil {
		fmt.Println(err, 1)
		return
	}
	te := time.Now()
	fmt.Fprintf(w, "<h1>" + te.Sub(ts).String() + "</h1>")

}