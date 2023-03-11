package main

import (
	// "encoding/json"
	"reflect"
	"flag"
	"fmt"
	"os"
	// "strconv"
	// "log"
	"bytes"
	// "database/sql"
	// "io"
	// "net/http"
	"regexp"
	// "encoding/json"
	// "github.com/satori/go.uuid"
	
)

import log "github.com/sirupsen/logrus"
// import data "trustkbb.de/daosgenerate/appdata"
// import conf "trustkbb.de/daosgenerate/appconfig"
// import errmod "trustkbb.de/daosgenerate/apperror"
import data "appdata"
import conf "appconfig"
import errmod "apperror"

var Buf    bytes.Buffer
var LoggerDB *log.Logger
// var DBT dbs.DBTarget

func init() {

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	
	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		// FullTimestamp: true,
		// log.SetFormatter(&log.JSONFormatter{})
	
	})
	log.Info("Setup Configuration! DBTOOL")
	
	conf.Dbconf = conf.InitConfig("./appconfig/UrlConfig.json")
	err=conf.Dbconf.LoadConfig()
	errmod.CheckError("LoadJson Configuration",err) 
	log.Info("Starting MAIN")	
}
func main() {

	var index string
	var category string
	var count int
	
	flag.StringVar(&category, "l", "Zahlensequenz", "Load category. Default is Zufallszahlen")
	flag.StringVar(&index, "i", "S20", "Load index within the category. Default is Sequenz 1..10")
	flag.IntVar(&count, "c", 1, "Load count within the category. Default is Sequenz 1..10")
	flag.Parse() // after declaring flags we need to call it

	
			
	log.Debug("Category: ", category)
	log.Debug("Index: ", index)
	log.Debug("Count: ", count)

	// conf.Dbconf.ShowConfig()

	switch category {

	case "Array":

	x := [...]int{10, 20, 30}

	log.Debug("Mit Anzahl:",reflect.ValueOf(x).Kind())
	log.Debug(len(x))

	y := []int{10, 20, 30}

	log.Debug("Ohne Anzahl:",reflect.ValueOf(y).Kind())
	log.Debug(len(y))

	case "Test":

		zeichen  := []string {"Hofladen","Hof'tür","Hofstrasse"}
		re := regexp.MustCompile(`'`)
		for i, z := range zeichen {
			fmt.Printf("Vorher: %-02d %s\n",i,z)
		    fmt.Println(re.FindString(z))
			
			fmt.Println("Ausgabe",re.ReplaceAllString(z, "\\'"))
			z= re.ReplaceAllString(z, "\\'")
			fmt.Printf("Nachher: %-02d %s\n",i,z)
			
			z= re.FindString(z)
			if z != "" {
				fmt.Println("Match ...")
			}	
		}

	case "Persons":
		{
			// var Persons data.Persons
			// Persons := data.CreatePersons()
			// Persons.CreateNPerson(conf.Dbconf, int(count))
			Persons.CreateNObjects(int(count))
			// data.Report(Persons)
			err := Persons.StoreDB()
			errmod.CheckError("Store Persons to DB",err)	
		}
	
		case "Adressen":
		{
			// body := data.GetCategoryElements(category, index, conf.Dbconf)
			err:=conf.Dbconf.CheckConfig(category, index)
			errmod.CheckError("Konfiguration Adressen",err)
		
			var adressen data.Adressen
			adressen.LoadRandObjects(index)
			data.Report(adressen)
		}
	case "Zahlensequenz":
		{
			err:=conf.Dbconf.CheckConfig(category, index)
			errmod.CheckError("Konfiguration Zahlensequenz",err)
			var seq data.Zahlensequenzen
			aseq := &seq
			aseq.LoadRandObjects(index)
			data.Report(seq)
		}
	case "Namen":
		{
			err:=conf.Dbconf.CheckConfig(category, index)
			errmod.CheckError("Konfiguration Namen",err)
			var ns data.Names
			ns.LoadRandObjects(index)
			data.Report(ns)
		}

	case "SocialNumbers":
		{
			err:=conf.Dbconf.CheckConfig(category, index)
			errmod.CheckError("Konfiguration Socialnumbers",err)
			var sns data.SocialNumbers
			sns.LoadRandSocialObjects(index)
			data.Report(sns)
		}

	default:
		err:=conf.Dbconf.CheckConfig(category, index)
		errmod.CheckError("Default: fehlende Kategorie",err)
		// body := net.GetCategoryElements(category, index, conf.Dbconf)
		// fmt.Printf("Default:\n%s\n", body)
	}
	os.Exit(0)
}
