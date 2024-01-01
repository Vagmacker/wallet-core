package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAValidParams_WhenCallsNewCustomer_ShouldReturnCustomer(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "john@doe.com"

	customer, err := NewCustomer(expectedName, expectedEmail)

	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.CreatedAt)
	assert.NotNil(t, customer.UpdatedAt)
	assert.Equal(t, expectedName, customer.Name)
	assert.Equal(t, expectedEmail, customer.Email)
}

func TestGivenAnInvalidName_WhenCallsNewCustomer_ShouldReturnError(t *testing.T) {
	expectedName := ""
	expectedEmail := "john@doe.com"

	customer, err := NewCustomer(expectedName, expectedEmail)

	expectedErrorMessage := "name should not be null"

	assert.NotNil(t, err)
	assert.Nil(t, customer)
	assert.Equal(t, err.Error(), expectedErrorMessage)
}

func TestGivenAnInvalidEmail_WhenCallsNewCustomer_ShouldReturnError(t *testing.T) {
	expectedEmail := ""
	expectedName := "John Doe"

	customer, err := NewCustomer(expectedName, expectedEmail)

	expectedErrorMessage := "email should not be null"

	assert.NotNil(t, err)
	assert.Nil(t, customer)
	assert.Equal(t, err.Error(), expectedErrorMessage)
}

func TestGivenAValidParams_WhenCallsUpdate_ShouldReturnCustomerUpdated(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "john@doe.com"

	customer, _ := NewCustomer("John Wick", expectedEmail)

	createdAt := customer.CreatedAt
	updatedAt := customer.UpdatedAt

	err := customer.Update(expectedName, expectedEmail)

	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, expectedName, customer.Name)
	assert.Equal(t, expectedEmail, customer.Email)
	assert.Equal(t, createdAt, customer.CreatedAt)
	assert.NotEqual(t, updatedAt, customer.UpdatedAt)
}
