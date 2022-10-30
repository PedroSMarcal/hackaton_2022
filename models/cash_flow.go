package models

import (
	"time"
)

type CashFlow struct {
	Data    time.Time
	Account interface{}
	Amount  float64
	Entry   float64
	Out     float64
}
