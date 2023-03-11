package appdata

import (
	"testing"
	"fmt"
	// "github.com/holmes89/hello-api/translation"
)
// import conf "trustkbb.de/tools/dbtools/appconfig"

func TestInterface (t *testing.T) {

	// Arrange
	// var zks Zeichenketten
	//index := "Z10"

	var b Basic
	var c Car
	var p Paar
	
	b.name = "BasicName"
	c.name = "CarName"
	p.name = "PaarNachname"
	p.vorname = "PaarVorname"
	//c.vorname ="CarVorname"
	
	var l1 Listen 
	l1.name = "Liste1"
	l1.Items = make([]interface{}, 3)
	
	l1.Items[0] = b
	l1.Items[1] = p
	l1.Items[2] = c
	
	for _, l := range l1.Items {
		switch l.(type) {
		case Basic:
			fmt.Printf("%+v %T\n",l,l)
		case Car:
			fmt.Printf("%+v %T\n",l,l)
		case Paar:
			fmt.Printf("%+v %T\n",l,l)
			fmt.Printf("Vorname: %s\n",l.(Paar).vorname)
			fmt.Printf("Nachname: %s\n",l.(Paar).name)
		}
			}

	// ListTestData(b)
	// ListTestData(c.name)
	 ListTestData(p)
	//ListTestData(index)

}
