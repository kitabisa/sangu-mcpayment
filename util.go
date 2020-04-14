package mcpayment

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

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

func validateSignatureKey(xSignKey, registerID, SignatureKey string) bool {
	dataToHash := []byte(fmt.Sprint(xSignKey, registerID))
	hashToValidate := sha256.Sum256(dataToHash)

	return fmt.Sprintf("%x", hashToValidate) == SignatureKey
}
