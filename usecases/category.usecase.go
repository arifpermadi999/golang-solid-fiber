package usecases

import (
	"github.com/arifpermadi999/core-echo-golang/models"
	"github.com/arifpermadi999/core-echo-golang/repositories"
)

type CategoryUsecase struct {
	CategoryRepository *repositories.CategoryRepository
}
type CategoryUsecaseContract interface {
	GetAllDataCategory() []models.Category
	ViewDataCategory(data *models.Category) (*models.Category, error)
	EditDataCategory(data *models.Category) (*models.Category, error)
	SaveCategory(data *models.Category) (*models.Category, error)
	UpdateCategory(data *models.Category) (*models.Category, error)
	DeleteCategory(data *models.Category) error
}

func ProviderCategoryUsecase(r *repositories.CategoryRepository) CategoryUsecase {
	return CategoryUsecase{CategoryRepository: r}
}

func (u CategoryUsecase) GetAllDataCategory() []models.Category {
	return u.CategoryRepository.FindAllCategory()
}
func (u CategoryUsecase) ViewDataCategory(data *models.Category) (*models.Category, error) {
	data, error := u.CategoryRepository.FindCategory(data)
	if error != nil {
		return data, error
	}
	return data, nil
}

func (u CategoryUsecase) SaveCategory(data *models.Category) (*models.Category, error) {

	data, error := u.CategoryRepository.SaveCategory(data)
	if error != nil {
		return data, error
	}
	return data, nil
}

func (u CategoryUsecase) UpdateCategory(data *models.Category) (*models.Category, error) {
	data, error := u.CategoryRepository.UpdateCategory(data)
	if error != nil {
		return data, error
	}
	return data, nil
}

func (u CategoryUsecase) DeleteCategory(data *models.Category) error {
	error := u.CategoryRepository.DeleteCategory(data)
	if error != nil {
		return error
	}
	return nil
}
