package mcpayment

var (
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeAuth          = "AUTHENTICATION_ERROR"
	ErrCodeMissingData   = "MISSING_DATA"
	ErrCodeInvalidData   = "INVALID_DATA"
	ErrCodeSessExp       = "SESSION_EXPIRED"
	ErrCode500           = "INTERNAL_SERVER_ERROR"
	ErrCodeBank          = "BANK_ERROR"
	ErrCodeCard          = "CARD_ERROR"
	ErrCodePaymentSystem = "PAYMENT_SYSTEM_ERROR"
	ErrCodePayment       = "PAYMENT_ERROR"
)
