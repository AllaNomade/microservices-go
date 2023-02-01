package car

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarValue(t *testing.T) {
	repository := &ValueRepositoryMock{}
	repository.On("ValueCar", 20000).Return(errors.New("error ins show value"))
	repository.On("ValueCar", 55000).Return(errors.New("error ins show value"))
	repository.On("ValueCar", 10000).Return(nil)

	err := SavePriceCar(20000.0, repository)
	assert.Error(t, err, "error in show value")

	err = SavePriceCar(55000.0, repository)
	assert.Error(t, err, "error in show value")

	err = SavePriceCar(10000, repository)
	assert.Nil(t, err)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "ValueCar", 3)
}

func TestPrice(t *testing.T) {
	value := 500
	expected := 10000

	result := CarPrice(float64(value))
	if result != expected {
		t.Error(expected, result)
	}
}
