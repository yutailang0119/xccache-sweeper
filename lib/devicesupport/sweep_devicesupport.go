package devicesupport

import (
	"os/user"
	"strings"
)

func xcodeFilesPath() string {
	usr, _ := user.Current()
	return strings.Replace("~/Library/Developer/Xcode", "~", usr.HomeDir, 1)
}

// SweepDeviceSupports sweep ~/Library/Developer/Xcode/*DeviceSupport
// platform: "iOS", "watchOS", "tvOS" and "all"
// isAll: force delete all
func SweepDeviceSupports(platform string, isAll bool) error {

	if platform == "" {
		platform = "all"
	}

	iOS := deviceSupports{"iOS"}
	watchOS := deviceSupports{"watchOS"}
	tvOS := deviceSupports{"tvOS"}
	list := deviceSupportsList{iOS: iOS, watchOS: watchOS, tvOS: tvOS}

	if isAll {
		err := list.deleteAll(platform)
		if err != nil {
			return err
		}

		return nil
	}

	err := list.askDelete(platform)
	if err != nil {
		return err
	}

	return nil

}
