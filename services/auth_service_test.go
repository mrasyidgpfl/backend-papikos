package services_test

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/mocks"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthRepository_MatchingCredentialsReturnsUser(t *testing.T) {
	t.Run("Should return correct status code when signed in", func(t *testing.T) {
		// G
		mockRepository := new(mocks.AuthRepository)
		userService := services.NewAuthService(&services.AuthConfig{
			AuthRepository: mockRepository,
		})
		request := dto.SignInRequest{Email: "email1@mail.com", Password: "pass"}
		user := models.User{
			ID:       100001,
			Email:    "email10@mail.com",
			Password: "Pass",
			FullName: "Full Name 10",
			Address:  "Jakarta",
			CityId:   111111,
			Role:     "user",
		}

		// W
		mockRepository.On("MatchingCredential", "email1@mail.com", "pass").Return(&user, nil)
		tokenRes, _ := userService.SignIn(&request)

		// T
		// Token returned, therefore credentials matched.
		assert.NotEqual(t, nil, tokenRes.JWTToken.IDToken)
	})
}
