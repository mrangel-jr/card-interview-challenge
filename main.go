package main

import (
	"log"
	"net/http"
	"os"

	"stone/cards/authorizer/internal/adapter/ctrl"
	"stone/cards/authorizer/internal/adapter/db"
	"stone/cards/authorizer/internal/domain/authorizer"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Criar repositórios
	authorizerRepo := db.NewAuthorizerRepository()
	riskRepo := db.NewRiskRepository()

	// Inicializar o use case de autorização
	authorizerUseCase := authorizer.NewAuthorizerUC(authorizerRepo, riskRepo)

	// Inicializar o controlador
	authorizerController := ctrl.NewAuthorizerCtrl(authorizerUseCase)

	// Configurar o roteador Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rotas da API
	r.Post("/transactions", authorizerController.ProcessTransaction)

	// Iniciar o servidor
	port := "8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	log.Printf("Servidor iniciado na porta %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
