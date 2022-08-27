package app

import (
	"github.com/google/uuid"
	"github.com/puizeabix/user-service/pkg/domain"
)

// Implementation of domain.UserService
type userService struct {
	DB domain.UserDB
}

func NewUserService(db domain.UserDB) domain.UserService {
	return userService{
		DB: db,
	}
}

func (svc userService) Get(id uuid.UUID) (*domain.User, error) {
	return svc.DB.FindById(id)
}

func (svc userService) List() ([]*domain.User, error) {
	return svc.DB.FindAll()
}

func (svc userService) Add(u *domain.User) (uuid.UUID, error) {
	return svc.DB.Save(u)
}

func (svc userService) Delete(id uuid.UUID) error {
	return svc.DB.Delete(id)
}
