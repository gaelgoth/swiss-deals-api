package controllers

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

type GalaxusProduct struct {
	Store          string `json:"store"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	RemainingStock int    `json:"remaining_stock"`
	Stock          int    `json:"stock"`
	OfferPrice     int    `json:"offer_price"`
	RegularPrice   int    `json:"regular_price"`
	ImageUrl       string `json:"image_url"`
}

// Galaxus godoc
// @Summary     Get current deal from Digitec
// @Description receive the promotion of the day
// @Tags        Galaxus 🐢
// @Produce     json
// @Success     200 {object} controllers.GalaxusProduct
// @Router      /deals/digitec [get]
func GetDigitecOffer(c *fiber.Ctx) error {
	log.Println("GetDigitecOffer")
	var offerCollector []string = GetCollectorData([]string{"digitec.ch", "www.digitec.ch", "https://www.digitec.ch/en/daily-deal"})
	url := "https://www.digitec.ch"
	for k, v := range offerCollector {
		fmt.Println(k, v)
	}

	if len(offerCollector) == 0 {
		return errors.New("the deal cannot be loaded")
	}

	stockCollector := offerCollector[0]
	priceCollector := offerCollector[8]
	title := offerCollector[3]
	// The image is the last element collected, but depending on the structure of the page the index may change
	imageUrl := url + offerCollector[len(offerCollector)-1]

	price := getDealPrice(priceCollector)
	stock := getDealStock(stockCollector)

	product := &GalaxusProduct{
		Store:          "Digitec",
		Title:          title,
		Url:            url + "/liveshopping/",
		RemainingStock: stock["Remaining"],
		Stock:          stock["Available"],
		OfferPrice:     price["offerPrice"],
		RegularPrice:   price["RegularPrice"],
		ImageUrl:       imageUrl,
	}
	log.Println("Product : ", product)
	return c.Status(fiber.StatusOK).JSON(product)
}

// Galaxus godoc
// @Summary     Get current deal from Galaxus
// @Description receive the promotion of the day
// @Tags        Galaxus 🐢
// @Produce     json
// @Success     200 {object} controllers.GalaxusProduct
// @Router      /deals/galaxus [get]
func GetGalaxusOffer(c *fiber.Ctx) error {
	log.Println("GetGalaxusOffer")
	var offerCollector []string = GetCollectorData([]string{"galaxus.ch", "www.galaxus.ch", "https://www.galaxus.ch/fr/liveshopping/81"})
	url := "https://www.galaxus.ch"

	if len(offerCollector) == 0 {
		return errors.New("the deal cannot be loaded")
	}

	stockCollector := offerCollector[0]
	priceCollector := offerCollector[8]
	title := offerCollector[3]
	// The image is the last element collected, but depending on the structure of the page the index may change
	imageUrl := url + offerCollector[len(offerCollector)-1]
	price := getDealPrice(priceCollector)
	stock := getDealStock(stockCollector)

	product := &GalaxusProduct{
		Store:          "Galaxus",
		Title:          title,
		Url:            url + "/liveshopping/",
		RemainingStock: stock["Remaining"],
		Stock:          stock["Available"],
		OfferPrice:     price["offerPrice"],
		RegularPrice:   price["RegularPrice"],
		ImageUrl:       imageUrl,
	}
	log.Println("Product : ", product)

	return c.Status(fiber.StatusOK).JSON(product)
}

func GetCollectorData(domains []string) []string {
	collector := colly.NewCollector(
		colly.AllowedDomains(domains...),
	)

	data := make([]string, 0)

	collector.OnHTML(".sc-v6swez-17 div", func(element *colly.HTMLElement) {
		result := element.Text
		if len(result) != 0 {
			data = append(data, result)
		}
	})

	collector.OnHTML("article.sc-1k4pd2t-0", func(e *colly.HTMLElement) {
		image_element := e.DOM.Find("img").Eq(0)
		image, _ := image_element.Attr("src")
		data = append(data, image)
	})

	collector.Visit(domains[2])
	collector.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	for index, element := range data {
		fmt.Println(index, "=>", element)
	}

	return data
}

func getDealStock(stock string) map[string]int {
	stockDeal := make(map[string]int)
	re := regexp.MustCompile("[0-9]+")
	availableProductStock := re.FindAllString(stock, -1)

	if len(availableProductStock) == 1 { // Offer has ended
		availableStock, err := strconv.Atoi(availableProductStock[0])
		stockDeal["Remaining"] = 0
		stockDeal["Available"] = availableStock
		if err != nil {
			log.Println("Could not get availableStock when offer has ended")
		}
		return stockDeal
	}

	availableStock, err := strconv.Atoi(availableProductStock[1])
	remainingStock, err := strconv.Atoi(availableProductStock[0])
	if err != nil {
		log.Println("Could not get element remainingStock/availableStock")
	}
	stockDeal["Remaining"] = remainingStock
	stockDeal["Available"] = availableStock

	return stockDeal
}

func getDealPrice(priceCollector string) map[string]int {
	priceDeal := make(map[string]int)
	log.Println("price" + priceCollector)

	re := regexp.MustCompile("[0-9]+")
	price := re.FindAllString(priceCollector, -1)

	if len(price) <= 1 {
		log.Println("The price without discount is not indicated")
		offerPrice, err := strconv.Atoi(price[0])
		regularPrice := offerPrice
		if err != nil {
			log.Println("Could not get regularPrice/offerPrice when offer has ended")
		}
		priceDeal["offerPrice"] = offerPrice
		priceDeal["RegularPrice"] = regularPrice

	} else {
		offerPrice, err := strconv.Atoi(price[0])
		regularPrice, err := strconv.Atoi(price[1])
		if err != nil {
			log.Println("Could not get regularPrice/offerPrice when offer has ended")
		}
		priceDeal["offerPrice"] = offerPrice
		priceDeal["RegularPrice"] = regularPrice
	}
	return priceDeal
}
