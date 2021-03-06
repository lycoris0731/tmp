package main

import "github.com/urfave/cli"

var commands = []cli.Command{
	{
		Name:    "make",
		Aliases: []string{"m", "mk"},
		Usage:   "Make a new temporary directory",
		Action:  makeDir,
	},
	{
		Name:    "remove",
		Aliases: []string{"r", "rm"},
		Usage:   "Remove a target directory as passed by an argument",
		Action:  removeDir,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "all",
				Usage: "Remove `all` directories in tracking",
			},
		},
	},
	{
		Name:    "list",
		Aliases: []string{"l", "ls"},
		Usage:   "Show all tracking temporary directories",
		Action:  list,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "number, n",
				Usage: "Show identication `number`",
			},
		},
	},
	{
		Name:    "move",
		Aliases: []string{"mv"},
		Usage:   "Move temporary directory",
		Action:  move,
	},
}
