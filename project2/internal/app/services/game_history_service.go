package services

import (
	"project2/internal/domain/interfaces"
	"sync"
)

type GameHistoryService struct {
	gameHistoryRepo interfaces.GameHistoryRepository
	gameHistoryWG   *sync.WaitGroup
}

func NewGameHistoryService(gameHistoryRepo interfaces.GameHistoryRepository) *GameHistoryService {
	return &GameHistoryService{
		gameHistoryRepo: gameHistoryRepo,
		gameHistoryWG:   &sync.WaitGroup{},
	}
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
