package example

type UserService struct {
	Repository *UserRepository
}

func (s *UserService) GetUsers() []string {
	return s.Repository.FindAll()
}
