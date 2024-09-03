package main

import (
	"fmt"
	"net/http"

	"portfolio/src/handle"
	"portfolio/src/project"
	"portfolio/src/service"
)

func main() {
	database := service.OpenDB("database.db")
	projectList := project.GetProjects(database)
	fmt.Println(projectList)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/portfolio", handle.PortfolioHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	service.Check(err)

}
