package schema

import "time"

type AuthorizerResponse struct {
	Status      string `json:"status"`
	AuthorizeID string `json:"authorize_id,omitempty"`
	Error       string `json:"errors,omitempty"`
	Warning     string `json:"warning,omitempty"`
}

// Transaction representa os dados de uma transação de cartão
type Transaction struct {
	CardNumber string    `json:"card_number"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	Merchant   string    `json:"merchant"`
	Timestamp  time.Time `json:"timestamp"`
}
