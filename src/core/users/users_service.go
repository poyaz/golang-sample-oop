package users

type UsersService struct {
	usersRepository *IRepository
}

func NewUsersService(usersRepository *IRepository) *IService {
	userService := &UsersService{
		usersRepository: usersRepository,
	}
	var out IService = userService

	return &out
}

func (s *UsersService) GetAll() []Users {
	return (*s.usersRepository).GetAll()
}

func (s *UsersService) GetById(id uint64) Users {
	return (*s.usersRepository).GetById(id)
}
