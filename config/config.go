package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	SMTPDef SMTPConfig `yaml:"smtpDefinition"`
	EmailDef EmailConfig `yaml:"emailDefinition"`
}

// SMTPConfig holds SMTP server configuration
type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	FromEmail string `yaml:"from_email"`
	FromName  string `yaml:"from_name"`
}

// EmailConfig holds email-related configuration
type EmailConfig struct {
	SubjectPrefix  string `yaml:"subject_prefix"`
	RetryCount     int    `yaml:"retry_count"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
}

// LoadConfig loads configuration from the  YAML file
func LoadConfig(configPath string) (*Config, error) {
	// Read the config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse the YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}