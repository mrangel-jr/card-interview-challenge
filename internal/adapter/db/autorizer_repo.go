package db

import (
	"stone/cards/authorizer/internal/domain/entities"

	"github.com/google/uuid"
)

type AuthorizerRepository struct {
	db map[string]entities.Authorizer
}

func (r *AuthorizerRepository) InsertAuthorizer(authorizer entities.Authorizer) (uuid.UUID, error) {
	// ... implements here
	uuid := uuid.New()
	r.db[uuid.String()] = authorizer
	return uuid, nil
}

func NewAuthorizerRepository() *AuthorizerRepository {
	db := make(map[string]entities.Authorizer)
	return &AuthorizerRepository{db: db}
}
