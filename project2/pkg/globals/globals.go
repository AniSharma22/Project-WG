package globals

import (
	"project2/internal/domain/entities"
)

var UsersMap = make(map[string]entities.User) // uuid : User
var GamesMap = make(map[string]entities.Game)
var ResultsMap = make(map[string]entities.Result)
var SlotsMap = make(map[string]map[string][]entities.SlotStats) //key is game id
var ActiveUser string
