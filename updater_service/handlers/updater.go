package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	helpers "github.com/rajamaniraja/webscraper_task/updater_service/helpers"
	models "github.com/rajamaniraja/webscraper_task/updater_service/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	res := models.Health{Message: "Sucess"}
	fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func StoreData(w http.ResponseWriter, r *http.Request) {
	var product models.ProductDetails
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
		w.WriteHeader(400)
		return
	}
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
		w.WriteHeader(400)
		return
	}

	database, err := helpers.GetDatabase()

	productColl := database.Collection("products")

	result, err := productColl.InsertOne(context.Background(), product)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	response := models.UpdateResponse{Message: "Sucess"}
	json.NewEncoder(w).Encode(response)
}
