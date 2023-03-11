
package apppki

import "bytes"
import "crypto/x509"
// import "crypto/x509/pkix"

// import "github.com/satori/go.uuid"
// func init() {
// 	fmt.Println("Database-Prgramms -This will get called on main initialization")
//   }
// import pki "trustkbb.de/tools/dbtools/apppki"

type CertBlock struct{
    PemCert *bytes.Buffer
	PemPrivKey *bytes.Buffer
	Cert *x509.Certificate	
}
