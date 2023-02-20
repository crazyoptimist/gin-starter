package user

type userDAO interface {
	Get(id uint) (*User, error)
}

type UserService struct {
	dao userDAO
}

func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

func (s *UserService) Get(id uint) (*User, error) {
	return s.dao.Get(id)
}
