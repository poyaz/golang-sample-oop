package db

import "golang-sample-oop/src/core/users"

var FakeUser1 = users.Users{
	Id:     1,
	Name:   "Price",
	Age:    45,
	Gender: users.Male,
}

var FakeUser2 = users.Users{
	Id:     2,
	Name:   "Alex",
	Age:    30,
	Gender: users.Male,
}

var FakeUser3 = users.Users{
	Id:     3,
	Name:   "Farah",
	Age:    26,
	Gender: users.Female,
}
