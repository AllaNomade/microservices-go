package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)

	tax, err = CalculateTax(20000)
	assert.Nil(t, err)
	assert.Equal(t, 20.0, tax)

	tax, err = CalculateTax(555)
	assert.Nil(t, err)
	assert.Equal(t, 55.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error save tax"))
	repository.On("SaveTax", 55.0).Return(nil)
	repository.On("SaveTax", 20.0).Return(nil)

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	err = CalculateTaxAndSave(20000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(555.0, repository)
	assert.Nil(t, err)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 4)
}
