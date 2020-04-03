package mcpayment

func isOK(httpStatus int) bool {
	return httpStatus > 200 && httpStatus < 299
}
