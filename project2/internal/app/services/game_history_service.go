package services

import (
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/globals"
	"sync"
)

type GameHistoryService struct {
	gameHistoryRepo interfaces.GameHistoryRepository
	userService     *UserService
	gameHistoryWG   *sync.WaitGroup
}

func NewGameHistoryService(gameHistoryRepo interfaces.GameHistoryRepository, userService *UserService) *GameHistoryService {
	return &GameHistoryService{
		gameHistoryRepo: gameHistoryRepo,
		userService:     userService,
		gameHistoryWG:   &sync.WaitGroup{},
	}
}

func (gh *GameHistoryService) GetTotalGameHistory() ([]entities.GameHistory, error) {
	user, err := gh.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return nil, err
	}

	return gh.gameHistoryRepo.GetUserGameHistory(user.ID)
}

func (gh *GameHistoryService) GetResultsToUpdate() ([]entities.GameHistory, error) {
	user, err := gh.userService.GetUserByEmail(globals.ActiveUser)
	if err != nil {
		return nil, err
	}

	return gh.gameHistoryRepo.GetResultsToUpdate(user.ID)
}

//func (r *GameHistoryService) AddResult(result *entities.Result) error {
//	result.ResultId, _ = utils.GetUuid()
//	for _, user := range result.WinningUser {
//		err := r.userService.AddWinToUser(&user, result.GameId)
//		if err != nil {
//			return err
//		}
//	}
//	for _, user := range result.LosingUser {
//		err := r.userService.AddLossToUser(&user, result.GameId)
//		if err != nil {
//			return err
//		}
//	}
//	return r.resultRepo.AddResult(result)
//}
//
//func (r *GameHistoryService) GetResult(resultId string) (*entities.Result, error) {
//	result := r.resultRepo.FindResult(resultId)
//	if result == nil {
//		return nil, errors.New("result not found")
//	}
//	return result, nil
//}
//func (r *GameHistoryService) GetAllResults() ([]*entities.Result, error) {
//	results := r.resultRepo.GetAllResults()
//	if results == nil {
//		return nil, errors.New("no result found")
//	}
//	return results, nil
//}
//
