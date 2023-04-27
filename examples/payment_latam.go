package examples

type payloadPaymentLatam struct {
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
	ClientInfo struct {
		PhoneNumber string `json:"PhoneNumber"`
		FirstName   string `json:"FirstName"`
		LastName    string `json:"LastName"`
		Email       string `json:"Email"`
		TaxId       string `json:"TaxId"`
		Zip         string `json:"Zip"`
	} `json:"ClientInfo"`
}

func GetPaymentLatam() (payloadPaymentLatam, string) {
	const METHOD_TYPE = "CardsLatam.Payment"
	payload := payloadPaymentLatam{}

	payload.OrderId = "2200000"
	payload.Amount = 150
	payload.ReturnUrl = "https://yandex.ru"
	payload.Description = "DEBIT"
	payload.Currency = "RUB"
	payload.RebillFlag = true
	payload.CardInfo.CardNumber = "4111111111111111"
	payload.CardInfo.CardHolder = "CARD HOLDER"
	payload.CardInfo.ExpirationYear = "2024"
	payload.CardInfo.ExpirationMonth = "12"
	payload.CardInfo.Cvv = "736"
	payload.ClientInfo.PhoneNumber = "+75991435892"
	payload.ClientInfo.FirstName = "John"
	payload.ClientInfo.LastName = "Doe"
	payload.ClientInfo.Email = "doejohn@gmail.com"
	payload.ClientInfo.TaxId = "50284414727"
	payload.ClientInfo.Zip = "38082365"

	return payload, METHOD_TYPE
}
