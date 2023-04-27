package examples

type payloadPaymentLink struct {
	PaymentName string `json:"PaymentName"`
	OrderID     string `json:"Order ID"`
	Website     string `json:"Website"`
	Currency    string `json:"Currency"`
	Amount      int    `json:"Amount"`
	Description string `json:"Description"`
	Days        string `json:"Days"`
	Hours       string `json:"Hours"`
	Minutes     string `json:"Minutes"`
	Email       string `json:"Email"`
}

func GetPaymentLink() (payloadPaymentLink, string) {
	const METHOD_TYPE = "PaymentLink.Add"
	payload := payloadPaymentLink{}

	payload.PaymentName = "PaymentLinkName"
	payload.OrderID = "link-135"
	payload.Website = "616-test.com"
	payload.Currency = "RUB"
	payload.Amount = 150
	payload.Description = "test"
	//payload.Days = "00"
	payload.Hours = "2"
	//payload.Minutes = "00"
	payload.Email = "tratata@gmail.com"

	return payload, METHOD_TYPE
}
