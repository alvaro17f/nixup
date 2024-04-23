package utils

import (
	"os"
	"os/user"

	"github.com/alvaro17f/nixup/internal/errors"
)

type User struct {
	Name     string
	Username string
}

func GetUser() *User {
	sysUser, err := user.Current()
	if err != nil {
		errors.ErrorFormatFatal("Error getting user", err)
	}
	return &User{
		Name:     sysUser.Name,
		Username: sysUser.Username,
	}
}

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		errors.ErrorFormatFatal("Error getting hostname", err)
	}
	return hostname
}
