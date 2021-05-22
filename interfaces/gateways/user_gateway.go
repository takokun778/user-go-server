package gateways

import (
	"github.com/takokun778/user-go-server/domains/models"
	"github.com/takokun778/user-go-server/domains/repositories"
)

type UserGateway struct {
}

func NewUserGateway() repositories.UserRepository {
	return &UserGateway{}
}

func (gateway *UserGateway) Save(user *models.User) (*models.User, *models.BaseError) {
	return user, nil
}

func (gateway *UserGateway) Find(id string) (*models.User, *models.BaseError) {
	return nil, models.NewInternalServerError("Not implemented.").BaseError
}

func (gateway *UserGateway) FindAll() ([]*models.User, *models.BaseError) {
	return nil, models.NewInternalServerError("Not implemented.").BaseError
}

func (gateway *UserGateway) Update(user *models.User) (*models.User, *models.BaseError) {
	return nil, models.NewInternalServerError("Not implemented.").BaseError
}

func (gateway *UserGateway) Delete(id string) *models.BaseError {
	return models.NewInternalServerError("Not implemented.").BaseError
}
