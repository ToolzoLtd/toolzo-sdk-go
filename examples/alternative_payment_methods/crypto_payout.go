package alternative_payment_methods

type payloadCryptoPayout struct {
	OrderId  interface{} `json:"OrderId"`
	Amount   int         `json:"Amount"`
	Currency string      `json:"Currency"`
	Address  string      `json:"Address"`
}

func GetCryptoPayout() (payloadCryptoPayout, string) {
	const METHOD_TYPE = "Crypto.Payout"
	payload := payloadCryptoPayout{}

	payload.OrderId = "2900000"
	payload.Amount = 150
	payload.Currency = "RUB"
	payload.Address = "mqgsvC1CsPCW2NMaXAdFZFdqKhGH5kgAtC"

	return payload, METHOD_TYPE
}
