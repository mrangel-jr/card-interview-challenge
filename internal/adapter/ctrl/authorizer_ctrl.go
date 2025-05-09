package ctrl

import (
	"encoding/json"
	"stone/cards/authorizer/internal/adapter/ctrl/schema"
	"stone/cards/authorizer/internal/domain/entities"
	"stone/cards/authorizer/internal/domain/errors"

	"github.com/google/uuid"
)

type AuthorizerUseCase interface {
	Authorize(authorizer entities.Authorizer) (uuid.UUID, error)
}

type AuthorizerCtrl struct {
	authorizerUC AuthorizerUseCase
}

func (a AuthorizerCtrl) Authorize(payload json.RawMessage) schema.AuthorizerResponse {
	customTransaction, err := entities.UnmarshalTransaction(payload)
	if err != nil {
		return schema.AuthorizerResponse{
			Status: "rejected",
			Error:  err.Error(),
		}
	}

	uuid, err := a.authorizerUC.Authorize(customTransaction)

	if err != nil {
		return schema.AuthorizerResponse{
			Status: "reject",
			Error:  errors.ErrInvalidPayload.Error(),
		}
	}

	return schema.AuthorizerResponse{
		Status:      "approved",
		AuthorizeID: uuid.String(),
	}
}

func NewAuthorizerCtrl(authorizerUC AuthorizerUseCase) AuthorizerCtrl {
	return AuthorizerCtrl{
		authorizerUC: authorizerUC,
	}
}
