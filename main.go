package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
	cli "github.com/urfave/cli/v2"
	"github.com/ztrue/tracerr"
)

func main() {
	app := &cli.App{
		Name:  "git-auto-sync",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			{
				Name:  "watch",
				Usage: "Watch a folder for changes",
				Action: func(ctx *cli.Context) error {
					repoPath, err := os.Getwd()
					if err != nil {
						return tracerr.Wrap(err)
					}

					return WatchForChanges(repoPath)
				},
			},
			{
				Name:  "sync",
				Usage: "Sync a repo right now",
				Action: func(ctx *cli.Context) error {
					repoPath, err := os.Getwd()
					if err != nil {
						return tracerr.Wrap(err)
					}

					err = autoSync(repoPath)
					if err != nil {
						return tracerr.Wrap(err)
					}

					return nil
				},
			},
			{
				Name:  "notify",
				Usage: "Sync a repo right now",
				Action: func(ctx *cli.Context) error {
					err := beeep.Alert("Title", "Message body", "assets/warning.png")
					if err != nil {
						panic(err)
					}

					return nil
				},
			},
			{
				Name:  "daemon",
				Usage: "Interact with the background daemon",
				Subcommands: []*cli.Command{
					{
						Name:   "status",
						Usage:  "Show the Daemon's status",
						Action: daemonStatus,
					},
					{
						Name:    "list",
						Aliases: []string{"ls"},
						Usage:   "List of repos being auto-synced",
						Action:  daemonList,
					},
					{
						Name:   "add",
						Usage:  "Add a repo for auto-sync",
						Action: daemonAdd,
					},

					{
						Name:    "remove",
						Aliases: []string{"rm"},
						Usage:   "Remove a repo from auto-sync",
						Action:  daemonRm,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
