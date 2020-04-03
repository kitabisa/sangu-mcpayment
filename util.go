package mcpayment

var (
	// PaymentName for prefix logging
	PaymentName = "[MC Payment]"
)

func isOK(httpStatus int) bool {
	return httpStatus > 200 && httpStatus < 299
}
