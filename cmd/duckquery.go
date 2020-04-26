package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"github.com/zbioe/duckquery"
)


type input struct {Query string}

var port = "2610"

func main() {
	http.HandleFunc("/query", handleQuery)
	log.Printf("start server in port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), nil))
}

func resp(w http.ResponseWriter, data []byte) {
	fmt.Fprintf(w, `{"status": "%d", "data": %s}`, http.StatusOK, data)
}

func respError(w http.ResponseWriter, status int, detail string) {
	fmt.Fprintf(w, `{"status": "%d", "detail": "%s"}`, status, detail)
}

func handleQuery(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		respError(w, http.StatusBadRequest, "nil body")
		return
	}
	var i = input{}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respError(w, http.StatusInternalServerError, "read from body: " + err.Error())
		return
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		respError(w, http.StatusNotAcceptable, "invalid json: " + err.Error())
		return
	}
	if i.Query == "" {
		respError(w, http.StatusNotAcceptable, "empty query")
		return
	}
	o, err := duckquery.Query(i.Query)
	if err != nil {
		respError(w, http.StatusInternalServerError, "problem with query: " + err.Error())
		return
	}

	b, err = json.Marshal(o)
	if err != nil {
		respError(w, http.StatusInternalServerError, "problem with response of query: " + err.Error())
		return
	}
	resp(w, b)
	return
}
