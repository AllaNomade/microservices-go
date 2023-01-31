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

	err := SavePriceCar(20000.0, repository)
	assert.Error(t, err, "error in show value")

	err = SavePriceCar(55000.0, repository)
	assert.Error(t, err, "error in show value")

	repository.AssertExpectations(t)
}
