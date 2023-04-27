package examples

type payloadPayout struct {
	OrderId  string `json:"OrderId"`
	Amount   int    `json:"Amount"`
	Currency string `json:"Currency"`
	CardInfo struct {
		CardNumber      string `json:"CardNumber"`
		CardHolder      string `json:"CardHolder"`
		ExpirationYear  string `json:"ExpirationYear"`
		ExpirationMonth string `json:"ExpirationMonth"`
	} `json:"CardInfo"`
}

func GetPayout() (payloadPayout, string) {
	const METHOD_TYPE = "Cards.Payout"
	payload := payloadPayout{}

	payload.OrderId = "2300000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.CardInfo.CardNumber = "4111111111111111"
	payload.CardInfo.CardHolder = "CARD HOLDER"
	payload.CardInfo.ExpirationYear = "2024"
	payload.CardInfo.ExpirationMonth = "12"

	return payload, METHOD_TYPE
}
