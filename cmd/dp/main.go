package main

import (
	"log"
	"os"

	"github.com/ONSdigital/dp/cmd/commands/ssh"
	"github.com/ONSdigital/dp/cmd/commands/ui"
	"github.com/ONSdigital/dp/cmd/config"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "dp"
	app.Usage = "digital publishing helper command"

	app.Commands = []cli.Command{
		ssh.Command(cfg),
		ui.Command(cfg),
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
