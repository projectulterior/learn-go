package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		request, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=calebpark", nil)
		resp, _ := http.DefaultClient.Do(request)

		response, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, string(response))
		fmt.Fprintf(w, "Hello, Caleb!")
	})

	fmt.Println("starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
