package main

import (
    "flag"
    "fmt"
    coralogix "github.com/coralogix/go-coralogix-sdk"
    "github.com/sirupsen/logrus"
    yaml "gopkg.in/yaml.v3"
    "io/ioutil"
)

type AlertConfig struct {
	Name           string `yaml:"name"`
	Query          string `yaml:"query"`
	Severity       string `yaml:"severity"`
	Notification   string `yaml:"notification"`
	Threshold      int    `yaml:"threshold"`
	TimeWindow     string `yaml:"time_window"`
	NotificationID int    `yaml:"notification_id"`
}

// LoadAlertConfig function to load the alert configuration from a YAML file
func LoadAlertConfig(filePath string) ([]AlertConfig, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var alerts []AlertConfig
	err = yaml.Unmarshal(data, &alerts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return alerts, nil
}

func main()  {
	
	// Get the path to the alerts.yaml file from command line arguments
	alertsFilePath := flag.String("alerts-file", "alerts.yaml", "Path to the alerts YAML file")
	flag.Parse()
	
    config := NewConfig()
    
    // Load alert configurations from YAML file
	alerts, err := LoadAlertConfig(*alertsFilePath)
	if err != nil {
		logrus.Fatalf("Error loading alert configurations: %v", err)
	}
    
	CoralogixHook := coralogix.NewCoralogixHook(
        config.PrivateKey,
        config.ApplicationName,
        config.SubsystemName,
    )
    defer CoralogixHook.Close()

    log := logrus.New()
    log.SetLevel(logrus.DebugLevel)

    log.AddHook(CoralogixHook)
    
    // Create alerts based on loaded configuration
	for _, alert := range alerts {
		// Use Coralogix SDK to create alert
		log.Infof("Creating alert: %s", alert.Name)
		// This part is hypothetical as the actual API call depends on the SDK's capabilities
		// coralogix.CreateAlert(alert.Name, alert.Query, alert.Severity, alert.Notification, alert.Threshold, alert.TimeWindow, alert.NotificationID)
	}
    
}