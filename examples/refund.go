package examples

type payloadRefund struct {
	OrderId       string `json:"OrderId"`
	Amount        int    `json:"Amount"`
	ParentOrderId string `json:"ParentOrderId"`
	Currency      string `json:"Currency"`
	Description   string `json:"Description"`
}

func GetRefund() (payloadRefund, string) {
	const METHOD_TYPE = "Cards.Refund"
	payload := payloadRefund{}

	payload.OrderId = "10000018"
	payload.Amount = 150
	payload.ParentOrderId = "10000017"
	payload.Description = "TEST"
	payload.Currency = "RUB"

	return payload, METHOD_TYPE
}
