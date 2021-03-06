package models

// Scraper Data Model
type ProductDetails struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Reviews     int64   `json:"reviews"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

// Health Check API response Model
type Health struct {
	Message string `json:"message"`
}

//scraper request Model
type ScraperRequest struct {
	Url string `json:"url"`
}
