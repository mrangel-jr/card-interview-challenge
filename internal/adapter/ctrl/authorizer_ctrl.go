package ctrl

import (
	"encoding/json"
	"io"
	"net/http"
	"stone/cards/authorizer/internal/adapter/ctrl/schema"
	"stone/cards/authorizer/internal/domain/entities"
	"stone/cards/authorizer/internal/domain/errors"
)

type AuthorizerUseCase interface {
	ProcessTransaction(transaction entities.Authorizer) schema.AuthorizerResponse
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

	response := a.authorizerUC.ProcessTransaction(customTransaction)

	return response

}

func NewAuthorizerCtrl(authorizerUC AuthorizerUseCase) AuthorizerCtrl {
	return AuthorizerCtrl{
		authorizerUC: authorizerUC,
	}
}

func (ctrl *AuthorizerCtrl) ProcessTransaction(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		ctrl.sendErrorResponse(w, errors.ErrInvalidPayload)
		return
	}
	payload := json.RawMessage(bodyBytes)

	response := ctrl.Authorize(payload)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (ctrl *AuthorizerCtrl) sendErrorResponse(w http.ResponseWriter, err error) {
	response := schema.AuthorizerResponse{
		Status: "rejected",
		Error:  err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
