package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	helpers "github.com/rajamaniraja/webscraper_task/scraper_service/helpers"
	models "github.com/rajamaniraja/webscraper_task/scraper_service/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	res := models.Health{Message: "Sucess"}
	fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ScrapUrl(w http.ResponseWriter, r *http.Request) {
	var product models.ProductDetails
	var req models.ScraperRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
		w.WriteHeader(400)
		return
	}
	url := req.Url
	product, err = helpers.ScraperHelper(url)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
