package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type env struct {
	Foo string `json: foo`
	Bar string `json: bar`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	f := os.Getenv("FOO")
	b := os.Getenv("BAR")

	e := env{
		Foo: f,
		Bar: b,
	}
	json.NewEncoder(w).Encode(e)

}
func file(w http.ResponseWriter) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, dir)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
