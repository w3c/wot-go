package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("www/"))
	http.ListenAndServe(":8000", fileServer)
}