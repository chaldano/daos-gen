package appdata

import (
	"fmt"
	"regexp"
	"testing"
	// "github.com/holmes89/hello-api/translation"
)

func TestEmails(t *testing.T) {

	// Arrange
	count := 10
	
	// Act
	emails, err := CreateNObjects("Email", count, nil)
	if err != nil {
		t.Errorf("Erro in Emails generation %s\n", err)
	}

	// Assert

	for _, d := range emails.Items {
		email := d.(Email).Adresse
		matched, _ := regexp.MatchString(`^[a-zA-Z]+(\.[a-zA-Z]+)?\@\w+\.\w{2,4}$`, email)
		if matched != true {
			t.Errorf("No match found in %s\n", email)
		}
	}
	Report(emails)
}
func TestEmailsExtern(t *testing.T) {
	// Namen werden als externe Daten übergeben
	// Arrange
	var names RandomObjects
	var emails CreatedObjects

	
	// Act

	count := 10
	n := fmt.Sprintf("%02d", count)
	index := "N" + n

	cat := "Namen"
	names, err := LoadRObjects(cat, count, index)
	if err != nil {
		t.Errorf("Erro in external Names generation %s\n", err)
	}
	Report(names)

	cat = "Email"
	emails, err = CreateNObjects(cat, count, names)
	if err != nil {
		t.Errorf("Erro in Emails generation %s\n", err)
	}

	// Assert

	for _, d := range emails.Items {
		email := d.(Email).Adresse
		matched, _ := regexp.MatchString(`^[a-zA-Z]+(\.[a-zA-Z]+)?\@\w+\.\w{2,4}$`, email)
		if matched != true {
			t.Errorf("No match found in %s\n", email)
		}
	}
	Report(emails)

}
