package main

import (
	"net/http"

	"portfolio/src/handle"
	"portfolio/src/service"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/portfolio", handle.PortfolioHandler)
	http.HandleFunc("/portfolio/converter", handle.ProjectConverterHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	service.Check(err)
}
