package main

import (
	"ToolzoExample/examples"
	"fmt"
)

func main() {
	payload, method := examples.GetPayment()
	//payload, method := examples.GetPaymentLatam()
	//payload, method := examples.GetPaymentLink()
	//payload, method := examples.GetPayout()
	//payload, method := examples.GetRebill()
	//payload, method := examples.GetRefund()
	//payload, method := examples.GetReversal()

	//payload, method := examples.GetState()

	//payload, method := alternative_payment_methods.GetBoleto()
	//payload, method := alternative_payment_methods.GetBankTransferChinaPayment()
	//payload, method := alternative_payment_methods.GetBankTransferChinaPayout()
	//payload, method := alternative_payment_methods.GetBankTransferEuPayment()
	//payload, method := alternative_payment_methods.GetCryptoPayment()
	//payload, method := alternative_payment_methods.GetCryptoPayout()
	//payload, method := alternative_payment_methods.GetDepositExpress()
	//payload, method := alternative_payment_methods.GetLottery()
	//payload, method := alternative_payment_methods.GetPicPayment()
	//payload, method := alternative_payment_methods.GetPix()
	//payload, method := alternative_payment_methods.GetSepa()

	httpMessageFinal := MakeRequest(payload, method)
	result := SendRequest(httpMessageFinal)

	fmt.Printf("%+v\n", result)
}
