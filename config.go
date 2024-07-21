package main

// Config struct to hold the configuration values
import (
    "github.com/sirupsen/logrus"
    "os"
)
type Config struct {
	PrivateKey      string
	ApplicationName string
	SubsystemName   string
}

// NewConfig function to load configuration from environment variables
func NewConfig() *Config {
	privateKey := os.Getenv("CORALOGIX_PRIVATE_KEY")
	if privateKey == "" {
		logrus.Fatal("CORALOGIX_PRIVATE_KEY environment variable is required")
	}

	applicationName := os.Getenv("CORALOGIX_APPLICATION_NAME")
	if applicationName == "" {
		logrus.Fatal("CORALOGIX_APPLICATION_NAME environment variable is required")
	}

	subsystemName := os.Getenv("CORALOGIX_SUBSYSTEM_NAME")
	if subsystemName == "" {
		logrus.Fatal("CORALOGIX_SUBSYSTEM_NAME environment variable is required")
	}

	return &Config{
		PrivateKey:      privateKey,
		ApplicationName: applicationName,
		SubsystemName:   subsystemName,
	}
}