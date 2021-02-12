package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() map[string]string {
	godotenv.Load()
	envs := make(map[string]string)
	envs["PORT"] = os.Getenv("PORT")
	envs["POSTGRES_HOST"] = os.Getenv("POSTGRES_HOST")
	envs["POSTGRES_USER"] = os.Getenv("POSTGRES_USER")
	envs["POSTGRES_PASSWORD"] = os.Getenv("POSTGRES_PASSWORD")
	envs["POSTGRES_DB"] = os.Getenv("POSTGRES_DB")
	envs["POSTGRES_PORT"] = os.Getenv("POSTGRES_PORT")
	return envs
}
