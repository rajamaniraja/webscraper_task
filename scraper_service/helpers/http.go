package helpers


import (
	"encoding/json"
	"net/http"
	"io/ioutil"
    "bytes"
	models "github.com/rajamaniraja/webscraper_task/scraper_service/models"
)

var serviceUrl string = "http://updater-service/api/v1/store"



func PostScrapData (product models.ProductDetails) (interface{}, error){
	requestBody, err := json.Marshal(product)
	if err != nil {
		return nil,err
	}

	response, err := http.Post(serviceUrl,"application/json",bytes.NewBuffer(requestBody))
	if err != nil {
		return nil,err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}

	return body, nil
}