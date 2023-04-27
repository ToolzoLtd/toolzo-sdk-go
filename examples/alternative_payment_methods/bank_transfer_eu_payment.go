package alternative_payment_methods

type payloadBankTransferEuPayment struct {
	OrderId   string `json:"OrderId"`
	Amount    int    `json:"Amount"`
	Currency  string `json:"Currency"`
	ReturnUrl string `json:"ReturnUrl"`
}

func GetBankTransferEuPayment() (payloadBankTransferEuPayment, string) {
	const METHOD_TYPE = "BankTransferEu.Pay"
	payload := payloadBankTransferEuPayment{}

	payload.OrderId = "2700000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.ReturnUrl = "https://yandex.ru"

	return payload, METHOD_TYPE
}
