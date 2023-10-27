package repositories

import (
	"github.com/arifpermadi999/core-echo-golang/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepositoryContract interface {
	SaveCategory(data *models.Category) (*models.Category, error)
	UpdateCategory(data *models.Category) (*models.Category, error)
	FindAllCategory() []models.Category
	FindCategory(data *models.Category) (*models.Category, error)
	DeleteCategory(data *models.Category) error
}

func ProviderCategoryRepository(DB *gorm.DB) CategoryRepository {
	return CategoryRepository{DB: DB}
}
func (r CategoryRepository) FindCategory(data *models.Category) (*models.Category, error) {
	err := r.DB.Where(data).First(&data)
	if err != nil {
		return data, err.Error
	}
	return data, nil
}

func (r CategoryRepository) SaveCategory(data *models.Category) (*models.Category, error) {
	if err := r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r CategoryRepository) UpdateCategory(data *models.Category) (*models.Category, error) {
	if err := r.DB.Updates(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r CategoryRepository) FindAllCategory() []models.Category {
	var datas []models.Category
	r.DB.Find(&datas)
	return datas

}

func (r CategoryRepository) DeleteCategory(data *models.Category) error {
	if err := r.DB.Delete(data).Error; err != nil {
		return err
	}
	return nil
}
