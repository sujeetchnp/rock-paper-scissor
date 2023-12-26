package service

import (
	"github.com/sujeetchnp/rock-paper-scissor/model"
)

// CacheService interface
type CacheService interface {
	CacheResult(result model.GameResult)
	GetResultFromCache(player1Choice, player2Choice model.Choice) *model.GameResult
}

// CacheServiceImpl struct
type CacheServiceImpl struct {
	ResultMap map[string]model.GameResult
}

func NewCacheServiceImpl() *CacheServiceImpl {
	resultMap := make(map[string]model.GameResult)
	instance := CacheServiceImpl{ResultMap: resultMap}
	instancePointer := &instance
	return instancePointer

	// return &CacheServiceImpl{ResultMap: resultMap}   -- putting all steps together above
}

func (cs *CacheServiceImpl) CacheResult(result model.GameResult) {
	cacheKey := getCacheKey(result.Player1Choice, result.Player2Choice)
	cs.ResultMap[cacheKey] = result
}

func (cs *CacheServiceImpl) GetResultFromCache(player1Choice, player2Choice model.Choice) *model.GameResult {
	cacheKey := getCacheKey(player1Choice, player2Choice)
	if result, ok := cs.ResultMap[cacheKey]; ok {
		return &result
	}
	return nil
}

func getCacheKey(player1Choice, player2Choice model.Choice) string {
	return player1Choice.String() + player2Choice.String()
}

// GameService interface
type GameService interface {
	ComparePlayerChoice(player1Choice, player2Choice model.Choice) model.Choice
}

// GameServiceImpl struct
type GameServiceImpl struct{}

func NewGameServiceImpl() *GameServiceImpl {
	return &GameServiceImpl{}
}

func (gs *GameServiceImpl) ComparePlayerChoice(player1Choice, player2Choice model.Choice) model.Choice {
	if player1Choice == player2Choice {
		return 0 // Draw
	}
	// Logic for determining the winner

	switch player1Choice {
	case model.ROCK:
		if player2Choice == model.SCISSOR {
			return model.ROCK // Rock beats Scissors
		}
	case model.PAPER:
		if player2Choice == model.ROCK {
			return model.PAPER // Paper beats Rock
		}
	case model.SCISSOR:
		if player2Choice == model.PAPER {
			return model.SCISSOR // Scissors beats Paper
		}
	}

	// If none of the above conditions are met, player2 wins
	return player2Choice

}
