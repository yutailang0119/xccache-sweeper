package deriveddata

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func derivedDataPaths() ([]string, error) {
	xcodeBuildLocationStyle, err := exec.Command("defaults", "read", "com.apple.dt.Xcode", "IDEBuildLocationStyle").Output()
	if err != nil {
		xcodeBuildLocationStyle = []byte("Unique")
	}

	usr, _ := user.Current()
	if strings.TrimSpace(string(xcodeBuildLocationStyle)) == "Unique" {
		return []string{strings.Replace("~/Library/Developer/Xcode/DerivedData", "~", usr.HomeDir, 1)}, nil
	}

	paths := []string{}
	err = filepath.Walk(usr.HomeDir, func(path string, info os.FileInfo, err error) error {
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
