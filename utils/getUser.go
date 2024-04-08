package utils

import (
	"os/user"
)

type User struct {
	Name     string
	Username string
}

func GetUser() *User {
	sysUser, err := user.Current()
	if err != nil {
		ErrorFormatFatal("Error getting user", err)
	}
	return &User{
		Name:     sysUser.Name,
		Username: sysUser.Username,
	}
}
