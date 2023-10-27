package injection

import (
	"github.com/arifpermadi999/core-echo-golang/handlers"
	"github.com/arifpermadi999/core-echo-golang/repositories"
	"github.com/arifpermadi999/core-echo-golang/usecases"
	"gorm.io/gorm"
)

func InitAuthApi(db *gorm.DB) handlers.AuthHandler {
	repository := repositories.ProviderUserRepository(db)
	usecase := usecases.ProviderAuthUsecase(&repository)
	handler := handlers.ProviderAuthHandler(&usecase)
	return handler
}

func InitCategoryApi(db *gorm.DB) handlers.CategoryHandler {
	repository := repositories.ProviderCategoryRepository(db)
	usecase := usecases.ProviderCategoryUsecase(&repository)
	handler := handlers.ProviderCategoryHandler(&usecase)
	return handler
}
