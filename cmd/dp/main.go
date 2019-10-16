package main

import (
	"log"
	"os"

	configCmd "github.com/ONSdigital/dp/cmd/commands/config"
	"github.com/ONSdigital/dp/cmd/commands/remote"
	"github.com/ONSdigital/dp/cmd/commands/ssh"
	"github.com/ONSdigital/dp/cmd/commands/ui"
	"github.com/ONSdigital/dp/cmd/config"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "dp"
	app.Version = "0.2.0"
	app.Usage = "digital publishing helper command"

	app.Commands = append(
		ssh.Command(cfg),
		ui.Command(cfg),
		remote.Command(cfg),
		configCmd.Command(cfg),
	)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
