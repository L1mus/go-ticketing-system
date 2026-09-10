package main

import (
	"fmt"
	"net/http"
)

func helloHundler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello! Ticketing System")
}

func healtCheck(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	responseMessage := fmt.Sprintf(`{"status" : "%s"}`, status)

	fmt.Fprintln(w, responseMessage)
}

func sayHelloTo(w http.ResponseWriter, r *http.Request) {
	// q := r.FormValue("name")
	q := r.URL.Query()
	name := q.Get("name")
	fmt.Println(name)
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	http.HandleFunc("/", helloHundler)
	http.HandleFunc("/health", healtCheck)
	http.HandleFunc("/hello", sayHelloTo)

	fmt.Println("Server running on http://localhost:8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
