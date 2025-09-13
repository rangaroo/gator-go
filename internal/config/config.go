package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := filepath.Join(userHomeDir, configFileName)
	return configFilePath, nil
}

func (c Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	err := write(c)
	if err != nil {
		return err
	}

	return nil
}

func write(config Config) error {
	// Convert Config struct to JSON format
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Open temporary config file
	tempConfigPath := "temp_config.json"
	f, err := os.OpenFile(tempConfigPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write(jsonData); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	// Replace the original config with temp config
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.Rename(tempConfigPath, configFilePath); err != nil {
		return err
	}

	return nil
}
