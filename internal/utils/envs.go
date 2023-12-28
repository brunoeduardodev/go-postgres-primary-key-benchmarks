package utils

import "github.com/joho/godotenv"

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
