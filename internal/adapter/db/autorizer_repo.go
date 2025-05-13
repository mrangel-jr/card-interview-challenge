package db

import (
	"stone/cards/authorizer/internal/domain/entities"

	"github.com/google/uuid"
)

type AuthorizerRepository struct {
	db map[string]entities.Authorizer
}

// Adicionar transação ao repositório
func (r *AuthorizerRepository) InsertAuthorizer(transaction entities.Authorizer) (uuid.UUID, error) {
	uuid := uuid.New()
	r.db[uuid.String()] = transaction
	return uuid, nil
}

func NewAuthorizerRepository() *AuthorizerRepository {
	db := make(map[string]entities.Authorizer, 0)
	return &AuthorizerRepository{db: db}
}
