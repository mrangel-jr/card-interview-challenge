package entities

import (
	"encoding/json"
	"fmt"
	"stone/cards/authorizer/internal/domain/errors"
	"time"
)

type Authorizer struct {
	CardNumber string
	Amount     float64
	Currency   string
	Merchant   string
	Timestamp  time.Time
}

func UnmarshalTransaction(data []byte) (Authorizer, error) {
	var shadow struct {
		CardNumber string  `json:"card_number"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
		Merchant   string  `json:"merchant"`
		Timestamp  string  `json:"timestamp"`
	}

	if err := json.Unmarshal(data, &shadow); err != nil {
		return Authorizer{}, fmt.Errorf("invalid JSON: %w", err)
	}

	parsedTime, err := time.Parse(time.RFC3339, shadow.Timestamp)
	if err != nil {
		return Authorizer{}, errors.ErrInvalidTimestamp
	}

	if parsedTime.After(time.Now()) {
		return Authorizer{}, errors.ErrInTheFuture
	}

	return Authorizer{
		CardNumber: shadow.CardNumber,
		Amount:     shadow.Amount,
		Currency:   shadow.Currency,
		Merchant:   shadow.Merchant,
		Timestamp:  parsedTime,
	}, nil
}
