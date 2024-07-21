package main

import (
	"os"
	"testing"

	"github.com/coralogix/go-coralogix-sdk"
	"github.com/sirupsen/logrus"
)

// Helper function to set environment variables for the test
func setEnv(key, value string) error {
	return os.Setenv(key, value)
}

// Helper function to unset environment variables after the test
func unsetEnv(key string) error {
	return os.Unsetenv(key)
}

func TestNewConfig(t *testing.T) {
	err := setEnv("CORALOGIX_PRIVATE_KEY", "test_private_key")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	err = setEnv("CORALOGIX_APPLICATION_NAME", "test_application_name")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	err = setEnv("CORALOGIX_SUBSYSTEM_NAME", "test_subsystem_name")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	defer func() {
        err := unsetEnv("CORALOGIX_PRIVATE_KEY")
        if err != nil {
            logrus.Errorln("error unsetting CORALOGIX_PRIVATE_KEY")
        }
    }()
	defer func() {
        err := unsetEnv("CORALOGIX_APPLICATION_NAME")
        if err != nil {
			logrus.Errorln("error unsetting CORALOGIX_APPLICATION_NAME")
        }
    }()
	defer func() {
        err := unsetEnv("CORALOGIX_SUBSYSTEM_NAME")
        if err != nil {
			logrus.Errorln("error unsetting CORALOGIX_SUBSYSTEM_NAME")
        }
    }()

	config := NewConfig()

	if config.PrivateKey != "test_private_key" {
		t.Errorf("Expected PrivateKey to be 'test_private_key', got '%s'", config.PrivateKey)
	}
	if config.ApplicationName != "test_application_name" {
		t.Errorf("Expected ApplicationName to be 'test_application_name', got '%s'", config.ApplicationName)
	}
	if config.SubsystemName != "test_subsystem_name" {
		t.Errorf("Expected SubsystemName to be 'test_subsystem_name', got '%s'", config.SubsystemName)
	}
}

func TestCoralogixHookInitialization(t *testing.T) {
	err := setEnv("CORALOGIX_PRIVATE_KEY", "test_private_key")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	err = setEnv("CORALOGIX_APPLICATION_NAME", "test_application_name")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	err = setEnv("CORALOGIX_SUBSYSTEM_NAME", "test_subsystem_name")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	defer func() {
        err := unsetEnv("CORALOGIX_PRIVATE_KEY")
        if err != nil {
            logrus.Errorln("error unsetting CORALOGIX_PRIVATE_KEY")
        }
    }()
	defer func() {
        err := unsetEnv("CORALOGIX_APPLICATION_NAME")
        if err != nil {
			logrus.Errorln("error unsetting CORALOGIX_APPLICATION_NAME")
        }
    }()
	defer func() {
        err := unsetEnv("CORALOGIX_SUBSYSTEM_NAME")
        if err != nil {
			logrus.Errorln("error unsetting CORALOGIX_SUBSYSTEM_NAME")
        }
    }()

	config := NewConfig()
	CoralogixHook := coralogix.NewCoralogixHook(
		config.PrivateKey,
		config.ApplicationName,
		config.SubsystemName,
	)
	defer CoralogixHook.Close()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.AddHook(CoralogixHook)

	entry := logrus.NewEntry(log)
	err = CoralogixHook.Fire(entry)
	if err != nil {
		t.Errorf("Expected no error from CoralogixHook.Fire, got '%v'", err)
	}
}
