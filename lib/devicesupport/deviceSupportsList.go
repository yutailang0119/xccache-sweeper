package devicesupport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type deviceSupportsList struct {
	iOS     deviceSupports
	watchOS deviceSupports
	tvOS    deviceSupports
}

func (v deviceSupportsList) osList(platform string) ([]deviceSupports, error) {
	switch platform {
	case "iOS":
		return []deviceSupports{v.iOS}, nil
	case "watchOS":
		return []deviceSupports{v.watchOS}, nil
	case "tvOS":
		return []deviceSupports{v.tvOS}, nil
	case "all":
		var list []deviceSupports
		iOSVersions, err := v.iOS.versions()
		if err != nil {
			return nil, err
		}
		if len(iOSVersions) != 0 {
			list = append(list, v.iOS)
		}

		watchOSVersions, err := v.watchOS.versions()
		if err != nil {
			return nil, err
		}
		if len(watchOSVersions) != 0 {
			list = append(list, v.watchOS)
		}

		tvOSVersion, err := v.tvOS.versions()
		if err != nil {
			return nil, err
		}
		if len(tvOSVersion) != 0 {
			list = append(list, v.tvOS)
		}

		return list, nil
	default:
		return nil, fmt.Errorf("not found a platform: %s", platform)
	}
}

func (v deviceSupportsList) allDeviceSupports(platform string) ([]deviceSupport, error) {
	osList, err := v.osList(platform)
	if err != nil {
		return nil, err
	}
	var all []deviceSupport
	for _, supports := range osList {
		versions, err := supports.versions()
		if err != nil {
			return nil, err
		}
		all = append(all, versions...)
	}
	return all, nil
}

func (v deviceSupportsList) deleteAll(platform string) error {
	osList, err := v.osList(platform)
	if err != nil {
		return err
	}
	if osList == nil {
		fmt.Println("not found a platform: ", platform)
		return nil
	}
	allDeviceSupports, err := v.allDeviceSupports(platform)
	if err != nil {
		return err
	}
	if len(allDeviceSupports) == 0 {
		fmt.Println("not found DeviceSupport: ", platform)
		return nil
	}
	message := "Deleted: "
	for _, supports := range osList {
		versions, err := supports.versions()
		if err != nil {
			return err
		}
		message += "\n - " + supports.osName
		for _, support := range versions {
			message += "\n    - " + support.version
			err = os.RemoveAll(support.path)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println(message)

	return nil
}

func (v deviceSupportsList) askDelete(platform string) error {
	osList, err := v.osList(platform)
	if err != nil {
		return err
	}

	allDeviceSupports, err := v.allDeviceSupports(platform)
	if err != nil {
		return err
	}
	if len(allDeviceSupports) == 0 {
		fmt.Println("not found DeviceSupport: ", platform)
		return nil
	}

	message := "Which do you want to delete? Please select index"
	count := 0
	for _, supports := range osList {
		versions, err := supports.versions()
		if err != nil {
			return err
		}
		message += "\n" + supports.osName
		for _, version := range versions {
			message += "\n " + strconv.Itoa(count) + ": " + version.description()
			count++
		}
	}
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	inputIndex, _ := reader.ReadString('\n')
	inputIndex = strings.TrimSuffix(inputIndex, "\n")
	index, err := strconv.Atoi(inputIndex)
	if err != nil {
		fmt.Println("Please input number")
		return v.askDelete(platform)
	}

	target := allDeviceSupports[index]
	err = os.RemoveAll(target.path)
	if err != nil {
		return err
	}

	fmt.Println("Deleted: ", target.description())

	allDeviceSupports, _ = v.allDeviceSupports(platform)
	if len(allDeviceSupports) == 0 {
		fmt.Println("Complete all deleted 🎉")
		return nil
	}

	return v.askDelete(platform)
}
