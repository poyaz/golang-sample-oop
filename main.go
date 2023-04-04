package main

import (
	"fmt"
	"golang-sample-oop/src/core/users"
	"golang-sample-oop/src/infrastructure/db"
	"strings"
)

func main() {
	// DI start

	// Repository start

	var usersRepository = db.NewUsersRepository()

	// Repository end

	// Service start

	var usersService = users.NewUsersService(usersRepository)
	var doubleAgeService = users.NewDoubleAgeService(usersService)

	// Service end

	// DI end

	a4 := (*usersService).GetAll()
	fmt.Println("UsersService -> GetAll()", a4)

	a2 := (*doubleAgeService).GetAll()
	fmt.Println("DoubleAgeService -> GetAll()", a2)

	a3 := (*usersService).GetAll()
	fmt.Println("UsersService -> GetAll()", a3)

	fmt.Println(strings.Repeat("#", 100))

	i1 := (*usersService).GetById(1)
	fmt.Println("UsersService -> GetById(1)", i1)

	i2 := (*doubleAgeService).GetById(1)
	fmt.Println("DoubleAgeService -> GetById(1)", i2)

	i3 := (*usersService).GetById(1)
	fmt.Println("UsersService -> GetById(1)", i3)
}
