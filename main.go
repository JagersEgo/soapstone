package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"soapstone/database_controller"
	"soapstone/printer"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "soapstone",
		Usage: "Interact with the soapstone database and execute commands.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "noex",
				Usage: "Do not execute commands from system path",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "like",
				Usage:   "Like a message",
				Aliases: []string{"up"},
				Action: func(c *cli.Context) error {
					arg := strings.Join(c.Args().Slice(), " ")
					database_controller.VoteMessage(arg, true)
					return nil
				},
			},
			{
				Name:    "dislike",
				Usage:   "Dislike a message",
				Aliases: []string{"down"},
				Action: func(c *cli.Context) error {
					arg := strings.Join(c.Args().Slice(), " ")
					database_controller.VoteMessage(arg, false)
					return nil
				},
			},
			{
				Name:    "comment",
				Usage:   "Add comment to command, [COMMAND] [COMMENT...]",
				Aliases: []string{"add", "c", "a"},
				Action: func(c *cli.Context) error {
					if len(c.Args().Slice()) < 2 {
						printer.PrintWarning("Not enough arguments")
						return nil
					}

					msg := strings.Join(c.Args().Slice()[1:], " ")
					cmd := (c.Args().Slice()[0])

					database_controller.NewMessage(msg, cmd)

					database_controller.Save()
					return nil
				},
			},
			{
				Name:    "delete",
				Usage:   "Remove comment [COMMENT] [INDEX]",
				Aliases: []string{"remove"},
				Action: func(c *cli.Context) error {
					if len(c.Args().Slice()) < 2 {
						printer.PrintWarning("Not enough arguments")
						return nil
					}

					cmd := (c.Args().Slice()[0])
					idx, e := strconv.Atoi(c.Args().Slice()[1])

					if e != nil {
						printer.PrintWarning("Argument 1 is not a valid integer")
						return nil
					}

					database_controller.RemoveMessage(cmd, idx)

					database_controller.Save()
					return nil
				},
			},
			{
				Name:  "help",
				Usage: "Show help information",
				Action: func(c *cli.Context) error {
					print_help()
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			if os.Geteuid() == 0 {
				printer.PrintWarning("Don't run me as root!")
				return nil
			}

			if c.Args().Len() == 0 {
				fmt.Println("No command specified.")
				print_help()
				return nil
			}

			cmd := c.Args().Get(0)
			args := strings.Join(c.Args().Slice()[1:], " ")

			if c.Bool("noex") {
				database_controller.PrintSoapstone(cmd)
				return nil
			}

			database_controller.PrintSoapstone(cmd)

			command := exec.Command(cmd, args)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Stdin = os.Stdin

			if err := command.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}

			database_controller.Save()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

func print_help() {
	fmt.Println("<> Helpful information <>")
}
