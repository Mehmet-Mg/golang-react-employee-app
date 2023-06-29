package main

import (
	"example/employee-app/app"
	"example/employee-app/config"
	"fmt"
	"os"
)

const dev = "development"

func main() {
	cfg, err := config.NewConfig(dev)
	if err != nil {
		fmt.Println("Error reading config file, ", err)
		os.Exit(1)
	}
	app.Run(cfg)
}
