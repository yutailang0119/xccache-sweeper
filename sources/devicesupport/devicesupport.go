package devicesupport

type deviceSupport struct {
	osName string
	version string
	path    string
}

func (v deviceSupport) description() string {
	return v.osName + "/" + v.version
}
