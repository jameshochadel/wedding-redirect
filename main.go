package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "https://sayi.do/lizzyandjames")
		w.WriteHeader(301)
	}))
}
