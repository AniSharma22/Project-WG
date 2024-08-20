package services

import (
	"errors"
	"project2/internal/domain/entities"
	"project2/internal/domain/interfaces"
	"project2/pkg/utils"
	"sync"
)

type ResultService struct {
	resultRepo  interfaces.ResultRepository
	userService *UserService
	resultWG    *sync.WaitGroup
}

func NewResultService(resultRepo interfaces.ResultRepository, userService *UserService) *ResultService {
	return &ResultService{
		resultRepo:  resultRepo,
		userService: userService,
		resultWG:    &sync.WaitGroup{},
	}
}

func (r *ResultService) AddResult(result *entities.Result) error {
	result.ResultId, _ = utils.GetUuid()
	for _, user := range result.WinningUser {
		err := r.userService.AddWinToUser(&user, result.GameId)
		if err != nil {
			return err
		}
	}
	for _, user := range result.LosingUser {
		err := r.userService.AddLossToUser(&user, result.GameId)
		if err != nil {
			return err
		}
	}
	return r.resultRepo.AddResult(result)
}

func (r *ResultService) GetResult(resultId string) (*entities.Result, error) {
	result := r.resultRepo.FindResult(resultId)
	if result == nil {
		return nil, errors.New("result not found")
	}
	return result, nil
}
func (r *ResultService) GetAllResults() ([]*entities.Result, error) {
	results := r.resultRepo.GetAllResults()
	if results == nil {
		return nil, errors.New("no result found")
	}
	return results, nil
}

//func (r *ResultService) RemoveResult(resultId string) error {
//	return r.resultRepo.RemoveResult(resultId)
//}
