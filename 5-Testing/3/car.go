package car

type Repository interface {
	ValueCar(value int) error
}

func SavePriceCar(value float64, repository Repository) error {
	amount := CarPrice(value)
	return repository.ValueCar(amount)
}

func CarPrice(value float64) int {
	if value <= 20000 {
		return 20000.0
	}
	if value == 55000 {
		return 55000.0
	}
	return 10000
}
