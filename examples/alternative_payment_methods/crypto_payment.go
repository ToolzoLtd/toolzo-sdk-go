package alternative_payment_methods

type payloadCryptoPayment struct {
	OrderId     string `json:"OrderId"`
	Amount      int    `json:"Amount"`
	Currency    string `json:"Currency"`
	ReturnUrl   string `json:"ReturnUrl"`
	Description string `json:"Description"`
	ClientInfo  struct {
		Email string `json:"Email"`
	} `json:"ClientInfo"`
}

func GetCryptoPayment() (payloadCryptoPayment, string) {
	const METHOD_TYPE = "Crypto.Pay"
	payload := payloadCryptoPayment{}

	payload.OrderId = "2800000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.ReturnUrl = "https://yandex.ru"
	payload.Description = "test"
	payload.ClientInfo.Email = "ivanov@gmail.com"

	return payload, METHOD_TYPE
}
