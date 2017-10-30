package service

import (
	"entity"
	"fmt"
	// "utils"
)

type user = entity.User

func validateNewUser(user *user) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("username should not be empty")
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return fmt.Errorf("username '%s' already exists", user.Username)
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("password should not be empty")
	}
	// ...TODO
	return nil
}

// Register registers a user
func Register(username string, password string, email string, phone string) (err error) {
	newUser := &user{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	err = validateNewUser(newUser)
	if err != nil {
		return
	}
	//newUser.Password = utils.MD5(newUser.Password)
	newUser.Password = newUser.Password
	entity.UserModel.AddUser(newUser)
	return
}

// IsPwdMatch checks whether provided password matches that of the user
func IsPwdMatch(user *user, password string) bool {
	//return utils.MD5(password) == user.Password
	return password == user.Password
}
