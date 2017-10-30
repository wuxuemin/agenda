package agenda_service

import (
	"entity"
	"fmt"
	"utils"
)

type user = entity.User

func validateNewUser(user *user) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("empty username")
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return fmt.Errorf("username '%s' exits", user.Username)
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("empty password")
	}
	return nil
}

// Register registers a user
func Register(username string, password string, email string, telephone string) (err error) {
	newUser := &user{
		Username:  username,
		Password:  password,
		Email:     email,
		Telephone: telephone,
	}
	err = validateNewUser(newUser)
	if err != nil {
		return
	}
	newUser.Password = utils.sha(newUser.Password)
	entity.UserModel.AddUser(newUser)
	return
}

// checks whether provided password matches that of the user
func IsPwdMatch(user *user, password string) bool {
	return utils.sha(password) == user.Password
}
