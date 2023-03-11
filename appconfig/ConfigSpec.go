package appconfig

import (
	// "fmt"
	)
// import tools "trustkbb.de/tools/dbtools/common"

type Ressource struct {
	Cat	string				// Kategorie
	Url  string				// Parametrisierte Url
	Name string				// Url-Name innerhalb einer Kategorie
}

type Ressources struct {
	Configs []Ressource
} 


type DBConfigs struct{
	Path	string
	Urls 	*Ressources
}

type SqlRequest struct {
	elements []string
	table string
	condition string
}
