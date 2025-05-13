package errors

// AuthorizerError representa um erro de validação
type AuthorizerError struct {
	Message string
}

func (e AuthorizerError) Error() string {
	return e.Message
}

// Erros comuns
var (
	ErrInvalidPayload   = AuthorizerError{"invalid payload"}
	ErrInvalidTimestamp = AuthorizerError{"timestamp not valid"}
	ErrFutureTimestamp  = AuthorizerError{"timestamp on future"}
)
