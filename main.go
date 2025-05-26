package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(ResponseWriter, *Request)

func main() {
	http.ListenAndServe(":8080")
}