package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/yutailang0119/go-xccache-sweeper/lib/archives"
	"github.com/yutailang0119/go-xccache-sweeper/lib/deriveddata"
	"github.com/yutailang0119/go-xccache-sweeper/lib/devicesupport"
)

var (
	// Version is git tag version from Makefile `shell git describe --tags --abbrev=0`
	Version string
	// Revision is git HEAD revision from Makefile `shell git rev-parse --short HEAD`
	Revision string
)

func main() {

	app := cli.NewApp()
	app.Name = "xccache-sweeper"
	app.Usage = "Sweep Xcode caches"
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:  "archives",
			Usage: "Sweep Archives. Defaults is /Users/user/Library/Developer/Xcode/Archives",
			Action: func(c *cli.Context) error {
				return archives.SweepArchives()
			},
		},
		{
			Name:  "deriveddata",
			Usage: "Sweep DerivedData. Defaults is /Users/user/Library/Developer/Xcode/DerivedData",
			Action: func(c *cli.Context) error {
				return deriveddata.SweepDerivedData()
			},
		},
		{
			Name:  "caches",
			Usage: "Sweep Archives and DerivedData.",
			Action: func(c *cli.Context) error {
				err := deriveddata.SweepDerivedData()
				if err != nil {
					return err
				}

				err = archives.SweepArchives()
				return err
			},
		},
		{
			Name:  "devicesupport",
			Usage: "Sweep Device Support. ~/Library/Developer/Xcode/*DeviceSupport",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "force delete all",
				},
			},
			Action: func(c *cli.Context) error {
				err := devicesupport.SweepDeviceSupports(c.Args().First(), c.Bool("all"))
				return err
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
