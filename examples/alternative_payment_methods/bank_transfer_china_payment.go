package alternative_payment_methods

type payloadBankTransferChinaPayment struct {
	OrderId    string `json:"OrderId"`
	Amount     int    `json:"Amount"`
	Currency   string `json:"Currency"`
	ReturnUrl  string `json:"ReturnUrl"`
	ClientInfo struct {
		FirstName   string `json:"FirstName"`
		LastName    string `json:"LastName"`
		Email       string `json:"Email"`
		PhoneNumber string `json:"PhoneNumber"`
		Zip         string `json:"Zip"`
		Country     string `json:"Country"`
		City        string `json:"City"`
		Address     string `json:"Address"`
	} `json:"ClientInfo"`
}

func GetBankTransferChinaPayment() (payloadBankTransferChinaPayment, string) {
	const METHOD_TYPE = "BankTransferChina.Pay"
	payload := payloadBankTransferChinaPayment{}

	payload.OrderId = "2500000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.ReturnUrl = "https://yandex.ru"
	payload.ClientInfo.FirstName = "Ivan"
	payload.ClientInfo.LastName = "Ivanov"
	payload.ClientInfo.Email = "ivanov@gmail.com"
	payload.ClientInfo.PhoneNumber = "75991435892"
	payload.ClientInfo.Zip = "123456"
	payload.ClientInfo.Country = "Russia"
	payload.ClientInfo.City = "Moscow"
	payload.ClientInfo.Address = "Mira street"

	return payload, METHOD_TYPE
}
