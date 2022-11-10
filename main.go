package main

import (
	"net/http"
)

func main() {
	http.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "https://sayi.do/lizzyandjames")
		w.WriteHeader(301)
	}))

	http.ListenAndServe("", nil)
}
