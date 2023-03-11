package appnetwork

import "os"
// import "encoding/json"
import "fmt"
// import data "trustkbb.de/daosgenerate/data"
import conf "trustkbb.de/daosgenerate/appconfig"
import errmod "trustkbb.de/daosgenerate/apperror"

import "errors"
import "net/http"
import "net/url"
import "strings"
import "io"
import "time"

	


func LoadJSONBody(Body string, key interface{}) {
	// inFile, err := os.Open(fileName)
	// checkError(err)
	// decoder := json.NewDecoder(inFile)
	// err = decoder.Decode(key)
	// checkError(err)
	// inFile.Close()
}

func GetCategoryElements(category string, index string) ([]byte, error) {
	// fmt.Println("Durch GetCat")
	
	urlres, _ := SearchUrl(category, index)
	body, err := LoadUrl(urlres)
	if err != nil {
		errmod.CheckError("GetCategoryElements",err)
		return nil,err
	}
	
	return body,nil
}
// Search Url in Config-File
func SearchUrl(cat string, index string) (string, error) {
	configData:=conf.Dbconf
	for i := 0; i < len(configData.Urls.Configs); i++ {
		if configData.Urls.Configs[i].Cat == cat && configData.Urls.Configs[i].Name == index {
			searchUrl := configData.Urls.Configs[i].Url
			return searchUrl, nil
		}
	}
	return "", errors.New("Url not found")
}
// Load Content from Url
func LoadUrl(urlres string) ([]byte, error) {
	// fmt.Println("Runnig url parsing")
	url, err := url.Parse(urlres)
	errmod.CheckError("LoadUrl",err)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	// client := &http.Client{
	// 	Timeout: time.Second * 5,
	// }
	// // url := urlreq
	request, err := http.NewRequest("GET", url.String(), nil)
	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	request.Header.Add("accept", "*/*")
	request.Header.Add("X-API-Key", "6362b8114dde49edb17a913320ce5169")
	//fmt.Println("Request: ", url.String())

	response, err := client.Do(request)
	errmod.CheckError("ClientRequest",err)
	defer response.Body.Close()

	if response.Status != "200 OK" {
		fmt.Println("Response-State",response.Status)
		os.Exit(2)
	}
	
	chSet := getCharset(response)
	// fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	body, _ := io.ReadAll(response.Body)
		return body ,err
	}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}