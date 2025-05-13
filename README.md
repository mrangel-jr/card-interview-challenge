# Solução - Fase 2

## Mapa do Projeto

```
card-interview-challenge/
├── .github/ # Arquivos de configuração GitHub (existente)
├── internal/ # Código interno da aplicação
│ ├── adapter/ # Adapters para interfaces externas
│ │ ├── ctrl/ # Controladores HTTP
│ │ │ ├── schema/ # Esquemas de request/response
│ │ │ │ └── authorizer.go # Esquema para transações
│ │ │ ├── authorizer_ctrl.go # Controlador de transações
│ │ │ └── authorizer_ctrl_test.go # Testes do controlador
│ │ └── db/ # Repositórios para armazenamento
│ │ ├── autorizer_repo.go # Repositório de transações
│ │ └── risk_repo.go # Repositório de transações de risco
│ └── domain/ # Núcleo do domínio
│ ├── authorizer/ # Casos de uso do autorizador
│ │ └── uc_authorizer.go # Implementação do caso de uso
│ ├── entities/ # Entidades do domínio
│ │ ├── authorizer.go # Entidade de transação e resposta
│ │ └── risk.go # Entidade de risco
│ └── errors/ # Erros de domínio
│ └── authorizer.go # Definição de erros
├── .gitignore # Arquivos ignorados pelo Git
├── main.go # Ponto de entrada
├── go.mod # Módulos Go
├── go.sum # Checksums
├── LICENSE # Licença
├── PROBLEM.md # Descrição do problema
└── README.md # Documentação atualizada
```

## Descrição da Solução

1. **Detecção de Transações Suspeitas**
   No use case, foi criado um método chamado _CheckRisk_ que checa os seguintes casos de fraude:

- Detecção de transações muito altas, com valor acima de **$10,000** deve ser marcada como suspeita: _high amount_ -
  Através de uma condicional, verifica-se se o valor da transação é maior que o **riskAmountLimit**. Em caso afirmativo, retorna como transação suspeita **entities.RiskHighAmount**.

- Compras fora do padrão: Se o mesmo número de cartão realizar mais de 5 transações em menos de 1 minuto, as transações adicionais devem ser marcadas como suspeitas. _not standard_ - Nesse caso foi criado um _map[string][]time.Time_., na qual a chave é o nº do cartão e o value é um array de timestamps. A ideia principal é adicionar o timestamp da transação ao slice do map cuja key é o nº do cartão. Então, para poder consultar as últims 5 transações no último minuto,
  o usa-se o método GetCardTransactions, passando a key (nº do cartão) e o timestamp atual menos 1 minuto. O método irá retornar um slice com todos os timestamps de transação daquele cartão acima do timestamp fornecido, e adiante, se houver 5 ou mais elementos nesse slice, retornará como transação suspeita **entities.RiskNotStandard**.

## Observações

- Alguns refactors foram realizados no controller a partir da etapa 1. Alguns métodos foram renomeados, assim como **Test4** do _TestAuthorizerCtrlEx2_ foi modificado para estar alinhado a solução projetada.
