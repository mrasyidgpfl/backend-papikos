package config

import (
	"os"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type AppConfig struct {
	ENV                string
	AppName            string
	JWTSecret          []byte
	JWTExpireInMinutes int64
	DBConfig           dbConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	AppName:            "PAPIKOS",
	JWTSecret:          []byte("very-secret"),
	JWTExpireInMinutes: 180,
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", "ec2-34-231-190-161.compute-1.amazonaws.com"),
		User:     getENV("DB_USER", "hukwibgmmlwluu"),
		Password: getENV("DB_PASSWORD", "7f2baf2a9f6cb9e283e36fe7a8db3d38566647fc7e51c805a2e2d5aaaf81e3ed"),
		DBName:   getENV("DB_NAME", "dab0h51o47bf12"),
		Port:     getENV("DB_PORT", "5432"),
	},
}
