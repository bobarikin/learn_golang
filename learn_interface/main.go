package main

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

func (u *user) ChangeFIO(newFIO string) {
	u.FIO = newFIO
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func main() {

}