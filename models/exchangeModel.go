package models

type Exchange struct {
	From   string  `form:"from" binding:"required"`
	To     string  `form:"to" binding:"required"`
	Amount float64 `form:"amount" binding:"required"`
}

type ExchangeRatesResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Currency string `json:"currency"`
	Rates    Rates  `json:"rates"`
}

type Rates struct {
	USD string `json:"USD"`
	EUR string `json:"EUR"`
	BTC string `json:"BTC"`
	ETH string `json:"ETH"`
}
