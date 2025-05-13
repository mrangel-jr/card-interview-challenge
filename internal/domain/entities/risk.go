package entities

import (
	"time"
)

type Risk struct {
	CardNumber string
	Reason     RiskReason
	Timestamp  time.Time
}

// Tipo de Risco
type RiskReason = string

// Riscos comuns
const (
	RiskHighAmount  RiskReason = "high amount"
	RiskNotStandard RiskReason = "not standard"
)
