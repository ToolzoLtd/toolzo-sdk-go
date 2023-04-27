package alternative_payment_methods

type payloadSepa struct {
	OrderId    string `json:"OrderId"`
	Amount     int    `json:"Amount"`
	Currency   string `json:"Currency"`
	ClientInfo struct {
		FirstName   string `json:"FirstName"`
		LastName    string `json:"LastName"`
		PhoneNumber string `json:"PhoneNumber"`
		Email       string `json:"Email"`
		Country     string `json:"Country"`
		City        string `json:"City"`
		Zip         string `json:"Zip"`
		Address     string `json:"Address"`
		Birth       string `json:"Birth"`
	} `json:"ClientInfo"`
	Account struct {
		Iban    string `json:"Iban"`
		Country string `json:"Country"`
	} `json:"Account"`
}

func GetSepa() (payloadSepa, string) {
	const METHOD_TYPE = "Sepa.Payout"
	payload := payloadSepa{}

	payload.OrderId = "2140000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.ClientInfo.FirstName = "Ivan"
	payload.ClientInfo.LastName = "Ivanov"
	payload.ClientInfo.PhoneNumber = "75991435892"
	payload.ClientInfo.Email = "ivanov@gmail.com"
	payload.ClientInfo.Country = "Russia"
	payload.ClientInfo.City = "Moscow"
	payload.ClientInfo.Zip = "123456"
	payload.ClientInfo.Address = "Mira street"
	payload.ClientInfo.Birth = "1990-03-01T00:00:00Z"
	payload.Account.Iban = "DE89370400440532013000"
	payload.Account.Country = "AT"

	return payload, METHOD_TYPE
}
