package devicesupport

import (
	"io/ioutil"
	"strings"
)

type deviceSupports struct {
	osName string
}

func (v deviceSupports) dirName() string {
	return v.osName + " DeviceSupport"
}

func (v deviceSupports) path() string {
	return xcodeFilesPath() + "/" + v.dirName()
}

func (v deviceSupports) versions() ([]deviceSupport, error) {
	files, err := ioutil.ReadDir(v.path())
	if err != nil {
		return nil, err
	}

	var list []deviceSupport
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		path := v.path() + "/" + file.Name()
		list = append(list, deviceSupport{v.osName, file.Name(), path})
	}

	return list, nil
}
