package apppki

import (
	// "database/sql"
	// "fmt"
    // "crypto/x509"
    // "bytes"
	// "crypto/x509/pkix"
	
	// _ "github.com/lib/pq"
)


type Certificate interface {
    getCertExpiration()                 // List all items of a selected categorie
    isCertValid() (bool)
}

func showExpiration(c Certificate) {
    c.getCertExpiration()
}

