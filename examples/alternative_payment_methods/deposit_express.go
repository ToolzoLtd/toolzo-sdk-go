package alternative_payment_methods

type payloadDepositExpress struct {
	OrderId    string `json:"OrderId"`
	Amount     int    `json:"Amount"`
	Currency   string `json:"Currency"`
	Bank       string `json:"Bank"`
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

func GetDepositExpress() (payloadDepositExpress, string) {
	const METHOD_TYPE = "DepositExpress.Pay"
	payload := payloadDepositExpress{}

	payload.OrderId = "2110000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.Bank = "mockBank"
	payload.ClientInfo.FirstName = "Ivan"
	payload.ClientInfo.LastName = "Ivanov"
	payload.ClientInfo.Birth = "1990-03-01T00:00:00Z"
	payload.ClientInfo.Email = "ivanov@gmail.com"
	payload.ClientInfo.TaxId = "50284414727"
	payload.ClientInfo.PhoneNumber = "75991435892"
	payload.ClientInfo.Zip = "38082365"

	return payload, METHOD_TYPE
}
