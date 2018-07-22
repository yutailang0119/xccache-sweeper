package deriveddata

import (
	"log"
	"os"
)

// SweepDerivedData sweep DerivedData
func SweepDerivedData() error {

	derivedDataPaths, err := derivedDataPaths()
	if err != nil {
		return err
	}

	for _, path := range derivedDataPaths {
		err := os.RemoveAll(path)
		if err == nil {
			log.Println(path)
		}
	}

	return nil
}
