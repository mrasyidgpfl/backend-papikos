package handlers_test

import (
	"encoding/json"
	"errors"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/mocks"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	t.Run("Should return the correct status code and body", func(t *testing.T) {
		// G
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		router := server.NewRouter(&config)
		request := dto.SignUpRequest{Email: "email10@mail.com", FullName: "Full Name 10", Password: "Pass"}
		response := dto.SignUpResponse{StatusCode: 200, Code: "SUCCESS", Message: "SUCCESS.",
			User: models.User{
				ID:       100001,
				Email:    "email10@mail.com",
				Password: "Pass",
				FullName: "Full Name 10",
				Address:  "Jakarta",
				CityId:   111111,
				Role:     "user",
			},
			Wallet: models.Wallet{
				ID:      900001,
				UserId:  100001,
				Balance: 0,
			},
			GamesChances: models.GamesChances{
				ID:      200001,
				UserId:  100001,
				Chance:  0,
				History: 0,
			},
		}

		// W
		mockService.On("SignUp", &request).Return(&response, nil)
		reqData, _ := json.Marshal(request)
		expectedBody, _ := json.Marshal(map[string]any{"data": response})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(string(reqData)))
		router.ServeHTTP(w, req)

		// T
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(expectedBody), w.Body.String())
	})
}

func TestUserHandler_SignUpDuplicateEmail(t *testing.T) {
	t.Run("Should return correct status code (error code) when response is nil", func(t *testing.T) {
		// G
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		router := server.NewRouter(&config)
		request := dto.SignUpRequest{Email: "email10@mail.com", FullName: "Full Name 10", Password: "Pass"}

		// W
		mockService.On("SignUp", &request).Return(nil, errors.New("duplicate email error"))
		reqData, _ := json.Marshal(request)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(string(reqData)))
		router.ServeHTTP(w, req)

		// T
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestUserHandler_SignInStatusOK(t *testing.T) {
	t.Run("Should return correct status code when signed in", func(t *testing.T) {
		// G
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		router := server.NewRouter(&config)
		request := dto.SignInRequest{Email: "email1@mail.com", Password: "pass"}
		response := dto.SignInResponse{
			StatusCode: 200,
			Code:       "SIGN IN SUCCESS",
			Message:    "UserRes signed in",
			JWTToken:   dto.IDToken{},
		}

		// W
		mockService.On("SignIn", &request).Return(&response, nil)
		reqJson, _ := json.Marshal(request)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/signin", strings.NewReader(string(reqJson)))
		router.ServeHTTP(w, req)

		// T
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUserHandler_SignInFailed(t *testing.T) {
	t.Run("Should return correct status code when signed in", func(t *testing.T) {
		// G
		mockService := new(mocks.AuthService)
		config := server.RouterConfig{AuthService: mockService}
		router := server.NewRouter(&config)
		request := dto.SignInRequest{Email: "email1@mail.com", Password: "wrongpassword"}

		// W
		mockService.On("SignIn", &request).Return(nil, errors.New("unauthorized error"))
		reqJson, _ := json.Marshal(request)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/signin", strings.NewReader(string(reqJson)))
		router.ServeHTTP(w, req)

		// T
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
