package tax

import (
	"time"
)

func Calculate(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	if amount == 555 {
		return 55.0
	}
	return 5.0
}

func Calculate2(amount float64) float64 {
	time.Sleep(time.Millisecond * 700)
	if amount == 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
