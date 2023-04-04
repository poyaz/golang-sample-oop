package users

type IService interface {
	GetAll() []Users
	GetById(id uint64) Users
}

type IRepository interface {
	GetAll() []Users
	GetById(id uint64) Users
}
