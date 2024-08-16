package services

import (
	"fmt"
	"project2/internal/domain/entities"
	"project2/pkg/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Signup(user *entities.User) error {
	// some basic manipulations
	user.Role = "user"
	uuid, err := utils.GetUuid()
	if err != nil {
		fmt.Println("Error generating password hash: ", err)
		return err
	}
	user.UserId = uuid
	password, err := utils.GetHashedPassword(user.Password)
	if err != nil {
		fmt.Println("Error generating password hash: ", err)
		return err
	}
	user.Password = password

	// logic to store the entry in the local map and the users.json file

	return nil
}
