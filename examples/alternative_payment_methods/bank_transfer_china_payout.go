package alternative_payment_methods

type payloadBankTransferChinaPayout struct {
	OrderId          interface{} `json:"OrderId"`
	Amount           int         `json:"Amount"`
	Currency         string      `json:"Currency"`
	BankCode         string      `json:"BankCode"`
	RecipientAccount string      `json:"RecipientAccount"`
	ClientInfo       struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
	} `json:"ClientInfo"`
}

func GetBankTransferChinaPayout() (payloadBankTransferChinaPayout, string) {
	const METHOD_TYPE = "BankTransferChina.Payout"
	payload := payloadBankTransferChinaPayout{}

	payload.OrderId = "2600000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.BankCode = "mockBank"
	payload.RecipientAccount = "4564984561151"
	payload.ClientInfo.FirstName = "Ivan"
	payload.ClientInfo.LastName = "Ivanov"

	return payload, METHOD_TYPE
}
