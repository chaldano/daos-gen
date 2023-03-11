package err

import (
	"fmt"
	"os"
	"errors"
	// "encoding/json"
)
// import  "trustkbb.de/tools/dbtools/config"
// import log "github.com/sirupsen/logrus"



// type ErrorReason struct {
// 	Path string
// }

func (e *ErrorReason) Error() string {
	return fmt.Sprintf("%v", e.Path)
}

func throwError(path string) error {
	return &ErrorReason{Path: path}
}

// Fehlerbehandlung
func CheckError(path string, err error) {
	if err != nil {
		// log.Error()
		fmt.Printf("Error: %s - Reason: %s\n",throwError(path), err)
		os.Exit(1)
	}
}
// func OopsError(path string, reason string) error {
func OopsError(reason string) error {
	// fmt.Printf("Error: %s - Reason: %s\n",throwError(path), reason)
	return errors.New(reason)
	// os.Exit(1)
}

