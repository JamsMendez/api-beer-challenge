package settings

import (
	_ "embed"
	"encoding/json"
	"os"
	"strconv"
)

const (
	envProd = "production"
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
	s := &Settings{}
	var err error

	env := os.Getenv("GO_ENV")
	if env == envProd {
		s, err = setUpEnvironments()
		if err != nil {
			return nil, err
		}

		return s, nil
	}

	err = json.Unmarshal(settingsFile, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func setUpEnvironments() (*Settings, error) {
	p := os.Getenv("API_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}

	dbp := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbp)
	if err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")

	s := Settings{
		Port: port,
		DBConfig: Database{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPwd,
			Name:     dbName,
		},
	}

	return &s, nil
}
