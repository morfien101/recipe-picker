package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	From           string   `json:"from"`
	Password       string   `json:"password"`
	Prefix         string   `json:"file_prefix"`
	Location       string   `json:"file_location"`
	RecipeCount    int      `json:"recipes_required"`
	EmailAddresses []string `json:"email_to"`
}

func newConfig(location string) (*config, error) {
	b, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, fmt.Errorf("Failed to read configuration file. Error: %s", err)
	}
	config := &config{}
	if err := json.Unmarshal(b, config); err != nil {
		return nil, fmt.Errorf("Failed to marshal configuration file. Error: %s", err)
	}
	if err := validate(config); err != nil {
		return nil, err
	}
	return config, nil
}

func validate(config *config) error {
	if config.From == "" {
		return fmt.Errorf("from can not be blank")
	}
	if config.Password == "" {
		return fmt.Errorf("password can not be blank")
	}
	if config.Prefix == "" {
		return fmt.Errorf("file_prefix can not be blank")
	}
	if config.Location == "" {
		return fmt.Errorf("file_location can not be blank")
	}
	if config.RecipeCount == 0 {
		return fmt.Errorf("recipes_required can not be blank")
	}
	if len(config.EmailAddresses) < 1 {
		return fmt.Errorf("email_to can not be blank")
	}
	return nil
}
