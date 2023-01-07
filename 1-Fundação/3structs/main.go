package main

type user struct {
	ID int
}

// C *user aponta diretamente para a Struct user
func (C *user) main() {
	C.ID = 1
	println(C.ID)
}
