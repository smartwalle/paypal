package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/paypal", func(writer http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		var b, _ = ioutil.ReadAll(req.Body)
		fmt.Println(string(b))
	})

	http.ListenAndServe(":6565", nil)
}
