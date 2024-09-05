package handle

import (
	"html/template"
	"net/http"
	"portfolio/src/service"
)

func PortfolioHandler(writer http.ResponseWriter, request *http.Request) {
	htmlFile, err := template.ParseFiles("templates/mainView.html")
	service.Check(err)

	err = htmlFile.Execute(writer, nil)
	service.Check(err)
}
