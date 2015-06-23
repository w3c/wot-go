// this is a stepping stone to a Go equivalent of httpd.js, see
// https://github.com/w3c/web-of-things-framework/blob/master/httpd.js

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    mime_types := map[string]string {
        "html": "text/html",
        "txt": "text/plain",
        "js": "text/javascript",
        "json": "application/json",
        "css": "text/css",
        "jpeg": "image/jpeg",
        "jpg": "image/jpeg",
        "png": "image/png",
        "gif": "image/gif",
        "ico": "image/x-icon",
    }

	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}