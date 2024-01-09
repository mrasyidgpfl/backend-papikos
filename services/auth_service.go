package services

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	SignIn(*dto.SignInRequest) (*dto.SignInResponse, error)
	SignUp(d *dto.SignUpRequest) (*dto.SignUpResponse, error)
	GetUserDetail(id uint) (*dto.UserDetailResponse, error)
	UpdateProfile(d *dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error)
	SignOut(token string) (*dto.SignOutResponse, error)
	CheckToken(token string) (bool, error)
	CleanBlacklists() (*dto.CleanBlacklistResponse, error)
}

type authService struct {
	authRepository repositories.AuthRepository
	appConfig      config.AppConfig
}

type AuthConfig struct {
	AuthRepository repositories.AuthRepository
	AppConfig      config.AppConfig
}

type jwtPayload struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	CityID   uint   `json:"city_id"`
	Role     string `json:"role"`
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	JWTUserPayload jwtPayload `json:"user"`
}

func NewAuthService(c *AuthConfig) AuthService {
	return &authService{
		authRepository: c.AuthRepository,
		appConfig:      c.AppConfig,
	}
}

func (a *authService) SignIn(req *dto.SignInRequest) (*dto.SignInResponse, error) {
	user, err := a.authRepository.MatchingCredential(req.Email, req.Password)
	if err != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
	}
	token, err := a.GenerateJWTToken(user)
	return new(dto.SignInResponse).CreateSignInResponse(200, "SIGNING IN SUCCESS", "User signed "+
		"in successfully", token), err
}

func (a *authService) GenerateJWTToken(user *models.User) (string, error) {
	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}
	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		JWTUserPayload: jwtPayload{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			Address:  user.Address,
			CityID:   user.CityId,
			Role:     user.Role,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (a *authService) SignUp(req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	userId, walletId, gamesChancesId, _ := a.GenerateUID()
	hashedPassword, _ := a.HashPassword(req.Password)
	user := models.User{
		ID:       uint(userId),
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
		Address:  req.Address,
		CityId:   req.CityID,
		Role:     "user",
	}
	wallet := models.Wallet{
		ID:      uint(walletId),
		UserId:  uint(userId),
		Balance: 0,
	}
	gamesChances := models.GamesChances{
		ID:     uint(gamesChancesId),
		UserId: uint(userId),
		Chance: 0,
	}
	insertedUser, rowsAffected, err := a.authRepository.SaveUser(&user)
	insertedWallet, _, err := a.authRepository.SaveWallet(&wallet)
	insertedGC, _, err := a.authRepository.SaveGamesChances(&gamesChances)
	if err == nil && rowsAffected == 0 {
		return new(dto.SignUpResponse), httperror.AppError{
			StatusCode: http.StatusConflict,
			Code:       "DUPLICATE EMAIL CONFLICT",
			Message:    "Email already registered, please try to sign up again with different email",
		}
	}
	return new(dto.SignUpResponse).CreateSignUpResponse(200, "USER CREATED SUCCESSFULLY", "User sign up and registered", insertedUser, insertedWallet, insertedGC), err
}

func (a *authService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *authService) GenerateUID() (int, int, int, error) {
	totalUser, _ := a.authRepository.CountTotalUser()
	totalUser += 1
	walletId := 900000 + totalUser
	userId := 100000 + totalUser
	gamesChancesId := 0 + totalUser
	return walletId, userId, gamesChancesId, nil
}

func (a *authService) GetUserDetail(userId uint) (*dto.UserDetailResponse, error) {
	user, wallet, gamesChances, err := a.authRepository.FindUserDetailsById(userId)
	if err != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusNotFound,
			Code:       "SOURCE NOT FOUND",
			Message:    "Source not found.",
		}
	}
	return new(dto.UserDetailResponse).CreateUserDetailRes(200, "USER DETAILS RETRIEVED", "UserRes details retrieved successfully", user, wallet, gamesChances), err
}

func (a *authService) UpdateProfile(request *dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error) {
	if len(request.Password) >= 1 {
		request.Password, _ = a.HashPassword(request.Password)
	}
	user, err := a.authRepository.UpdateProfile(request)
	if err != nil || user == nil {
		return nil, err
	}
	return new(dto.UpdateProfileResponse).CreateUpdateProfileRes(200, "USER PROFILE UPDATED", "User profile updated successfully", user), err
}

func (a *authService) SignOut(token string) (*dto.SignOutResponse, error) {
	_, err := a.authRepository.BlackListToken(token)
	return new(dto.SignOutResponse).CreateSignOutResponse(200, "SIGN OUT SUCCESSFUL", "User signed out successfully, token blacklisted"), err
}

func (a *authService) CheckToken(token string) (bool, error) {
	b, err := a.authRepository.CheckToken(token)
	return b, err
}

func (a *authService) CleanBlacklists() (*dto.CleanBlacklistResponse, error) {
	_, err := a.authRepository.CleanBlacklists()
	return new(dto.CleanBlacklistResponse).CreateCleanBlacklistsResponse(200, "BLACKLISTS CLEANUP SUCCESSFUL", "Blacklists table cleaned successfully"), err
}
