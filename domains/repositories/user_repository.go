package repositories

import "github.com/takokun778/user-go-server/domains/models"

type UserRepository interface {
	Save(user *models.User) (*models.User, models.IBaseError)
	Find(id string) (*models.User, models.IBaseError)
	FindAll() ([]*models.User, models.IBaseError)
	Update(user *models.User) (*models.User, models.IBaseError)
	Delete(id string) models.IBaseError
}
