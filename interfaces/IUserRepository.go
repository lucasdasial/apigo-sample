package interfaces

import "github.com/luccasalves/apigo-sample/models"

type IUserRepository interface {
	Create(u models.User) (models.User, error)
	// Delete(id uint) (models.User, error)
	// Update(p models.User) error
	Get(id uint) (models.User, error)
	GetAll() ([]models.User, error)
}
