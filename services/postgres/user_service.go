package postgres

import (
	"github.com/luccasalves/apigo-sample/interfaces"
	"github.com/luccasalves/apigo-sample/models"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func (s *UserService) CreateUser(data UserDTO) (models.User, error) {
	user := models.NewUser(data.Name, data.Age)

	return s.repository.Create(*user)
}

func (s *UserService) ShowUser(id uint) (models.User, error) {
	return s.repository.Get(id)
}

func (s *UserService) ShowAllUsers() ([]models.User, error) {
	return s.repository.GetAll()
}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}
