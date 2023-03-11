package controllers

import (
	"encoding/json"
	"exchange-api/models"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var client *http.Client

func GetExchangeRates(exchange models.Exchange) float64 {
	url := fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", exchange.From)

	var exchangeRatesResponse models.ExchangeRatesResponse

	GetJson(url, &exchangeRatesResponse)

	str := reflect.ValueOf(exchangeRatesResponse.Data.Rates).FieldByName(exchange.To).Interface()

	f, err := strconv.ParseFloat(str.(string), 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return 0
	} else {
		return exchange.Amount * f
	}
}

func GetJson(url string, target interface{}) error {
	client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var data = json.NewDecoder(resp.Body).Decode(target)

	return data
}

func GetExchange(c *gin.Context) {
	path := c.Request.RequestURI
	method := c.Request.Method

	var exchange models.Exchange
	
	user, _ := c.Get("user")
	userObj, _ := user.(models.User)

	if err := c.BindQuery(&exchange); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		LogRequest(userObj.ID , method, path, 400)
		return
	}

	var result = GetExchangeRates(exchange)

	if result == 0 {
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, gin.H{
		"from": exchange.From,
		"to": exchange.To,
		"amount": exchange.Amount,
		"result": result,
	})

	
	LogRequest(userObj.ID , method, path, 200)
}
