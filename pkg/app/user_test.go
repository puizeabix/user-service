package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/puizeabix/user-service/mocks"
	"github.com/puizeabix/user-service/pkg/domain"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockUserDB(mockCtrl)
	testUser := &domain.User{
		Username:  "unit",
		FirstName: "Unit",
		LastName:  "Test",
		Email:     "unit@test.com",
	}
	id := uuid.New()

	mockDB.EXPECT().Save(testUser).Return(id, nil).Times(1)

	svc := NewUserService(mockDB)
	ret, err := svc.Add(testUser)

	assert.Nil(t, err)
	assert.Equal(t, id, ret)
}

func TestGetUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockUserDB(mockCtrl)
	id := uuid.New()
	mockUser := &domain.User{
		ID:        id,
		FirstName: "Unit",
		LastName:  "Test",
		Email:     "unit@test.com",
		Username:  "unit",
	}
	mockDB.EXPECT().FindById(id).Return(mockUser, nil).Times(1)

	svc := NewUserService(mockDB)

	ret, err := svc.Get(id)

	assert.Nil(t, err)
	assert.Equal(t, mockUser.FirstName, ret.FirstName)
	assert.Equal(t, mockUser.LastName, ret.LastName)
	assert.Equal(t, mockUser.Username, ret.Username)
	assert.Equal(t, mockUser.Email, ret.Email)
}

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockUserDB(mockCtrl)
	id := uuid.New()

	mockDB.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	svc := NewUserService(mockDB)
	err := svc.Delete(id)
	assert.Nil(t, err)
}
