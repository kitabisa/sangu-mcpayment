package mcpayment

import "errors"

var (
	// PaymentName for prefix logging
	PaymentName = "[MC Payment]"

	// ErrInvalidRequest error type for invalid request
	ErrInvalidRequest = errors.New("Invalid Request")

	// ErrUnauthorize error type for 401
	ErrUnauthorize = errors.New("Unauthorize")
)

func isOK(httpStatus int) bool {
	return httpStatus > 200 && httpStatus < 299
}
