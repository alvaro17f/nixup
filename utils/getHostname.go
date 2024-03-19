package utils

import (
	"log"
	"os"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}
	return hostname
}
