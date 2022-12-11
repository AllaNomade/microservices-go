package main

type user struct {
	ID int
}

func (C *user) main() {
	C.ID = 1
	println(C.ID)
}
