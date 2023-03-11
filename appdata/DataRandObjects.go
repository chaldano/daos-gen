package appdata

import (
	// "encoding/json"
	"fmt"
	"regexp"
	 "math/rand"
	 "time"
	//  "os/exec"
	// "bytes"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
	// "log"
)

// import net "trustkbb.de/daosgenerate/appnetwork"
import errmod "trustkbb.de/daosgenerate/apperror"

// import conf "trustkbb.de/daosgenerate/appconfig"

// Methods UUID
func LoadRObjects(cat string, anz int, index string) (RandomObjects, error) {
	var ro RandomObjects
	ro.Category = cat
	ro.Items = make([]interface{}, anz)
	var body []byte
	var err error

	//fmt.Printf("Cat:%s Anzahl:%d Index:%s\n",cat,anz,index)
	if index != "" {
		//fmt.Println("Durch...")
		// body, err = net.GetCategoryElements(ro.Category, index)
		if err != nil {
			return ro, err
		}
	}
	switch cat {
	case "UUID":
		for i := 0; i < anz; i++ {
			u := new(Uuid)
			u.UId = uuid.NewV4()
			// fmt.Println("UUID:",u.UId)
			ro.Items[i] = *u
		}
		return ro, nil
	
	case "CASuggest":
		content := fmt.Sprintf("%s", body)
		// Aufsplitten in Array
		re := regexp.MustCompile(`,`)
		names := re.Split(content, -1)
		// Filtern Namen
		r := regexp.MustCompile(`"(.*?)"`)
		for i, n := range names {
			namesneu := r.FindStringSubmatch(n)
			c := new(Name)
			c.Firstname = namesneu[1]
			c.Surname = "Certificate-Authority"
			ro.Items[i] = *c
		}
		return ro, nil

	// case "Zeichenketten":
	// 	content := fmt.Sprintf("%s", body)
	// 	r := regexp.MustCompile(`(?ms)\<pre.*?\>(.*?)\<\/pre\>`)
	// 	match := r.FindStringSubmatch(content)
	// 	if match != nil {
	// 		rs := regexp.MustCompile(`(?ms)^\w+?$`)
	// 		results := rs.FindAllString(match[1], -1)
	// 		for i, z := range results {
	// 			ro.Items[i] = z
	// 		}
	// 		return ro, nil
	// 	} else {
	// 		err := errmod.OopsError("No-String Match")
	// 		ro.Items[0] = "No-Match"
	// 		return ro, err
	// 	}
	// 	return ro, nil
	case "Zeichenketten":
		for i := 0; i < anz; i++ {
			z:= fake.DigitsN(10)
		 	ro.Items[i] = z
		}
		return ro, nil
	
	case "Adressen":
		
		for i := 0; i < anz; i++ {
			var a = new(Adresse)
			uuids, res := LoadRObjects("UUID", anz, "")
			if res != nil {
			  return ro, res
			}
			a.Id = uuids.Items[0].(Uuid)
			a.Etage = fake.DigitsN(1)
			a.Strasse = fake.Street()
			a.PLZ = fake.Zip()
			a.Ort = fake.City()
			a.Bundesland = "Bundesland"
			a.Land = fake.Country()	
			ro.Items[i] = *a
		}
		return ro, nil

	case "Zahlensequenz":
		content := fmt.Sprintf("%s", body)
		r := regexp.MustCompile(`(?ms)\d+?$`)
		match := r.FindAllString(content, -1)
		if match == nil {
			err := errmod.OopsError("No-Sequenze Match")
			return ro, err
		}
		for i, s := range match {
			ro.Items[i] = s
		}
		return ro, nil
	case "DigitsN":
		// Generate N Digits (beliebige Reihenfolge von Ziffern)
		Digits := fake.DigitsN(anz)
		// fmt.Println("DigitGeneratedX",Digits)
		for i := 0; i < anz; i++ {
			ro.Items[i] = string(Digits[i])
		}
		return ro, nil
	
	
	case "Namen":
		for i := 0; i < anz; i++ {
			var n = new(Name)
			n.Firstname = fake.MaleFirstName()
			n.Surname = fake.MaleLastName();
			ro.Items[i] = *n
		}
		return ro, nil

	case "RegisterNumbers":
		// Entfernen von " ... "
		re := regexp.MustCompile(`\"(.*?)\"`)
		var sn = new(RegisterNumber)
		sn.Country = "US"
		match := re.FindStringSubmatch(string(body))
		if match == nil {
			err := errmod.OopsError("No-RegisterNumber found")
			return ro, err
		}
		sn.RId = match[1]
		ro.Items[0] = *sn
		return ro, nil
	
	case "DomainNames":
		for i := 0; i < anz; i++ {
			dn := new(DomainName)
			dn.name = fake.DomainName()
			ro.Items[i] = *dn
			// dnzone := fake.DomainZone()
		}
		return ro, nil
	default:
		fmt.Println("Durch")
		err := errmod.OopsError("Default: Wrong DataCategory in LoadRObjects")
		return ro, err
	}
	return ro, err
}
func (ro RandomObjects) Listitems() {
	fmt.Println("\nReport Kategorie: ", ro.Category)
	for i, s := range ro.Items {
		switch ro.Category {
		case "UUID":
			fmt.Printf("%02d : %+v\n", i, s.(Uuid).UId)
		case "Zeichenketten":
			fmt.Printf("%02d : %+v\n", i, s)
		case "Adressen":
			fmt.Printf("%02d :\n", i)
			fmt.Printf("%-11s:%-5v\n", "Id", s.(Adresse).Id.UId)
			fmt.Printf("%-11s:%-5s\n", "Plz", s.(Adresse).PLZ)
			fmt.Printf("%-11s:%-10s\n", "Ort", s.(Adresse).Ort)
			fmt.Printf("%-11s:%-12s\n", "Strasse", s.(Adresse).Strasse)
			fmt.Printf("%-11s:%-12s\n", "Bundesland", s.(Adresse).Bundesland)
			fmt.Printf("%-11s:%-12s\n", "Land", s.(Adresse).Land)
		case "Zahlensequenz":
			fmt.Printf("%02d : %+v\n", i, s)
		case "Namen":
			fmt.Printf("%02d : %-8s %-8s\n", i, s.(Name).Firstname, s.(Name).Surname)
		case "CASuggest":
			fmt.Printf("%03d : %-8s %-10s\n", i, s.(Name).Firstname,s.(Name).Surname)
		case "RegisterNumbers":
			fmt.Printf("%02d : %-8s\n", i, s.(RegisterNumber).RId)

		case "DomainNames":
			// for i, s := range dns.Items {
			fmt.Printf("%02d : %+v\n", i, s.(DomainName).name)
			// }
		case "DigitsN":
			fmt.Printf("%02d : %s\n", i, s.(string))
			// }
		}
	}
}
func RandDigitFrom(anz int) (int,error) {
	var err error
	rand.Seed(time.Now().UnixNano())

	n:=rand.Intn(anz)
	if n < 0 {
		err = errmod.OopsError("Error in Random Digit")			
	}
	return n, err
}

func ReplaceSingleQuote(z string) string {
	re := regexp.MustCompile(`'`)
	// DoppleQuoting '' um SingleQuote im Value zu kennzeichnen
	r := re.ReplaceAllString(z, "''")
	return r
}

func Report(re ReportCategory) {
	re.Listitems()
}

