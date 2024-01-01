package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := customer.Validate()

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (customer *Customer) Validate() error {
	if customer.Name == "" {
		return errors.New("name should not be null")
	}

	if customer.Email == "" {
		return errors.New("email should not be null")
	}

	return nil
}

func (customer *Customer) Update(name string, email string) error {
	customer.Name = name
	customer.Email = email
	customer.UpdatedAt = time.Now()

	err := customer.Validate()

	if err != nil {
		return err
	}

	return nil
}
