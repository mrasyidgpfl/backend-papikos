package services

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"fmt"
	"math/rand"
	"time"
)

type GamesService interface {
	FlipCoin(userId uint, input *dto.FlipCoinsRequest) (*dto.FlipCoinsResponse, error)
}

type gamesService struct {
	gamesRepository repositories.GamesRepository
	appConfig       config.AppConfig
}

type GamesConfig struct {
	GamesRepository repositories.GamesRepository
	AppConfig       config.AppConfig
}

func NewGamesService(c *GamesConfig) GamesService {
	return &gamesService{
		gamesRepository: c.GamesRepository,
		appConfig:       c.AppConfig,
	}
}

func (g gamesService) FlipCoin(userId uint, input *dto.FlipCoinsRequest) (*dto.FlipCoinsResponse, error) {
	result := "MIN"
	prizes := []int{5000, 10000, 15000, 20000, 30000, 50000}
	prize := 0
	coin := []string{"heads", "tails"}
	rand.Seed(time.Now().UnixNano())
	side := coin[rand.Intn(len(coin))]
	fmt.Println(side)
	if side == input.CoinSide {
		result = "MAX"
		prize += prizes[rand.Intn(len(prizes))]
	}
	userChances, _ := g.gamesRepository.UpdateUserChances(userId)
	flipCoinsData := dto.FlipCoinsData{
		Chance:     userChances,
		CoinSide:   input.CoinSide,
		GameResult: result,
		Prize:      prize,
	}
	wallet, _ := g.gamesRepository.UpdateUserWallet(userId, prize)
	return new(dto.FlipCoinsResponse).CreateCoinsResponse(200, "COIN FLIP SUCCESS", "Games successfully executed and results are updated", flipCoinsData, wallet), nil
}
