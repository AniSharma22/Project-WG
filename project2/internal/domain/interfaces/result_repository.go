package interfaces

import "project2/internal/domain/entities"

type ResultRepository interface {
	AddResult(result *entities.Result) error
	RemoveResult(resultId string) error
	FindResult(resultId string) *entities.Result
	GetAllResults() []*entities.Result
}
