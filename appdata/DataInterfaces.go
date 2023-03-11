package appdata

import (
	// "database/sql"
	//"fmt"
	// _ "github.com/lib/pq"
)

type ReportCategory interface {
    // GetCategory() string
    Listitems()                 // List all items of a selected categorie
}

type RandomObjectMethods interface {
    LoadRandObjects (index string)
}

// Funktionen mit der Objektkonfiguration
type ObjectConfig interface {
    ShowConfig()
    CheckConfig()
}

