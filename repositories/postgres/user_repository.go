package postgres

import (
	"fmt"

	"github.com/luccasalves/apigo-sample/database"
	"github.com/luccasalves/apigo-sample/models"
)

type UserRepository struct {
}

func (r *UserRepository) Create(u models.User) (models.User, error) {

	db := database.GetDb()

	err := db.Create(&u).Error

	if err != nil {
		return models.User{}, err
	}

	return u, nil

}

func (s *UserRepository) Get(id uint) (models.User, error) {
	db := database.GetDb()
	var u models.User
	err := db.First(&u, id).Error

	if err != nil {
		return models.User{}, fmt.Errorf("cannot find user by id: %v", err)
	}

	return u, nil
}

func (s *UserRepository) GetAll() ([]models.User, error) {
	db := database.GetDb()
	var u []models.User
	err := db.Find(&u).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find users: %v", err)
	}

	return u, err
}
