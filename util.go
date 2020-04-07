package mcpayment

import "errors"

var (
	// PaymentName for prefix logging
	PaymentName = "[MC Payment]"

	// ErrInvalidRequest error type for invalid request
	ErrInvalidRequest = errors.New("Invalid Request")
)

func isOK(httpStatus int) bool {
	return httpStatus > 200 && httpStatus < 299
}
