package examples

type payloadReversal struct {
	OrderId       string `json:"OrderId"`
	ParentOrderId string `json:"ParentOrderId"`
	Currency      string `json:"Currency"`
	Description   string `json:"Description"`
}

func GetReversal() (payloadReversal, string) {
	const METHOD_TYPE = "Cards.Reversal"
	payload := payloadReversal{}

	payload.OrderId = "2010005"
	payload.ParentOrderId = "2010004"
	payload.Description = "TEST"
	payload.Currency = "RUB"

	return payload, METHOD_TYPE
}
