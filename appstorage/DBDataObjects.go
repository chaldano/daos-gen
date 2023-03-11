package appstorage

import (
    //"testing"
    "fmt"   
)
// import errmod "trustkbb.de/daosgenerate/apperror"
import dat "trustkbb.de/daosgenerate/appdata"
import log "github.com/sirupsen/logrus"


func StorePersonsNat(co dat.CreatedObjects, dbname string) error {
	log.Info(">> Save Persons TO Database ", dbname)
	// Speicherung der Adressen
	dbm := &DBM
	err := dbm.SetCurrentDB(dbname)
	if err != nil {
		return err
	}
	var sql SqlRequest
	// Speicherung Adressen
	sql.SetRequest(DBM.Current, "Adressen", "INSERT", "")
	var tabcolumns = []string{"ID", "STRASSE", "PLZ", "ORT", "BUNDESLAND", "LAND"}
	var tabvalues []string
	log.Debug(">> Create Adressen-Table Array")

	for _, p := range co.Items {
		tabvalues = make([]string, len(tabcolumns))
		// tabvalues[0]=strconv.Itoa(p.Adresse.Id)
		tabvalues[0] = p.(dat.PersonNat).Person.Entity.Id.UId.String()
		log.Debug("AdresseID " + tabvalues[0])
		tabvalues[1] = p.(dat.PersonNat).Person.Adresse.Strasse
		tabvalues[2] = p.(dat.PersonNat).Person.Adresse.PLZ
		tabvalues[3] = p.(dat.PersonNat).Person.Adresse.Ort
		tabvalues[4] = p.(dat.PersonNat).Person.Adresse.Bundesland
		tabvalues[5] = p.(dat.PersonNat).Person.Adresse.Land
		log.Debug(">> Tabvalues:", tabvalues)

		err := sql.DoInsertRequest(tabvalues, tabcolumns)
		if err != nil {
			return err
		}
	}
	log.Debug(fmt.Sprintf("adressen: inserted %-02d records", len(co.Items)))

	// Speicherung Personen
	var sqlP SqlRequest
	//sqlP.SetRequest(dbs.DBT, "kunden", "INSERT", "")
	log.Debug(">> Create Kunden-Table Array")
	sqlP.SetRequest(DBM.Current, "Kunden", "INSERT", "")
	tabcolumns = []string{"ID", "SOCIALID", "VORNAME", "NAME"}
	for _, p := range co.Items {
		Entity := p.(dat.PersonNat).Person.Entity
		Person := p.(dat.PersonNat).Person
		Number := p.(dat.PersonNat).Number
		tabvalues = make([]string, len(tabcolumns))
		tabvalues[0] = Entity.Id.UId.String()
		tabvalues[1] = Number.RId
		tabvalues[2] = Person.Name.Firstname
		tabvalues[3] = Person.Name.Surname
		err := sqlP.DoInsertRequest(tabvalues, tabcolumns)
		if err != nil {
			return err
		}
	}
	log.Debug(fmt.Sprintf("kunden: inserted %-02d records ", len(co.Items)))
	return nil
}

func StorePersonsJur(co dat.CreatedObjects, dbname string) error {
	log.Info(">> Save Persons TO Database ", dbname)
	// Speicherung der Adressen
	dbm := &DBM
	err := dbm.SetCurrentDB(dbname)
	if err != nil {
		return err
	}
	var sql SqlRequest
	// Speicherung Adressen
	sql.SetRequest(DBM.Current, "Adressen", "INSERT", "")
	var tabcolumns = []string{"ID", "STRASSE", "PLZ", "ORT", "BUNDESLAND", "LAND"}
	var tabvalues []string
	log.Debug(">> Create Adressen-Table Array")

	for _, p := range co.Items {
		tabvalues = make([]string, len(tabcolumns))
		// tabvalues[0]=strconv.Itoa(p.Adresse.Id)
		tabvalues[0] = p.(dat.PersonJur).Person.Entity.Id.UId.String()
		log.Debug("AdresseID " + tabvalues[0])
		tabvalues[1] = p.(dat.PersonJur).Person.Adresse.Strasse
		tabvalues[2] = p.(dat.PersonJur).Person.Adresse.PLZ
		tabvalues[3] = p.(dat.PersonJur).Person.Adresse.Ort
		tabvalues[4] = p.(dat.PersonJur).Person.Adresse.Bundesland
		tabvalues[5] = p.(dat.PersonJur).Person.Adresse.Land
		log.Debug(">> Tabvalues:", tabvalues)

		err := sql.DoInsertRequest(tabvalues, tabcolumns)
		if err != nil {
			return err
		}
	}
	log.Debug(fmt.Sprintf("adressen: inserted %-02d records", len(co.Items)))

	// Speicherung Juristische Personen
	var sqlP SqlRequest
	log.Debug(">> Create Oraganization-Table Array")
	sqlP.SetRequest(DBM.Current, "Orgs", "INSERT", "")
	tabcolumns = []string{"ID", "VORNAME", "NAME", "REGID", "TYPE"}
	for _, p := range co.Items {
		Entity := p.(dat.PersonJur).Person.Entity
		Person := p.(dat.PersonJur).Person
		Number := p.(dat.PersonJur).Number
		Rechtsform := p.(dat.PersonJur).Rechtsform
		tabvalues = make([]string, len(tabcolumns))
		tabvalues[0] = Entity.Id.UId.String()
		tabvalues[1] = Person.Name.Firstname
		tabvalues[2] = Person.Name.Surname
		tabvalues[3] = Number.RId
		tabvalues[4] = Rechtsform

		err := sqlP.DoInsertRequest(tabvalues, tabcolumns)
		if err != nil {
			return err
		}
	}
	log.Debug(fmt.Sprintf("organisation: inserted %-02d records ", len(co.Items)))
	return nil
}

func StorePersonsRelations(co dat.CreatedObjects, dbname string) error {
	log.Info(">> Save Person-Relations TO Database ", dbname)
	// Speicherung der Adressen
	dbm := &DBM
	err := dbm.SetCurrentDB(dbname)
	if err != nil {
		return err
	}
	var sql SqlRequest
	// Speicherung Relations
	sql.SetRequest(DBM.Current, "RELATIONS", "INSERT", "")
	var tabcolumns = []string{"RSOURCE", "RTARGET", "TYPE"}
	var tabvalues []string
	log.Debug(">> Create Relations-Table Array")
	for _, p := range co.Items {
		tabvalues = make([]string, len(tabcolumns))
		// tabvalues[0]=strconv.Itoa(p.Adresse.Id)
		//tabvalues[0] = p.(Relation).Entity.Id.UId.String()
		//log.Debug("AdresseID " + tabvalues[0])
		tabvalues[0] = p.(dat.Relation).RSource.UId.String()
		tabvalues[1] = p.(dat.Relation).RTarget.UId.String()
		tabvalues[2] = p.(dat.Relation).Type
		log.Debug(">> Tabvalues:", tabvalues)

		err := sql.DoInsertRequest(tabvalues, tabcolumns)
		if err != nil {
			return err
		}
	}
	log.Debug(fmt.Sprintf("adressen: inserted %-02d records", len(co.Items)))

	return nil
}