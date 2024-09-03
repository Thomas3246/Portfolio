package project

import (
	"database/sql"
	"portfolio/src/service"
)

type Project struct {
	ID          int
	Name        string
	Descr       string
	LangImgPath string
}

// func GetProjects(database *sql.DB) []Project
// GetProjects makes a query to the database and returns a slice of Project type
// Project type struct:
// ID          int
// Name        string
// Descr       string
// LangImgPath string
func GetProjects(database *sql.DB) []Project {
	rows, err := database.Query(`SELECT P.ProjectID, P.ProjectName, P.ProjectDescr, L.ImagePath
								FROM Projects as P JOIN Languages as L
								ON P.LangID = L.ID `)
	service.Check(err)

	var projectList []Project
	var project Project
	for rows.Next() {
		rows.Scan(&project.ID, &project.Name, &project.Descr, &project.LangImgPath)
		projectList = append(projectList, project)
	}
	return (projectList)
}
