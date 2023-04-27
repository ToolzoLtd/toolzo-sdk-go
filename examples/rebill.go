package examples

type payloadRebill struct {
	OrderId   string `json:"OrderId"`
	Amount    int    `json:"Amount"`
	Currency  string `json:"Currency"`
	BindingId int    `json:"BindingId"`
}

func GetRebill() (payloadRebill, string) {
	const METHOD_TYPE = "Cards.Rebill"
	payload := payloadRebill{}

	payload.OrderId = "10000069"
	payload.Amount = 150
	payload.Currency = "USD"
	payload.BindingId = 106837

	return payload, METHOD_TYPE
}
