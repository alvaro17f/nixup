package utils

import (
	"os"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		ErrorFormatFatal("Error getting hostname", err)
	}
	return hostname
}
