package duckquery_test

import (
	"testing"
	"bytes"
	"encoding/json"
	"github.com/iuryfukuda/duckquery"
)

func TestDuckquery(t *testing.T) {
	o := duckquery.Query("Eduardo Campos")
	if len(o.Labels) == 0 {
		t.Fatal("not find any labels")
	}
	for _, l := range o.Labels {
		if l.Label == "" {
			t.Fatal("not find label in label object")
		}
		if l.Value == "" {
			t.Fatal("not find value in label object")
		}
	}
	if len(o.Topics) == 0 {
		t.Fatal("not find any topic")
	}
	if o.Abstract == "" {
		t.Fatal("not find any abstract")
	}
	b, err := json.Marshal(o)
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	t.Logf("output: %s", out.String())
}
