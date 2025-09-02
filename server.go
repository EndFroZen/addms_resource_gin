package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int `json:"port"`
}

func main() {
	// โหลด config.json
	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(fmt.Errorf("failed to read config.json: %w", err))
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(fmt.Errorf("failed to parse config.json: %w", err))
	}

	if config.Port == 0 {
		config.Port = 3000
	}

	r := gin.Default()

	r.Static("/views", "./views")
	r.Static("/css", "./css")
	r.Static("/js", "./script")

	r.GET("/", func(c *gin.Context) {
		c.File("./views/index.html")
	})

	addr := fmt.Sprintf(":%d", config.Port)
	if err := r.Run(addr); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
