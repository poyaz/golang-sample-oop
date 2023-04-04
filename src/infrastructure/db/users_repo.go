package db

import "golang-sample-oop/src/core/users"

var data = []users.Users{
	FakeUser1,
	FakeUser2,
	FakeUser3,
}

type UsersRepository struct {
}

func NewUsersRepository() *users.IRepository {
	usersRepository := &UsersRepository{}
	var out users.IRepository = usersRepository

	return &out
}

func (r *UsersRepository) GetAll() []users.Users {
	return data
}

func (r *UsersRepository) GetById(id uint64) users.Users {
	var find users.Users

	for _, user := range data {
		if user.Id == id {
			find = user
			break
		}
	}

	return find
}
