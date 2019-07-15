package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	SMTPUsername   string   `json:"smtp_username"`
	SMTPPassword   string   `json:"smtp_password"`
	Prefix         string   `json:"file_prefix"`
	Location       string   `json:"file_location"`
	RecipeCount    int      `json:"recipes_required"`
	From           string   `json:"from"`
	EmailAddresses []string `json:"email_to"`
	SendEmail      bool     `json:"send_email"`
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
	if config.SMTPUsername == "" {
		return fmt.Errorf("smtp_username can not be blank")
	}
	if config.SMTPPassword == "" {
		return fmt.Errorf("smtp_password can not be blank")
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

func exampleConfig() (string, error) {
	example := &config{
		SMTPUsername:   "username@test.com",
		SMTPPassword:   "test123",
		Prefix:         "https://fileserver.local/recipes",
		Location:       "/var/recipes",
		RecipeCount:    5,
		From:           "head_chef@yummy.com",
		EmailAddresses: []string{"candy@test.com", "sandy@test.com", "andy@test.com"},
		SendEmail:      true,
	}

	output, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}
