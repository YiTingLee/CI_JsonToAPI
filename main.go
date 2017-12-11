package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Serve int

func (m Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadFile("data.json")
	if err != nil {

	}
	fmt.Fprintln(w, string(res))
}

func main() {
	var s Serve
	http.ListenAndServe(":8080", s)
}
