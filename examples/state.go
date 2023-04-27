package examples

type payloadState struct {
	OrderId string `json:"OrderId"`
}

func GetState() (payloadState, string) {
	const METHOD_TYPE = "GetOrderState"
	payload := payloadState{}

	payload.OrderId = "2110000"

	return payload, METHOD_TYPE
}
