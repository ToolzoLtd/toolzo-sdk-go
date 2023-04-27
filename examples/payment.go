package examples

type payloadPayment struct {
	OrderId     string `json:"OrderId"`
	Amount      int    `json:"Amount"`
	ReturnUrl   string `json:"ReturnUrl"`
	Description string `json:"Description"`
	Currency    string `json:"Currency"`
	RebillFlag  bool   `json:"RebillFlag"`
	CardInfo    struct {
		CardNumber      string `json:"CardNumber"`
		CardHolder      string `json:"CardHolder"`
		ExpirationYear  string `json:"ExpirationYear"`
		ExpirationMonth string `json:"ExpirationMonth"`
		Cvv             string `json:"Cvv"`
	} `json:"CardInfo"`
}

func GetPayment() (payloadPayment, string) {
	const METHOD_TYPE = "Cards.Payment"
	payload := payloadPayment{}
	payload.OrderId = "2000005"
	payload.Amount = 150
	payload.ReturnUrl = "https://yandex.ru"
	payload.Description = "TEST"
	payload.Currency = "USD"
	payload.RebillFlag = true
	payload.CardInfo.CardNumber = "4111111111111111"
	payload.CardInfo.CardHolder = "CARD HOLDER"
	payload.CardInfo.ExpirationYear = "2024"
	payload.CardInfo.ExpirationMonth = "12"
	payload.CardInfo.Cvv = "736"
	return payload, METHOD_TYPE
}
