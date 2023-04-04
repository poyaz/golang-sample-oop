package users

type DoubleAgeService struct {
	usersService *IService
}

func NewDoubleAgeService(usersService *IService) *IService {
	doubleAgeService := &DoubleAgeService{
		usersService: usersService,
	}
	var out IService = doubleAgeService

	return &out
}

func (s *DoubleAgeService) GetAll() []Users {
	users := (*s.usersService).GetAll()

	usersClone := make([]Users, len(users))
	copy(usersClone, users)

	for i := range users {
		usersClone[i].Age *= 2
	}

	return usersClone
}

func (s *DoubleAgeService) GetById(id uint64) Users {
	user := (*s.usersService).GetById(id)
	user.Age *= 2

	return user
}
