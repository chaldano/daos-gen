package apppki

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	// "crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	// "io/ioutil"
	"math/big"
	// "net"
	// "net/http"
	// // "net/http/httptest"
	// "strings"
	"time"
)
import log "github.com/sirupsen/logrus"
import errmod "trustkbb.de/daosgenerate/apperror"


var err error

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
	log.Info("Setup Config (Init Authority)")
}
// Erzeugung eines Certicate-Block
func CreateCACertificate(Name pkix.Name,Issuer pkix.Name) (*CertBlock,error) {

	// var cb CertBlock
	var cacert = new(x509.Certificate)	
	cb := new(CertBlock)
	
	// var Name pkix.Name
	cacert.SerialNumber = big.NewInt(2019)
	cacert.Subject = 			  Name
	cacert.Issuer = 			  Issuer

	cacert.NotBefore=             time.Now()
	cacert.NotAfter=              time.Now().AddDate(5, 0, 0)			// 5 Jahre
	cacert.IsCA=                  true
	cacert.ExtKeyUsage=           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	cacert.KeyUsage=              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
	cacert.BasicConstraintsValid= true

	
	log.Info("Create CA-Key")
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return cb,err
	}
	// Generate certificate in byte
	caBytes, err := x509.CreateCertificate(rand.Reader, cacert, cacert, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return cb,err
	}
	
	// Encoding in PEM 
    caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	cb.PemCert = caPEM
	
	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})
	log.Info("CAKeyPEM")
	cb.PemPrivKey = caPrivKeyPEM
	
	// fmt.Println("KeyPEM",caPrivKeyPEM)
	cb.Cert = cacert
	return cb,nil
}
func CreateUserCertificate(Name pkix.Name,Issuer pkix.Name, cacb *CertBlock) (*CertBlock,error) {

	// user CertBlock
	var usercert = new(x509.Certificate)	
	usercb := new(CertBlock)
	
	cacert := cacb.Cert
	
	caPrivKeyPEM := cacb.PemPrivKey 
	var caPrivKeyPEMData = caPrivKeyPEM.Bytes()
	block, rest := pem.Decode(caPrivKeyPEMData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing ca private key")
		err := errmod.OopsError("failed to decode PEM block containing ca private key")	
		return usercb, err 
	}	
	fmt.Println("Rest",rest)
	caPrivKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)	
	if err != nil {
		return usercb,err
	}
	// var Name pkix.Name
	fmt.Println("UserCert herstellen")
	
	usercert.SerialNumber = big.NewInt(2019)	
	usercert.Subject = 			  	Name
	usercert.Issuer = 			  	Issuer

	usercert.NotBefore=             time.Now()
	usercert.NotAfter=              time.Now().AddDate(2, 0, 0)			// 2 Jahre
	usercert.IsCA=                  false
	usercert.ExtKeyUsage=           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	// usercert.KeyUsage=              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
	usercert.KeyUsage=              x509.KeyUsageDigitalSignature
	usercert.BasicConstraintsValid= true

	


	log.Info("Create User-Key")
	userPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return usercb,err
	}
	// Generate certificate in byte (DER encoded)
	caBytes, err := x509.CreateCertificate(rand.Reader, usercert, cacert, &userPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return usercb,err
	}
	
	// Encoding in PEM 
    userPEM := new(bytes.Buffer)
	pem.Encode(userPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	usercb.PemCert = userPEM
	userPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(userPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(userPrivKey),
	})
	log.Info("UserKeyPEM")
	usercb.PemPrivKey = userPrivKeyPEM
	
	// fmt.Println("KeyPEM",userPrivKeyPEM)
	usercb.Cert = usercert
	return usercb,nil
}



