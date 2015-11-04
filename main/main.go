package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/masterlink720/go_backup_db/db"
	"os"
)

const VERSION = "0.0.1"

func parseArgs() *cli.App {

	app := cli.NewApp()
	app.Name = "Go DB Backup"
	app.Usage = "Quickly backups up a database"
	app.Version = VERSION

	app.Commands = []cli.Command{
		{
			Name:      "mysql",
			ShortName: "MySQL commands",
			Usage:     "Uses a MySQL database when executing",
			Flags
			Flags: []cli.Flag{
				cli.StringFlag{Name: "host", Value: "localhost", Usage: "[Optional] Name of the host providing MySQL"},
				cli.IntFlag{Name: "port p", Value: 3306, Usage: "[Optional] Name of the host providing MySQL"},
				cli.StringFlag{Name: "user, u", Value: "root", Usage: "[Optional] Username for accessing the database"},
				cli.StringFlag{Name: "database, d", Usage: "Databaes name to backup"},
				cli.StringSliceFlag{Name: "tables, t", Usage: "[Optional]: Table names to include"},
			},
			Action: func(c *cli.Context) {
				cli.ShowAppHelp(c)
				os.Exit(1)
			},
			Subcommands: []cli.Command{
				{
					Name:    "backup",
					HideHelp:  true,
					Aliases: []string{"b"},
					Usage:   "Create a backup of a MySQL database",
					Flags: []cli.Flag{
						cli.StringSliceFlag{Name: "tables, t", Usage: "[Optional]: Table names to include"},
					},
					Action: mysqlBackup,
				},
			},
		},
	}

	// Global flags
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "output, o", Value: "", Usage: "Filename to import / export"},
	}

	return app
}

func main() {
	app := parseArgs()
	app.Run(os.Args)
}

func mysqlBackup(c *cli.Context) {

	// Build the backup options and verify
	opts := db.MysqlOptions{
		Host:   c.String("host"),
		Port:   c.Int("port"),
		User:   c.String("user"),
		Db:     c.String("database"),
		Tables: c.StringSlice("tables"),
	}

	if opts.Host == "" || opts.Port == 0 || opts.User == "" || opts.Db == "" {
		cli.ShowCommandHelp(c, "mysql") // "Host name, Port number, Username, and Database name are all required\n\tUnable to continue\n")
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	fmt.Printf("Backing up MySQL... ")
}
