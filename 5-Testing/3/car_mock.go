package car

import "github.com/stretchr/testify/mock"

type ValueRepositoryMock struct {
	mock.Mock
}

func (m *ValueRepositoryMock) ValueCar(value int) error {
	args := m.Called(value)
	return args.Error(0)
}
