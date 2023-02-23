package user

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return UserService{UserRepository: repository}
}

func (u *UserService) Save(user User) (User, error) {
	return u.UserRepository.Save(user)
}

func (u *UserService) FindAll() []User {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindById(id uint) (User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserService) Delete(user User) error {
	return u.UserRepository.Delete(user)
}
