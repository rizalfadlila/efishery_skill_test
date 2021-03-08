package main

import (
	"github.com/fetch_app/boostrap"
	"github.com/fetch_app/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Panic(err)
	}
}

func main() {
	boostrap.RunApp()
}
