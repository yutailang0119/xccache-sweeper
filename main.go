package main

import (
	"github.com/yutailang0119/go-xccache-sweeper/sources/deriveddata"
)

var (
	// Version is git tag version from Makefile `shell git describe --tags --abbrev=0`
	Version string
	// Revision is git HEAD revision from Makefile `shell git rev-parse --short HEAD`
	Revision string
)

func main() {
	err := deriveddata.SweepDerivedData()

	if err != nil {
		panic(err)
	}
}
