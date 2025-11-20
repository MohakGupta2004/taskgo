package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Context struct {
	CurrentGroup  string            `json:"current_group"`
	GroupValidity map[string]string `json:"group_validity"`
}

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".taskgo", "context.json"), nil
}

func LoadContext() (*Context, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Context{
			CurrentGroup: "",
			GroupValidity: map[string]string{
				"General": "24h",
			},
		}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ctx Context
	if err := json.Unmarshal(data, &ctx); err != nil {
		return nil, err
	}

	return &ctx, nil
}

func SaveContext(ctx *Context) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
