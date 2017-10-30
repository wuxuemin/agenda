package service

import (
	"entity"
	"fmt"
)

var (
	// ErrInvalidCredentials is error for incorrent username and password combination
	ErrInvalidCredentials = fmt.Errorf("incorrent username or password")
)

// Login logs in with provided user
func Login(username string, password string) error {
	curUsername := entity.CurSessionModel.GetCurUser()
	if len(curUsername) > 0 {
		return fmt.Errorf("you have already logged in as '%s'. Please logout first", curUsername)
	}
	user := entity.UserModel.FindByUsername(username)
	if user == nil || IsPwdMatch(user, password) == false {
		return ErrInvalidCredentials
	}
	entity.CurSessionModel.SetCurUser(user)
	return nil
}

// Logout logs out
func Logout() error {
	err := checkIfLoggedin()
	if err != nil {
		return err
	}
	entity.CurSessionModel.SetCurUser(&user{})
	return nil
}

func checkIfLoggedin() error {
	curUsername := entity.CurSessionModel.GetCurUser()
	if len(curUsername) == 0 {
		return fmt.Errorf("please login first")
	}
	return nil
}
