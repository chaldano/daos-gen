package appconfig


import (
	"fmt"
	// "database/sql"
	// "io"
	// "net/http"
	// "regexp"
	//"encoding/json"
)
import errmod "trustkbb.de/daosgenerate/apperror"
import log "github.com/sirupsen/logrus"

// Globale Config-Variable
var Dbconf DBConfigs

func InitConfig(Path string) DBConfigs {
	// var LoggerConfig *log.Logger
	log.Info("Setup Config! (Init ConfigMethods)")
	var conf DBConfigs
	res := new(Ressources)
	conf.Path = Path
	conf.Urls = res 
	return conf 
}


func (conf DBConfigs) LoadConfig() (error) {
	// fmt.Println("ConfigPath",conf.Path)
	err:=LoadJSON(conf.Path, conf.Urls)
	if err != nil{
		return err
	}
	return nil
}

func (conf DBConfigs) CheckConfig(Cat string, Index string) (error) {
	for _, s := range conf.Urls.Configs {
		if (s.Cat == Cat && s.Name == Index) {
			return nil
		}
	}
	reason:= fmt.Sprintf("Category: <%s> or Index: <%s>  are wrong\n",Cat, Index)
	err := errmod.OopsError(reason)	
	return err
	
}

func (conf DBConfigs) ShowConfig(){
	for _, s := range conf.Urls.Configs {
		fmt.Printf("Category:%-10s\nName: %-10s\nIndex: %-5s\n\n",s.Cat,s.Url,s.Name)
	}	
}

