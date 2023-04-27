package alternative_payment_methods

type payloadBoleto struct {
	OrderId    string `json:"OrderId"`
	Amount     int    `json:"Amount"`
	Currency   string `json:"Currency"`
	ClientInfo struct {
		FirstName   string      `json:"FirstName"`
		LastName    string      `json:"LastName"`
		Birth       interface{} `json:"Birth"`
		Email       string      `json:"Email"`
		TaxId       string      `json:"TaxId"`
		PhoneNumber string      `json:"PhoneNumber"`
		Zip         string      `json:"Zip"`
	} `json:"ClientInfo"`
}

func GetBoleto() (payloadBoleto, string) {
	const METHOD_TYPE = "Boleto.Pay"
	payload := payloadBoleto{}
	payload.OrderId = "2100000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.ClientInfo.FirstName = "Ivan"
	payload.ClientInfo.LastName = "Ivanov"
	payload.ClientInfo.Birth = "1990-03-01T00:00:00Z"
	payload.ClientInfo.Email = "ivanov@gmail.com"
	payload.ClientInfo.TaxId = "50284414727"
	payload.ClientInfo.PhoneNumber = "75991435892"
	payload.ClientInfo.Zip = "38082365"
	return payload, METHOD_TYPE
}
