package db

import (
	"stone/cards/authorizer/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type RiskRepository struct {
	alerts           map[string]entities.Risk
	cardTransactions map[string][]time.Time
}

// Adicionar risco ao repositório
func (r *RiskRepository) InsertRisk(risk entities.Risk) (uuid.UUID, error) {
	uuid := uuid.New()
	r.alerts[uuid.String()] = risk
	return uuid, nil
}

func NewRiskRepository() *RiskRepository {
	cardTransactions := make(map[string][]time.Time)
	alerts := make(map[string]entities.Risk, 0)
	return &RiskRepository{
		alerts:           alerts,
		cardTransactions: cardTransactions,
	}
}

func (r *RiskRepository) GetCardTransactions(cardNumber string, after time.Time) []time.Time {
	timestamps, exists := r.cardTransactions[cardNumber]
	if !exists {
		return []time.Time{}
	}

	// Filtra as transações após a data especificada
	recent := []time.Time{}
	for _, ts := range timestamps {
		if ts.After(after) {
			recent = append(recent, ts)
		}
	}
	return recent
}

// AddCardTransaction adiciona uma transação de cartão
func (r *RiskRepository) AddCardTransaction(cardNumber string, timestamp time.Time) {

	// Obtém as timestamps existentes
	timestamps, exists := r.cardTransactions[cardNumber]
	if !exists {
		timestamps = []time.Time{}
	}

	// Adiciona a nova timestamp
	timestamps = append(timestamps, timestamp)
	r.cardTransactions[cardNumber] = timestamps
}
