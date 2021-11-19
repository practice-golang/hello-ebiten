package main // import "server"

import (
	"net/http"
)

func main() {
	http.ListenAndServe("127.0.0.1:8864", http.FileServer(http.Dir(".")))
}
