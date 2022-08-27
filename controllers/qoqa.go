package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/gofiber/fiber/v2"
)

type QoqaProduct struct {
	Store          string `json:"store"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	Url            string `json:"url"`
	RemainingStock int64  `json:"remaining_stock"`
	Stock          int    `json:"stock"`
	OfferPrice     string `json:"offer_price"`
	RegularPrice   string `json:"regular_price"`
	ImageUrl       string `json:"image_url"`
}

// Qoqa godoc
// @Summary      Get current deal from QoQa
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qoqa [get]
func GetQoqaOffer(c *fiber.Ctx) error {
	log.Println("GetQoqaOffer")

	url_api := "https://api.qoqa.ch/v2/websites/wwwqoqach/offer_preview"

	return getCurrentDeal("Qoqa", url_api, c)
}

// Qoqa godoc
// @Summary      Get current deal from Qwine
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qwine [get]
func GetQwineOffer(c *fiber.Ctx) error {
	log.Println("GetQwineOffer")

	url_api := "https://api.qoqa.ch/v2/websites/qwineqoqach/offer_preview"

	return getCurrentDeal("Qwine", url_api, c)
}

// Qoqa godoc
// @Summary      Get current deal from Qbeer
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qbeer [get]
func GetBeerOffer(c *fiber.Ctx) error {
	log.Println("GetBeerOffer")

	url_api := "https://api.qoqa.ch/v2/websites/qbeerqoqach/offer_preview"

	return getCurrentDeal("Qbeer", url_api, c)
}

// Qoqa godoc
// @Summary      Get current deal from Qsport
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qsport [get]
func GetQsportOffer(c *fiber.Ctx) error {
	log.Println("GetQsportOffer")

	url_api := "https://api.qoqa.ch/v2/websites/qsportqoqach/offer_preview"

	return getCurrentDeal("Qsport", url_api, c)
}

// Qoqa godoc
// @Summary      Get current deal from Qooking
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qooking  [get]
func GetQookingOffer(c *fiber.Ctx) error {
	log.Println("GetQookingOffer")

	url_api := "https://api.qoqa.ch/v2/websites/qookingqoqach/offer_preview"

	return getCurrentDeal("Qooking", url_api, c)
}

// Qoqa godoc
// @Summary      Get current deal from Qlock
// @Description  receive the promotion of the day
// @Tags         Qoqa 🦦
// @Produce      json
// @Success      200  {object}  controllers.QoqaProduct
// @Router       /deals/qlock  [get]
func GetQlockOffer(c *fiber.Ctx) error {
	log.Println("GetQlockOffer")

	url_api := "https://api.qoqa.ch/v2/websites/qlockqoqach/offer_preview"

	return getCurrentDeal("Qlock", url_api, c)
}

func getCurrentDeal(store string, url_api string, c *fiber.Ctx) error {
	err, response := httpGetRequest(url_api)
	title := gjson.Get(response, "title").String()
	subtitle := gjson.Get(response, "subtitle").String()
	url := gjson.Get(response, "offer_url").String()
	remaining_stock := gjson.Get(response, "remaining_stock_percent").Int()
	price := gjson.Get(response, "offer_price_text").String()
	regular_price := gjson.Get(response, "regular_price_text").String()
	image_url := gjson.Get(response, "image_urls.standard.url").String()

	if err != nil {
		fmt.Println(err)
	}

	product := &QoqaProduct{
		Store:          store,
		Title:          title,
		Subtitle:       subtitle,
		Url:            url,
		RemainingStock: remaining_stock,
		Stock:          100,
		OfferPrice:     price,
		RegularPrice:   regular_price,
		ImageUrl:       image_url,
	}
	log.Println("Product : ", product)
	return c.Status(fiber.StatusOK).JSON(product)
}

func httpGetRequest(url_api string) (error, string) {
	resp, err := http.Get(url_api)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Can not get Qoqa API", err)
	}
	defer resp.Body.Close()
	response := string(body)
	return err, response
}
