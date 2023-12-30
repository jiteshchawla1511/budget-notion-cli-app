package main

import (
	"fmt"
	"os"

	"github.com/jiteshchawla1511/budget-notion-cli-app/cli"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading the env file %v", err)
	}
	cli.RootCmd.Run = cli.AddInteractive
	err = cli.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
