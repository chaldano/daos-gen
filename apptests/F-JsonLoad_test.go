package apptests

import "testing"
import "fmt"

import config "trustkbb.de/tools/dbtools/appconfig"

func TestLoadConfig(t *testing.T) {
	var dbconf config.DBConfigs
	var expected []string = []string{"Adressen", "Namen", "Zufallszahlen", "Farbcode"}
	var err error

	dbconf = config.InitConfig("./../appconfig/UrlConfig.json")
	// conf.Dbconf = conf.InitConfig("./../appconfig/UrlConfig.json")
	
	
	err = dbconf.LoadConfig()
		
	// _, err := fs.LoadJSON("./../con/dbconfig.json", &dbconf.Urls)
	if err != nil {
		t.Errorf("Error in loading config file %v", err)
	}
	// fmt.Println(dbconf.Urls.Configs)
	for i,_ :=range dbconf.Urls.Configs{
		fmt.Printf("%02d: %s\n",i,dbconf.Urls.Configs[i].Cat)
	}
	for index, value := range expected {
		want := value
		got := dbconf.Urls.Configs[index].Cat
		if got != want {
			t.Errorf("LoadJson() = %q, want %q", got, want)
		}
	}
}
