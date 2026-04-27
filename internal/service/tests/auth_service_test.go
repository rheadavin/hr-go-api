package tests

import (
	"testing"

	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/models"
	"github.com/rheadavin/hr-go-api/internal/service"
	"github.com/rheadavin/hr-go-api/mocks"
	"github.com/rheadavin/hr-go-api/pkg/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthService_Login_Success(t *testing.T) {
	// 1. buat mock repository
	mockRepo := new(mocks.UserRepositoryInterface)

	// 2. setup expected behavior mock
	// ketika findByEmail dipanggil dengan email ini, return user ini
	hashedpw, _ := hash.HashPassword("admin123")
	mockUser := &models.User{
		Base: models.Base{
			ID: 1,
		},
		Email:    "jasondoe@yopmail.com",
		Password: hashedpw,
		Name:     "Jason Doe",
		Role:     "staff",
	}
	mockRepo.On("FindByEmail", "jasondoe@yopmail.com").Return(mockUser, nil)

	// 3. buat service dengan mock repo
	authService := service.NewAuthService(mockRepo)

	// 4. test
	req := dto.LoginRequest{
		Email:    "jasondoe@yopmail.com",
		Password: "admin123",
	}
	result, err := authService.Login(req)

	// 5. assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, "jasondoe@yopmail.com", result.User.Email)

	// pastikan mock dipanggil seperti yang diexpect
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	hashedPw, _ := hash.HashPassword("wrongpass")

	mockUser := &models.User{Email: "jasondoe@yopmail.com", Password: hashedPw}
	mockRepo.On("FindByEmail", "jasondoe@yopmail.com").Return(mockUser, nil)

	authService := service.NewAuthService(mockRepo)
	req := dto.LoginRequest{Email: "jasondoe@yopmail.com", Password: "wrongpassword"}
	result, err := authService.Login(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "email or password not valid", err.Error())
	mockRepo.AssertExpectations(t)
}
func TestAuthService_Login_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)

	mockRepo.On("FindByEmail", mock.AnythingOfType("string")).
		Return(nil, gorm.ErrRecordNotFound)

	authService := service.NewAuthService(mockRepo)

	req := dto.LoginRequest{Email: "unregistereduser@yopmail.com", Password: "any"}
	result, err := authService.Login(req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
