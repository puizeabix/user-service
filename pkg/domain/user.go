package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	Username  string    `json:"username" bson:"username"`
	FirstName string    `json:"firstname" bson:"firstname,omitempty"`
	LastName  string    `json:"lastname" bson:"lastname,omitempty"`
	Email     string    `json:"email" bson:"lastname"`
}

//go:generate mockgen -destination=../../mocks/mock_usersvc.go -package=mocks github.com/puizeabix/user-service/pkg/domain UserService

type UserService interface {
	Get(id uuid.UUID) (*User, error)
	List() ([]*User, error)
	Add(u *User) (uuid.UUID, error)
	Delete(id uuid.UUID) error
}

//go:generate mockgen -destination=../../mocks/mock_userdb.go -package=mocks github.com/puizeabix/user-service/pkg/domain domain.UserDB
type UserDB interface {
	FindById(uuid.UUID) (*User, error)
	FindAll() ([]*User, error)
	Save(u *User) (uuid.UUID, error)
	Delete(id uuid.UUID) error
}
