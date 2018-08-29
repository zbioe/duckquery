package duckquery

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const ApiURL = "https://api.duckduckgo.com/?format=json"

type apiOutput struct {
	Abstract string
	RelatedTopics []topic
	Infobox	infobox
}

type topic struct {
	Text string
}

type infobox struct {
	Content []Label
}

func Request(query string) (apiOutput, error) {
	var (
		o apiOutput
		rawurl = ApiURL + fmt.Sprintf("&q=%s", query)
	)
	resp, err := http.Get(rawurl)
	if err != nil {
		return o, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return o, err
	}
	err = json.Unmarshal(b, &o)
	switch err.(type) {
	// ignore bad types without pattern returned by api
	case *json.UnmarshalTypeError:
		err = nil
	default:
	}
	return o, err
}

