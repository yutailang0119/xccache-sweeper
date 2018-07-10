package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/yutailang0119/go-xccache-sweeper/sources/archives"
	"github.com/yutailang0119/go-xccache-sweeper/sources/deriveddata"
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
	app.Version = "0.1.0"

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
	}

	err := app.Run(os.Args)

	if err != nil {
		panic(err)
	}
}
