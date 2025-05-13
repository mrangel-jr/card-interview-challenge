package authorizer

import (
	"time"

	"stone/cards/authorizer/internal/adapter/ctrl/schema"
	"stone/cards/authorizer/internal/domain/entities"
	"stone/cards/authorizer/internal/domain/errors"

	"github.com/google/uuid"
)

const (
	riskAmountLimit      = 10_000.0
	riskAuthorizersLimit = 5
)

type AuthorizerRepository interface {
	InsertAuthorizer(authorizer entities.Authorizer) (uuid.UUID, error)
}

type RiskRepository interface {
	InsertRisk(risks entities.Risk) (uuid.UUID, error)
	GetCardTransactions(cardNumber string, after time.Time) []time.Time
	AddCardTransaction(cardNumber string, timestamp time.Time)
}

type AuthorizerUC struct {
	authRepo AuthorizerRepository
	riskRepo RiskRepository
}

// NewUCAuthorizer cria uma nova instância do use case de autorização
func NewAuthorizerUC(repo AuthorizerRepository, riskRepo RiskRepository) *AuthorizerUC {
	return &AuthorizerUC{
		authRepo: repo,
		riskRepo: riskRepo,
	}
}

// ValidateTransaction valida uma transação
func (uc *AuthorizerUC) ValidateTransaction(transaction entities.Authorizer) error {
	// Verifica se a data está no futuro
	if transaction.Timestamp.After(time.Now()) {
		return errors.ErrFutureTimestamp
	}

	// Verifica se os campos obrigatórios estão preenchidos
	if transaction.CardNumber == "" || transaction.Amount <= 0 || transaction.Currency == "" || transaction.Merchant == "" {
		return errors.ErrInvalidPayload
	}

	return nil
}

// CheckRisk verifica se uma transação é de risco
func (uc *AuthorizerUC) CheckRisk(transaction entities.Authorizer) (entities.RiskReason, bool) {
	// Verifica transações de valor alto
	if transaction.Amount > riskAmountLimit {
		return entities.RiskHighAmount, true
	}

	// Verifica se houve mais de 5 transações no último minuto
	oneMinuteAgo := time.Now().Add(-1 * time.Minute)
	transactions := uc.riskRepo.GetCardTransactions(transaction.CardNumber, oneMinuteAgo)

	// Registra a transação atual
	uc.riskRepo.AddCardTransaction(transaction.CardNumber, time.Now())

	// Verifica se há mais de 5 transações no último minuto
	if len(transactions) >= riskAuthorizersLimit {
		return entities.RiskNotStandard, true
	}

	return "", false
}

// ProcessTransaction processa uma transação
func (uc *AuthorizerUC) ProcessTransaction(transaction entities.Authorizer) schema.AuthorizerResponse {
	// Valida a transação
	if err := uc.ValidateTransaction(transaction); err != nil {
		return schema.AuthorizerResponse{
			Status: "rejected",
			Error:  err.Error(),
		}
	}

	// Verifica se a transação é de risco
	riskType, isRisk := uc.CheckRisk(transaction)

	// Se for de risco, registra o alerta
	if isRisk {
		alert := entities.Risk{
			CardNumber: transaction.CardNumber,
			Timestamp:  transaction.Timestamp,
			Reason:     riskType,
		}
		transactionID, err := uc.riskRepo.InsertRisk(alert)

		if err != nil {
			// Retorna o erro quando não consegue persistir no banco
			return schema.AuthorizerResponse{
				Status: "rejected",
				Error:  err.Error(),
			}
		}
		return schema.AuthorizerResponse{
			Status:      "approved_with_warning",
			AuthorizeID: transactionID.String(),
			Warning:     "transaction marked as suspicious: " + string(riskType),
		}
	}

	// Armazena a transação válida
	transactionID, err := uc.authRepo.InsertAuthorizer(transaction)
	if err != nil {
		// Retorna o erro quando não consegue persistir no banco
		return schema.AuthorizerResponse{
			Status: "rejected",
			Error:  err.Error(),
		}
	}
	// Retorna a resposta de sucesso
	return schema.AuthorizerResponse{
		Status:      "approved",
		AuthorizeID: transactionID.String(),
	}
}
