package services

import (
	"clean_gin_api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepoTestifyMock struct {
	mock.Mock
}

func (u UserRepoTestifyMock) CreateUser(user *models.User) (*models.User, error) {
	// user is args here ..
	args := u.Called()
	// get returns args at specific index
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := UserRepoTestifyMock{}

	user := models.User{Name: "dnes", Email: "test@gmail.com"}

	// setup expectations on mock repo
	mockRepo.On("CreateUser").Return(&user, nil)
	testService := NewUserService(mockRepo)
	result, err := testService.CreateUser(&user)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "dnes", result.Name)
	assert.Nil(t, err)
}
