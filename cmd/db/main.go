package main

import (
	"context"
	"log"
	"os"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/db"
	"github.com/samber/do"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{}

	cmd.Commands = []*cli.Command{
		{
			Name:  "up",
			Usage: "Apply all pending migrations",
			Action: func(ctx context.Context, c *cli.Command) error {
				migrator := do.MustInvoke[*db.Migrator](bootstrap.Injector)
				return migrator.Up(ctx)
			},
		},
		{
			Name:  "down",
			Usage: "Revert the last applied migration",
			Action: func(ctx context.Context, c *cli.Command) error {
				migrator := do.MustInvoke[*db.Migrator](bootstrap.Injector)
				return migrator.Down(ctx)
			},
		},
		{
			Name:  "create",
			Usage: "Create a new database",
			Action: func(ctx context.Context, c *cli.Command) error {
				return db.CreateSqlDatabase()
			},
		},
		{
			Name:  "drop",
			Usage: "Drop the database",
			Action: func(ctx context.Context, c *cli.Command) error {
				return db.DropSqlDatabase()
			},
		},
		{
			Name:  "seed",
			Usage: "Seed the database with initial data",
			Action: func(ctx context.Context, c *cli.Command) error {
				seeder := do.MustInvoke[*db.Seeder](bootstrap.Injector)
				return seeder.Seed(ctx)
			},
		},
		{
			Name:  "setup",
			Usage: "Setup the database",
			Action: func(ctx context.Context, c *cli.Command) error {
				err := db.CreateSqlDatabase()
				if err != nil {
					return err
				}

				migrator := do.MustInvoke[*db.Migrator](bootstrap.Injector)
				err = migrator.Up(ctx)
				if err != nil {
					return err
				}

				seeder := do.MustInvoke[*db.Seeder](bootstrap.Injector)
				return seeder.Seed(ctx)
			},
		},
		{
			Name:  "reset",
			Usage: "Reset the database",
			Action: func(ctx context.Context, c *cli.Command) error {
				err := db.DropSqlDatabase()
				if err != nil {
					return err
				}

				err = db.CreateSqlDatabase()
				if err != nil {
					return err
				}

				migrator := do.MustInvoke[*db.Migrator](bootstrap.Injector)
				err = migrator.Up(ctx)
				if err != nil {
					return err
				}

				seeder := do.MustInvoke[*db.Seeder](bootstrap.Injector)
				return seeder.Seed(ctx)
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln("Failed to run command:", err)
	}
}
