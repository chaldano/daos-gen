package appdata

// import "bytes"
// import "crypto/x509"
// import "crypto/x509/pkix"

import "github.com/satori/go.uuid"

// func init() {
// 	fmt.Println("Database-Prgramms -This will get called on main initialization")
//   }
import pki "trustkbb.de/daosgenerate/apppki"

// type Entity struct{
// 	Name string
// 	Certname string
// }

type RandomObjects struct {
	Category string
	Items    []interface{}
}
type CreatedObjects struct {
	Category string
	Items    []interface{}
}

// Natürliche oder juristische Person
type Entity struct {
	Id      	*Uuid
	Type    	string
	Relations 	[] int // Index of Entity
}

type Person struct {
	Entity  *Entity
	Name    *Name
	Adresse *Adresse
	Email   *Email
	Options *Option
}

// Juristische Person (Company)
type PersonJur struct {
	Person     *Person
	Number     *RegisterNumber
	Company	string
	Rechtsform string
}

// Natürliche Person (Benutzer)
type PersonNat struct {
	Person *Person
	Number *PersonNumber
}

// type Person struct {
//  	Entity	*Entity
//  	Number  *RegisterNumber
//     }

type CA struct {
	// Entity *Entity
	PersonJur *PersonJur
	DN     string
}

type Option struct {
	Legal     bool
	CertBlock *pki.CertBlock
}

type Uuid struct {
	UId uuid.UUID
}

// Adresse
type Adresse struct {
	Id         Uuid
	Strasse    string
	Etage      string
	PLZ        string
	Ort        string
	Bundesland string
	Land       string
}

type PersonNumber struct {
	// Country string
	RId     string
}

type RegisterNumber struct {
	Country string
	RId     string
}

type Name struct {
	Firstname string
	Surname   string
}

type DomainName struct {
	name string
}

// E-Mail
type Email struct {
	Adresse string
}

// Relation (Beziehung zwischen Entitäten)

type Relation struct {
	Type    string
	RSource *Uuid
	RTarget *Uuid
}

// Anzahl bestehender Beziehungen
type Relations struct {
	Category string
	Items    []interface{}
}
