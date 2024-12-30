package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ConfigValue interface{}

type Config struct {
	options map[string]ConfigValue
}

func (c *Config) Set(key string, value ConfigValue) {
	c.options[key] = value
}

func (c *Config) Get(key string) ConfigValue {
	return c.options[key]
}

func (c *Config) Expect(key string, default_value ConfigValue) {
	if _, exists := c.options[key]; !exists {
		c.options[key] = default_value
	}
}

func (c *Config) GetAsInt(key string) (int, bool) {
	val, ok := c.options[key].(int)
	return val, ok
}

func (c *Config) GetAsString(key string) (string, bool) {
	val, ok := c.options[key].(string)
	return val, ok
}

func (c *Config) GetAsBool(key string) (bool, bool) {
	val, ok := c.options[key].(bool)
	return val, ok
}

func (c *Config) Load() error {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		log.Println("[Config] No configuration found, creating a default one.")

		if err := c.Save(); err != nil {
			return fmt.Errorf("failed to save default config: %v", err)
		}
		return nil
	}

	configJson, err := os.ReadFile("config.json")

	if err != nil {
		return fmt.Errorf("failed to read the config file: %v", err)
	}

	if err := json.Unmarshal(configJson, &c.options); err != nil {
		return fmt.Errorf("failed to unmarshall config: %v", err)
	}

	return nil
}

func (c *Config) Save() error {
	configJson, err := json.MarshalIndent(c.options, "", "  ")

	if err != nil {
		return fmt.Errorf("failed to marshall config during saving: %v", err)
	}

	if err := os.WriteFile("config.json", configJson, 0644); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		options: make(map[string]ConfigValue),
	}
}
