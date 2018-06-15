package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	// Version is git tag version from Makefile `shell git describe --tags --abbrev=0`
	Version string
	// Revision is git HEAD revision from Makefile `shell git rev-parse --short HEAD`
	Revision string
)

func cachedDerivedDataPaths() ([]string, error) {
	xcodeBuildLocationStyle, err := exec.Command("defaults", "read", "com.apple.dt.Xcode", "IDEBuildLocationStyle").Output()
	if err != nil {
		xcodeBuildLocationStyle = []byte("Unique")
	}

	usr, _ := user.Current()
	if strings.TrimSpace(string(xcodeBuildLocationStyle)) == "Unique" {
		return []string{strings.Replace("~/Library/Developer/Xcode/DerivedData", "~", usr.HomeDir, 1)}, nil
	} else {
		paths := []string{}
		err := filepath.Walk(usr.HomeDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {

				if filepath.Ext(path) == ".xcodeproj" {
					cmd := "xcodebuild -project " + path
					cmd = cmd + " -showBuildSettings | grep -e \"BUILD_ROOT\""
					buildRoot, _ := exec.Command("sh", "-c", cmd).Output()
					output := strings.TrimSpace(string(buildRoot))

					buildRootPath := strings.TrimPrefix(output, "BUILD_ROOT = ")
					derivedDataPath := strings.TrimSuffix(buildRootPath, "/Build/Products")

					paths = append(paths, derivedDataPath)

					return filepath.SkipDir
				}

			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return paths, nil

	}
}

func cachedArchivesPath() string {

	usr, _ := user.Current()
	xcodePlistPath := strings.Replace("~/Library/Preferences/com.apple.dt.Xcode", "~", usr.HomeDir, 1)
	archivesPath, err := exec.Command("defaults", "read", xcodePlistPath, "IDECustomDistributionArchivesLocation").Output()
	if err != nil {
		return strings.Replace("~/Library/Developer/Xcode/Archives", "~", usr.HomeDir, 1)
	}
	return string(archivesPath)

}

func checkExpired(dir string, expired time.Time) (bool, error) {

	splited := strings.Split(dir, "-")
	year, err := strconv.Atoi(splited[0])
	if err != nil {
		return false, err
	}

	month, err := strconv.Atoi(splited[1])
	if err != nil {
		return false, err
	}

	day, err := strconv.Atoi(splited[2])
	if err != nil {
		return false, err
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return date.Before(expired), nil
}

func main() {

	derivedDataPaths, err := cachedDerivedDataPaths()
	if err != nil {
		panic(err)
	}

	for _, path := range derivedDataPaths {
		err := os.RemoveAll(path)
		if err == nil {
			log.Println(path)
		}
	}

	archivesPath := cachedArchivesPath()
	matchingArchivesPath := filepath.Join(archivesPath, "*")

	now := time.Now()
	expired := now.AddDate(0, -1, 0)

	err = filepath.Walk(archivesPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {

			ok, err := filepath.Match(matchingArchivesPath, path)
			if err != nil {
				return err
			}

			if ok {
				dir := strings.Replace(path, archivesPath+"/", "", 1)
				isExpired, err := checkExpired(dir, expired)
				if err != nil {
					return err
				}

				if isExpired {
					err := os.RemoveAll(path)
					if err != nil {
						return err
					}
					log.Println(path)
					return filepath.SkipDir
				}
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}
