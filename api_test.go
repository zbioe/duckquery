package duckquery

import (
	"testing"
	"bytes"
	"encoding/json"
	"github.com/iuryfukuda/duckquery"
)

func TestApi(t *testing.T) {
	o, err := duckquery.Request("eduardo campos")
	if err != nil {
		t.Fatalf("expected nil err: %#v", err)
	}
	b, err := json.Marshal(o)
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	t.Logf("output: %s", out.String())
}
