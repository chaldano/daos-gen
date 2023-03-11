package appdata

import (
	"fmt"
	// "bytes"
	"crypto/x509/pkix"
	// "reflect"
	// "regexp"
)

import "github.com/icrowley/fake"
	
import errmod "trustkbb.de/daosgenerate/apperror"
import conf "trustkbb.de/daosgenerate/appconfig"
// import dbs "trustkbb.de/daosgenerate/appstorage"
import pki "trustkbb.de/daosgenerate/apppki"
import log "github.com/sirupsen/logrus"

// Methods CreatedObjects
func init() {
	log.SetLevel(log.DebugLevel)
	log.Info("Setup Config Data! (Init DataObjects)")
	path := "./../appconfig/UrlConfig.json"
	conf.Dbconf = conf.InitConfig(path)
	err := conf.Dbconf.LoadConfig()
	errmod.CheckError("Init Config", err)
}

func CreateNObjects(cat string, anz int, r interface{}) (CreatedObjects, error) {
	// anz: Anzahl der Objekte
	// r: RandomObjects für Erweiterungen
	var co CreatedObjects
	var err error
	var adresses, names, dns RandomObjects
	var registernumbers, personnumbers, emails CreatedObjects

	co.Category = cat
	co.Items = make([]interface{}, anz)
	// fmt.Println("Länge:",len(co.Items), cat)
	switch cat {
	case "Entity":
		// n := fmt.Sprintf("%02d", anz)
		uuids, err := LoadRObjects("UUID", anz, "")
		if err != nil {
			return co, err
		}
		// Generiere eine Entität mit Option
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "Entities"))
		for i := 0; i < anz; i++ {
			e := new(Entity)
			e.Type = "Entity"
			// o := new(Option)
			// e.Options = o
			// e.Options.Legal = false

			puuid := uuids.Items[i].(Uuid)
			e.Id = &puuid
			co.Items[i] = *e
		}
		return co, nil
	case "Person":
		n := fmt.Sprintf("%02d", anz)
		Entities, err := CreateNObjects("Entity", anz, "")
		if err != nil {
			return co, err
		}
		index := "A" + n
		adresses, err = LoadRObjects("Adressen", anz, index)
		if err != nil {
			return co, err
		}
		index = "N" + n
		names, err = LoadRObjects("Namen", anz, index)
		if err != nil {
			return co, err
		}
		// Generiere Emails für jeden Namen
		emails, err = CreateNObjects("Email", anz, names)
		if err != nil {
			return co, err
		}
		// Generiere eine Entität mit Option
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "Entities"))
		for i := 0; i < anz; i++ {
			p := new(Person)
			o := new(Option)
			p.Options = o
			p.Options.Legal = false
			// puuid := uuids.Items[i].(Uuid)
			// e.Id = &puuid
			Entity := Entities.Items[i].(Entity)
			p.Entity = &Entity
			name := names.Items[i].(Name)
			p.Name = &name
			email := emails.Items[i].(Email)
			p.Email = &email
			adresse := adresses.Items[i].(Adresse)
			p.Adresse = &adresse
			p.Adresse.Id = *Entity.Id
			co.Items[i] = *p
		}
		return co, nil
	case "PersonsNat":
		// n := fmt.Sprintf("%02d", anz)
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "Entities for PersonsNat"))
		// Entities, err := CreateNObjects("Entities", anz, nil)
		// if err != nil {
		// 	return co, err
		// }

		// Generiere Persons
		Persons, err := CreateNObjects("Person", anz, nil)
		if err != nil {
			return co, err
		}
		// Generiere Socialnumbers aus Zeichenketten
		personnumbers, err = CreateNObjects("PersonNumbers", anz, nil)
		if err != nil {
			return co, err
		}

		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "PersonsNat"))
		for i, _ := range Persons.Items {
			PersonNat := new(PersonNat)
			Person := Persons.Items[i].(Person)
			Person.Entity.Type = "NatürlichePerson"
			Person.Options.Legal = false

			sns := personnumbers.Items[i].(PersonNumber)
			PersonNat.Person = &Person
			PersonNat.Number = &sns
			co.Items[i] = *PersonNat
		}
		return co, nil
	case "PersonsJur":
		// n := fmt.Sprintf("%02d", anz)
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "Entities for PersonsJur"))
		// Entities, err := CreateNObjects("Entities", anz, nil)
		// if err != nil {
		// 	return co, err
		// }
		// Generiere Persons
		Persons, err := CreateNObjects("Person", anz, nil)
		if err != nil {
			return co, err
		}
		// Generiere Registrynumbers aus Zeichenketten
		registernumbers, err = CreateNObjects("RegisterNumbers", anz, nil)
		if err != nil {
			return co, err
		}

		// var Persons Persons
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "PersonsJur"))
		for i, _ := range Persons.Items {
			PersonJur := new(PersonJur)
			Person := Persons.Items[i].(Person)
			Person.Entity.Type = "JuristischePerson"
			Person.Options.Legal = true

			// entity := e.(Entity)
			// Person.Entity = &entity
			// Person.Entity.Type = "Person"
			// Person.Entity.Options = e.(Entity).Options
			// Person.Entity.Options.Legal = true

			sns := registernumbers.Items[i].(RegisterNumber)
			PersonJur.Person = &Person
			PersonJur.Number = &sns
			PersonJur.Rechtsform = "AG"
			PersonJur.Company = fake.Company()
			co.Items[i] = *PersonJur
		}
		return co, nil
	case "RegisterNumbers":
		n := fmt.Sprintf("%02d", anz)
		index := "Z" + n
		zks, err := LoadRObjects("Zeichenketten", anz, index)
		if err != nil {
			return co, err
		}
		for i, zk := range zks.Items {
			var sn = new(RegisterNumber)
			sn.Country = "US"
			sn.RId = "HRB:" + zk.(string)
			co.Items[i] = *sn
		}
		return co, nil
	case "PersonNumbers":
		n := fmt.Sprintf("%02d", anz)
		index := "Z" + n
		zks, err := LoadRObjects("Zeichenketten", anz, index)
		if err != nil {
			return co, err
		}
		for i, zk := range zks.Items {
			var sn = new(PersonNumber)
			// sn.Country = "US"
			sn.RId = "PKZ:" + zk.(string)
			co.Items[i] = *sn
		}
		return co, nil
	case "CA":
		Entities, err := CreateNObjects("Entity", anz, nil)
		if err != nil {
			return co, err
		}
		Persons, err := CreateNObjects("PersonsJur", anz, nil)
		if err != nil {
			return co, err
		}

		for i, _ := range Entities.Items {
			ca := new(CA)
			// Entity := e.(Entity)
			PersonJur := Persons.Items[i].(PersonJur)
			ca.PersonJur = &PersonJur
			// ca.Entity = &entity
			ca.PersonJur.Person.Entity.Type = "CA"

			// ca.Entity.Person.Options = e.(Entity).Options
			// ca.Entity.Person.Options.Legal = true

			ca.PersonJur.Person.Name.Firstname = fmt.Sprintf("CA%02d", i)
			ca.PersonJur.Person.Name.Surname = PersonJur.Person.Adresse.Bundesland

			// CA-Zertificat erzeugen
			err = ca.getCert()
			if err != nil {
				return co, err
			}
			// co.Items[i] = ca
			co.Items[i] = *ca
		}
		return co, nil
	case "Email":
		n := fmt.Sprintf("%02d", anz)
		dns, err = LoadRObjects("DomainNames", anz, "")
		if err != nil {
			return co, err
		}
		switch v := r.(type) {
		case RandomObjects:
			fmt.Printf("(%s von %s)\n", "Externe Random-Namen", v)
			names = r.(RandomObjects)
		default:
			fmt.Printf("(%s)\n", "Interne Random-Namen")
			index := "N" + n
			names, err = LoadRObjects("Namen", anz, index)
			if err != nil {
				return co, err
			}
		}
		log.Debug(fmt.Sprintf("%-6s %s %03d %-10s\n", "Create", ":", anz, "Email-DNS"))
		for i := 0; i < anz; i++ {
			Email := new(Email)
			dns := dns.Items[i].(DomainName)
			// Email.Domain = dns
			name := names.Items[i].(Name)
			// Email.Name = name
			Email.Adresse = fmt.Sprintf("%s.%s@%s", name.Firstname, name.Surname, dns.name)
			co.Items[i] = *Email
		}
		return co, nil

	default:
		// fmt.Println("Durch")
		err = errmod.OopsError("Default: Wrong ObjectCategory in CreateNObjects")
		return co, err
	}
	return co, err
}

func (co CreatedObjects) Listitems() {
	fmt.Println("\nReport Kategorie: ", co.Category)
	for i, s := range co.Items {
		fmt.Println("Category:", co.Category)
		switch co.Category {
		case "RegisterNumbers":
			// for i, s := range co.Items {
			fmt.Printf("%02d : %-4s %-8s \n", i, s.(RegisterNumber).Country, s.(RegisterNumber).RId)

		case "PersonsNat", "PersonsJur":
			rechtsform := "Natürliche"
			var Entity *Entity
			var Person *Person
			if co.Category == "PersonsNat" {
				Entity = s.(PersonNat).Person.Entity
				Person = s.(PersonNat).Person
				Entity = s.(PersonNat).Person.Entity
				fmt.Printf("%02d: %-8s  Person\n", i, rechtsform)
			} else {
				rechtsform = "Juristische"
				Entity = s.(PersonJur).Person.Entity
				Person = s.(PersonJur).Person
				Company := s.(PersonJur).Company
				fmt.Printf("%02d: %-8s Person (%s): ", i, rechtsform, s.(PersonJur).Rechtsform)
				fmt.Printf("Company (%s) \n", Company)
			}
			fmt.Printf("%-15s:%-5s\n", "UUID", Entity.Id.UId.String())
			// Differenzierte Ausgabe HRB und PKZ
			if co.Category == "PersonsNat" {
				Number := s.(PersonNat).Number
				Country:= s.(PersonNat).Person.Adresse.Land
				fmt.Printf("%-15s:%s (Country:%s)\n", "PKZNumber", Number.RId, Country)
				fmt.Printf("%-15s:%-5s %-8s\n", "Name", Person.Name.Firstname, Person.Name.Surname)
			
			} else {
				Number := s.(PersonJur).Number
				Country:= s.(PersonJur).Person.Adresse.Land

				fmt.Printf("%-15s:%s (Country:%s)\n", "HRBNumber", Number.RId, Country)
				fmt.Printf("%-15s:%-5s %-8s\n", "Inhaber", Person.Name.Firstname, Person.Name.Surname)
			
				//  fmt.Printf("%-15s:%s\n\n", "Rechtsform", Number.RId, Number.Country)
			}
			// Ausgabe der Entity-Daten
			fmt.Printf("%-15s:%-5s\n", "EMail", Person.Email.Adresse)
			fmt.Printf("%-15s:%-5s\n", "Plz", Person.Adresse.PLZ)
			fmt.Printf("%-15s:%-10s\n", "Ort", Person.Adresse.Ort)
			fmt.Printf("%-15s:%-12s\n", "Strasse", Person.Adresse.Strasse)
			fmt.Printf("%-15s:%-12s\n", "Bundesland", Person.Adresse.Bundesland)
			fmt.Printf("%-15s:%-12s\n\n", "Land", Person.Adresse.Land)
			if len(Entity.Relations) > 0 {
				fmt.Printf("%-15sAnzahl:%-12d\n", "Existierende Beziehungen:", len(Entity.Relations))
				for i, relation := range Entity.Relations {
					Person := co.Items[relation].(PersonNat).Person
					Entity := Person.Entity
					if i < len(Entity.Relations)-1 {
						fmt.Printf("%-2d: RelationTo->(%-20s) %-10s\n", i, Entity.Id.UId,Person.Name.Surname)
					} else {
						fmt.Printf("%-2d: RelationTo->(%-20s) %-10s\n\n", i, Entity.Id.UId,Person.Name.Surname)
					}
				}
			}
			if Person.Options.CertBlock != nil {
				fmt.Println("Certificate available")
				fmt.Printf("%-15s:%-12v\n", "Certificate-SN", Person.Options.CertBlock.Cert.SerialNumber)
				fmt.Printf("%-15s:%-12v\n", "Subject", Person.Options.CertBlock.Cert.Subject)
				fmt.Printf("%-15s:%-12v\n", "Issuer", Person.Options.CertBlock.Cert.Issuer)

			}
		case "CA":

			rechtsform := "Natürliche"
			if s.(CA).PersonJur.Person.Options.Legal == true {
				rechtsform = "Juristische"
			}
			fmt.Printf("%02d: %-8s CA\n", i, rechtsform)
			fmt.Printf("%-12s %-5s %-8s\n", "Name", s.(CA).PersonJur.Person.Name.Firstname, s.(CA).PersonJur.Person.Name.Surname)
			fmt.Printf("%-12s:%-5s\n", "CAID", s.(CA).PersonJur.Person.Entity.Id.UId.String())
			// fmt.Printf("%-12s:%-5s\n", "EMail", s.(Person).Email.Adresse)
			// fmt.Printf("%-12s:%-5s (Country:%-10s)\n", "SocialNumber", s.(CA).Number.SId, s.(Person).Number.Country)
			fmt.Printf("%-12s:%-5s\n", "Plz", s.(CA).PersonJur.Person.Adresse.PLZ)
			fmt.Printf("%-12s:%-10s\n", "Ort", s.(CA).PersonJur.Person.Adresse.Ort)
			fmt.Printf("%-12s:%-12s\n", "Strasse", s.(CA).PersonJur.Person.Adresse.Strasse)
			fmt.Printf("%-12s:%-12s\n", "Bundesland", s.(CA).PersonJur.Person.Adresse.Bundesland)
			fmt.Printf("%-12s:%-12s\n\n", "Land", s.(CA).PersonJur.Person.Adresse.Land)
			if s.(CA).PersonJur.Person.Options.CertBlock != nil {
				fmt.Println("Certificate available")
				fmt.Printf("%-12s:%-12v\n", "Certificate-SN", s.(CA).PersonJur.Person.Options.CertBlock.Cert.SerialNumber)
				fmt.Printf("%-12s:%-12v\n", "Subject", s.(CA).PersonJur.Person.Options.CertBlock.Cert.Subject)
				fmt.Printf("%-12s:%-12v\n", "Issuer", s.(CA).PersonJur.Person.Options.CertBlock.Cert.Issuer)

				// fmt.Printf("%-12s:%-12v\n\n", "Certificate_PEM", s.(CA).Entity.Options.CertBlock.PemCert)
				// fmt.Printf("%-12s:%-12v\n\n", "Certificate_KeyPEM", s.(CA).Entity.Options.CertBlock.PemPrivKey)

			}
		case "Email":
			fmt.Printf("%-8s(%02d)\n", "Email", i)
			fmt.Printf("%-8s:%-12s\n\n", "Email", s.(Email).Adresse)
		case "Relation":
			fmt.Printf("%02d:%-8s%-12s:\n", i, "Type:", s.(Relation).Type)
			fmt.Printf("%-8s(%-12s)\n", "Source:", s.(Relation).RSource.UId.String())
			fmt.Printf("%-8s(%-12s)\n\n", "Target:", s.(Relation).RTarget.UId.String())
		}
	}
}
func (ca *CA) getCert() error {
	var Name pkix.Name
	var Issuer pkix.Name

	organisation := fmt.Sprintf("%s(%s)-%s", ca.PersonJur.Person.Name.Firstname, "Authority", ca.PersonJur.Person.Name.Surname)
	Name.Organization = []string{organisation}
	Name.Country = []string{ca.PersonJur.Person.Adresse.Land}
	Name.Province = []string{""}
	Name.Locality = []string{ca.PersonJur.Person.Adresse.Ort}
	Name.StreetAddress = []string{ca.PersonJur.Person.Adresse.Strasse}
	Name.PostalCode = []string{ca.PersonJur.Person.Adresse.PLZ}
	Name.CommonName = fmt.Sprintf("%s", ca.PersonJur.Person.Name.Firstname)

	Issuer.Organization = Name.Organization

	cb, err := pki.CreateCACertificate(Name, Issuer)
	if err != nil {
		return err
	}
	ca.PersonJur.Person.Options.CertBlock = cb
	return nil
}
func (ca *CA) issueCert(p *PersonNat) error {
	var Name pkix.Name
	var Issuer pkix.Name

	Name.Organization = []string{p.Person.Name.Firstname}
	Name.Country = []string{p.Person.Adresse.Land}
	Name.Province = []string{""}
	Name.Locality = []string{p.Person.Adresse.Ort}
	Name.StreetAddress = []string{p.Person.Adresse.Strasse}
	Name.PostalCode = []string{p.Person.Adresse.PLZ}
	Name.CommonName = fmt.Sprintf("%s.%s", p.Person.Name.Firstname, p.Person.Name.Surname)

	organisation := fmt.Sprintf("%s(%s)-%s", ca.PersonJur.Person.Name.Firstname, "Authority", ca.PersonJur.Person.Name.Surname)
	Issuer.Organization = []string{organisation}

	CACertblock := ca.PersonJur.Person.Options.CertBlock
	cb, err := pki.CreateUserCertificate(Name, Issuer, CACertblock)
	if err != nil {
		return err
	}
	p.Person.Options.CertBlock = cb
	return nil
}
func (co CreatedObjects) CreateRelations(relanz int) (CreatedObjects, error) {
	// =============================================
	// Input: 
	// anz: Anzahl der Objekte
	// relanz: Beziehungen
	// =============================================
	// Output:
	// Array of Relations (source) => (target)
	// =============================================
	
	var err error
	anz := len(co.Items)
	if relanz >= anz {
		relanz = anz - 1
	}
	var Relations CreatedObjects
	count := anz * relanz
	Relations.Items = make([]interface{}, count)
	cat := "Relation"
	Relations.Category = cat

	fmt.Println("Array:", count)
	targets := make([]int, relanz)
	i := 0

	for k := 0; k < len(co.Items); k++ {
		// fmt.Println("Person: ", k)
		targets = make([]int, relanz)
		// Init targets with -1
		for i, _ := range targets {
			targets[i] = -1
		}
		for j := 0; j < len(targets); j++ {
			r := new(Relation)
			r.Type = co.Category
			r.RSource = co.Items[k].(PersonNat).Person.Entity.Id

			// Erzeugt Zufallswerte bis antrel - Relations gefunden
			for {
				number, err := RandDigitFrom(anz)
				if err != nil {
					err = errmod.OopsError("Wrong random count for relations")
					return Relations, err
				}
				// fmt.Println("Targets: ", targets)
				// fmt.Println("Number: ", number)

				if number != k && !Contains(targets, number) {
					targets[j] = number
					// co.Items[k].(PersonNat).Entity.Targets[j] = number
					r.RTarget = co.Items[number].(PersonNat).Person.Entity.Id
					err = nil
					Relations.Items[i+j] = *r
					break
				}
			}
		}
		// fmt.Println("raus")
		i = i + relanz
		// ResetValues(co.Items[k].(PersonNat).Entity.Targets)
	}

	//error = nil
	return Relations, err
}
func (co CreatedObjects) CreateRelationsByObject(relanz int) error {
	// =============================================
	// Input: 
	// anz: Anzahl der Objekte
	// relanz: Relations
    // =============================================
	// Output:
	// Relations are bound to Entity Relations-Array 
	// =============================================
	
	var err error
	anz := len(co.Items)
	// fmt.Println("Anzahl: ", anz)
	if relanz >= anz {
		relanz = anz - 1
	}
	i := 0
	// Durchlaufe jedes Objekt und baue dafür zufällige Beziehungen aus
	for k := 0; k < len(co.Items); k++ {
		targets := make([]int, relanz)
		// Init targets with -1, da auch 0 ein Target ist
		for i, _ := range targets {
			targets[i] = -1
		}
		// Ausgewähltes Object : j zufällige Beziehungen herstellen
		for j := 0; j < len(targets); j++ {
			// Erzeugt Zufallswerte bis anzrel - Relations gefunden
			for {
				number, err := RandDigitFrom(anz)
				if err != nil {
					err = errmod.OopsError("Wrong random count for relations")
					return err
				}
				// fmt.Println("Targets: ", targets)
				// fmt.Println("Number: ", number)

				if number != k && !Contains(targets, number) {
					targets[j] = number
					err = nil
					break
				}
			}
		}
		// Beziehungen für Items[i] liegen als Integerwerte in target[] vor
		// Integer-Werte auf UUID umschreiben
		// fmt.Println("Umschreiben auf Index")
		switch co.Category {
		case "PersonsNat":
			co.Items[k].(PersonNat).Person.Entity.Relations = make([]int, relanz)
			for r, relationnumber := range targets {
				co.Items[k].(PersonNat).Person.Entity.Relations[r] = relationnumber
				// UId.String()
			}
		}
		// Beziehungen für nächstes Object
		i = i + relanz
		// ResetValues(co.Items[k].(PersonNat).Entity.Targets)
	}

	// err := nil
	return err
}

// Durchsucht Slice nach Wert
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
