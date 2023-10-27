package repositories

import (
	"github.com/arifpermadi999/core-echo-golang/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryContract interface {
	SaveOrUpdateUser(data *models.User) (*models.User, error)
	FindAllUser() []models.User
	FindUser(data *models.User) (*models.User, error)
	DeleteUser(data *models.User) error
}

func ProviderUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}
func (r UserRepository) FindUser(data *models.User) (*models.User, error) {
	err := r.DB.Where(data).First(&data)
	if err != nil {
		return data, err.Error
	}
	return data, nil
}

func (r UserRepository) SaveOrUpdateUser(data *models.User) (*models.User, error) {
	if err := r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r UserRepository) FindAllUser() []models.User {
	var datas []models.User
	r.DB.Find(&datas)
	return datas

}

func (r UserRepository) DeleteUser(data *models.User) error {
	if err := r.DB.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
