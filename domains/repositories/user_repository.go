package repositories

import "github.com/takokun778/user-go-server/domains/models"

type UserRepository interface {
	Save(user *models.User) (*models.User, *models.BaseError)
	Find(id string) (*models.User, *models.BaseError)
	FindAll() ([]*models.User, *models.BaseError)
	Update(user *models.User) (*models.User, *models.BaseError)
	Delete(id string) *models.BaseError
}
