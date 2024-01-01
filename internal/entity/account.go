package entity

import "time"

type Account struct {
	ID        string
	Customer  *Customer
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
