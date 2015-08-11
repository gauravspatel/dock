package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func average(xs []float64) float64 {
	panic("Not Implemented")
}

func confirm(question string) bool {
	fmt.Printf("%s [y/N] ", question)
	var confirm string
	_, err := fmt.Scanln(&confirm)

	if err != nil {
		return false
	}

	if confirm == "y" || confirm == "Y" || confirm == "Yes" || confirm == "yes" {
		return true
	}

	return false
}

func main() {
	app := cli.NewApp()
	app.Name = "Dock"
	app.Usage = "Tame the docker inside of you."
	app.Version = "0.0.1"
	app.Author = "Gaurav Patel"
	app.Copyright = "2015 Gaurav Patel"
	app.Commands = []cli.Command{
		{
			Name:    "clean",
			Aliases: []string{"reset"},
			Usage:   "Clean the local environment of all docker images + containers",
			Before: func(c *cli.Context) error {
				var confirmReset = confirm("Are you sure you want to reset your environment?")

				if confirmReset {
					return nil
				}

				fmt.Printf("%s", "Clean operation was aborted\n")
				return errors.New("Clean operation was aborted")
			},
			Action: func(c *cli.Context) {
				exec.Command("/bin/sh", "-c", "docker stop $(docker ps -a -q) & docker rm $(docker ps -a -q) & docker rmi 0f $(docker images -q)").Output()

				fmt.Printf("All docker containers stopped and deleted and all images deleted\n")
			},
		},
	}

	app.EnableBashCompletion = true

	app.Run(os.Args)
}
