package settings

import (
	_ "embed"
	"encoding/json"
)

//go:embed settings.json
var settingsFile []byte

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Settings struct {
	Port     int      `json:"port"`
	DBConfig Database `json:"database"`
}

func New() (*Settings, error) {
	var s Settings

	err := json.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
