package repositories

import "github.com/takokun778/user-go-server/domains/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	Find(id string) (models.User, error)
	FindAll() ([]models.User, error)
	Update(models.User) (models.User, error)
	Delete(models.User) error
}
