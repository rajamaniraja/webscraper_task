package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/rajamaniraja/webscraper_task/scraper_service/models"
)

func ScraperHelper(url string) (models.ProductDetails, error) {
	res := models.ProductDetails{}
	c := colly.NewCollector()
	c.OnHTML("span[id=productTitle]", func(elem *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(elem.Text))
		res.Name = strings.TrimSpace(elem.Text)
	})
	c.OnHTML("span[id=acrCustomerReviewText]", func(elem *colly.HTMLElement) {
		formatedString := strings.Replace(strings.Split(elem.Text, " ")[0], ",", "", -1)
		ratings, _ := strconv.ParseInt(formatedString, 10, 32)
		res.Reviews = ratings
	})

	c.OnHTML(".twisterSwatchPrice", func(elem *colly.HTMLElement) {
		formatedString := strings.TrimSpace(strings.Replace(elem.Text, "$", "", -1))
		price, _ := strconv.ParseFloat(formatedString, 32)
		res.Price = price
	})

	c.OnHTML("div[id=productDescription]", func(elem *colly.HTMLElement) {
		description := strings.Replace(strings.TrimSpace(elem.Text), "\n", "", -1)
		res.Description = description
	})

	c.OnHTML("img[id=landingImage]", func(elem *colly.HTMLElement) {
		imageUrl := elem.Attr("data-a-dynamic-image")
		res.ImageUrl = imageUrl
	})

	c.Visit(url)
	return res, nil
}
