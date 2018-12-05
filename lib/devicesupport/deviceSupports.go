package devicesupport

import (
	"io/ioutil"
	"os"
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
	path := v.path()

	var list []deviceSupport

	exist, err := exists(path)

	if !exist {
		return list, nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		path := v.path() + "/" + file.Name()
		list = append(list, deviceSupport{v.osName, file.Name(), path})
	}

	return list, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
