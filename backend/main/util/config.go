package util

import (
	"encoding/json"
	"os"
	"strings"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

type Config struct {
	ApiPort  string   `json:"apiPort"`
	DbConfig DbConfig `json:"dbConfig"`
}

func LoadConfig(configFileName string) (*Config, error) {
	// Read the env.json file
	file, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into a map[string]interface{}
	var placeholders map[string]interface{}
	err = json.Unmarshal(file, &placeholders)
	if err != nil {
		return nil, err
	}

	// Replace placeholders with environment variables
	replacedPlaceholders := replacePlaceholders(placeholders)

	// Marshal the updated placeholders back to JSON
	updatedJSON, err := json.Marshal(replacedPlaceholders)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into the Config struct
	var config Config
	err = json.Unmarshal(updatedJSON, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func replacePlaceholders(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = replacePlaceholders(val)
		}
	case string:
		if strings.HasPrefix(v, "{{.") && strings.HasSuffix(v, "}}") {
			envVar := strings.TrimPrefix(strings.TrimSuffix(v, "}}"), "{{.")
			replacedValue := os.Getenv(envVar)
			if replacedValue == "" {
				// If the environment variable is not set, keep the placeholder
				return v
			}
			return replacedValue
		}
	}
	return data
}
