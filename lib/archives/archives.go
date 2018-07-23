package archives

import (
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func archivesPath() string {

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
