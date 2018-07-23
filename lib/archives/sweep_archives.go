package archives

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SweepArchives sweep Archives
func SweepArchives() error {

	archivesPath := archivesPath()
	matchingArchivesPath := filepath.Join(archivesPath, "*")

	now := time.Now()
	expired := now.AddDate(0, -1, 0)

	err := filepath.Walk(archivesPath, func(path string, info os.FileInfo, err error) error {
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

	return err

}
