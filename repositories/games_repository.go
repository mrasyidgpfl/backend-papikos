package repositories

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"gorm.io/gorm"
)

type GamesRepository interface {
	UpdateUserChances(userId uint) (int, error)
	UpdateUserWallet(userId uint, prize int) (*models.Wallet, error)
}

type gamesRepository struct {
	db *gorm.DB
}

type GamesConfig struct {
	DB *gorm.DB
}

func NewGamesRepository(b *GamesConfig) GamesRepository {
	return &gamesRepository{db: b.DB}
}

func (g gamesRepository) UpdateUserChances(userId uint) (int, error) {
	var gC *models.GamesChances
	g.db.Table("games_chances").Where("user_id = ?", userId).First(&gC)
	g.db.Table("games_chances").Where("user_id = ?", userId).Update("chance", gC.Chance-1)
	g.db.Table("games_chances").Where("user_id = ?", userId).First(&gC)
	g.db.Table("games_chances").Where("user_id = ?", userId).Update("history", gC.History+1)
	g.db.Table("games_chances").Where("user_id = ?", userId).First(&gC)
	if gC.History > 0 && gC.History%10 == 0 {
		g.db.Table("games_chances").Where("user_id = ?", userId).Update("chance", gC.Chance+1)
	}
	g.db.Table("games_chances").Where("user_id = ?", userId).First(&gC)
	return gC.Chance, nil
}

func (g gamesRepository) UpdateUserWallet(userId uint, prize int) (*models.Wallet, error) {
	var wallet *models.Wallet
	g.db.Table("wallets").Where("user_id = ?", userId).First(&wallet)
	g.db.Table("wallets").Where("user_id = ?", userId).Update("balance", wallet.Balance+prize)
	g.db.Table("wallets").Where("user_id = ?", userId).First(&wallet)
	return wallet, nil
}
