package repositories

import (
	"errors"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/httperror"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type AuthRepository interface {
	MatchingCredential(email, password string) (*models.User, error)
	CountTotalUser() (int, error)
	SaveWallet(waller *models.Wallet) (*models.Wallet, int, error)
	SaveGamesChances(gamesChances *models.GamesChances) (*models.GamesChances, int, error)
	SaveUser(user *models.User) (*models.User, int, error)
	FindUserDetailsById(userId uint) (*models.User, *models.Wallet, *models.GamesChances, error)
	UpdateProfile(request *dto.UpdateProfileRequest) (*models.User, error)
	BlackListToken(token string) (bool, error)
	CheckToken(token string) (bool, error)
	CleanBlacklists() (bool, error)
}

type authRepository struct {
	db *gorm.DB
}

type AuthConfig struct {
	DB *gorm.DB
}

func NewAuthRepository(b *AuthConfig) AuthRepository {
	return &authRepository{db: b.DB}
}

func (a *authRepository) MatchingCredential(email string, password string) (*models.User, error) {
	var user *models.User
	err := a.db.Table("users").Where("email=?", email).First(&user).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, err
	}
	var userPassword string
	a.db.Table("users").Select("password").Where("email=?", email).Scan(&userPassword)
	checkPass, _ := CheckPasswordHash(password, userPassword)
	if !checkPass {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
	}
	return user, err
}

func CheckPasswordHash(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, nil
}

func (a *authRepository) CountTotalUser() (int, error) {
	var totalUser int64
	a.db.Table("users").Count(&totalUser)
	return int(totalUser), nil
}

func (a *authRepository) SaveWallet(wallet *models.Wallet) (*models.Wallet, int, error) {
	result := a.db.Table("wallets").Clauses(clause.OnConflict{DoNothing: true}).Create(wallet)
	return wallet, int(result.RowsAffected), result.Error
}

func (a *authRepository) SaveGamesChances(gamesChances *models.GamesChances) (*models.GamesChances, int, error) {
	result := a.db.Table("games_chances").Clauses(clause.OnConflict{DoNothing: true}).Create(gamesChances)
	return gamesChances, int(result.RowsAffected), result.Error
}

func (a *authRepository) SaveUser(user *models.User) (b *models.User, rowsEffected int, err error) {
	result := a.db.Table("users").Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	return user, int(result.RowsAffected), result.Error
}

func (a *authRepository) FindUserDetailsById(userId uint) (*models.User, *models.Wallet, *models.GamesChances, error) {
	var user *models.User
	var wallet *models.Wallet
	var gamesChances *models.GamesChances
	err := a.db.Table("users").Where("id=?", userId).First(&user).Error
	err = a.db.Table("wallets").Where("user_id=?", userId).First(&wallet).Error
	err = a.db.Table("games_chances").Where("user_id=?", userId).First(&gamesChances).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, nil, nil, err
	}
	return user, wallet, gamesChances, nil
}

func (a *authRepository) UpdateProfile(request *dto.UpdateProfileRequest) (*models.User, error) {
	if len(request.FullName) > 0 {
		a.db.Table("users").Where("id=?", request.ID).Update("full_name", request.FullName)
	}
	if len(request.Password) > 0 {
		a.db.Table("users").Where("id=?", request.ID).Update("password", request.Password)
	}
	if len(request.Address) > 0 {
		a.db.Table("users").Where("id=?", request.ID).Update("address", request.Address)
	}
	a.db.Table("users").Where("id=?", request.ID).Update("city_id", request.CityID)
	a.db.Table("users").Where("id=?", request.ID).Update("role", request.Role)
	var user *models.User
	result := a.db.Table("users").Where("id=?", request.ID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (a *authRepository) BlackListToken(token string) (bool, error) {
	var s int64
	a.db.Table("blacklists").Count(&s)
	blacklist := models.Blacklist{
		Model: gorm.Model{},
		ID:    uint(s + 1),
		Token: token,
	}
	res := a.db.Table("blacklists").Clauses(clause.OnConflict{DoNothing: true}).Create(&blacklist)
	return true, res.Error
}

func (a *authRepository) CheckToken(token string) (bool, error) {
	var b *models.Blacklist
	a.db.Table("blacklists").Where("token=?", token).First(&b)
	if b.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (a *authRepository) CleanBlacklists() (bool, error) {
	var b *models.Blacklist
	a.db.Table("blacklists").Unscoped().Where("id > 0").Delete(&b)
	return true, nil
}
