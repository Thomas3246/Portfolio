package handle

import (
	"html/template"
	"net/http"
	"portfolio/src/project"
	"portfolio/src/service"
)

func PortfolioHandler(writer http.ResponseWriter, request *http.Request) {
	htmlFile, err := template.ParseFiles("templates/mainView.html")
	service.Check(err)

	database := service.OpenDB("database.db")
	projectList := project.GetProjects(database)

	err = htmlFile.Execute(writer, projectList)
	service.Check(err)
}

func ProjectConverterHandler(writer http.ResponseWriter, request *http.Request) {
	htmlFile, err := template.ParseFiles("templates/converterView.html")
	service.Check(err)

	err = htmlFile.Execute(writer, nil)
	service.Check(err)
}
