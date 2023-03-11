package appconfig

import (
	"fmt"
	"os"
	"encoding/json"
)
import errmod "trustkbb.de/daosgenerate/apperror"


// type ErrorReason struct {
// 	Path string
// }

// Serialisierung 

func LoadJSON(fileName string, key interface{}) (error) {
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Wrong file",fileName)
		reason:= fmt.Sprintf("Can´t open file\n",fileName)
		err := errmod.OopsError(reason)	
		return err
	}
	// errmod.CheckError("load jsonfile",err)
	decoder := json.NewDecoder(inFile)
	// fmt.Println("Decoder",decoder)
	err = decoder.Decode(key)
	errmod.CheckError("decoder",err)	
	inFile.Close()
	return nil
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	errmod.CheckError("SaveJson",err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	errmod.CheckError("Encode",err)
	outFile.Close()
}