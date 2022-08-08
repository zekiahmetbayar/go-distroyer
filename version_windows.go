package release

import (
	"syscall"

	"golang.org/x/sys/windows/registry"
)

func Version() (string, error) {
	key, err := registry.OpenKey(syscall.HKEY_LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE|registry.SET_VALUE|registry.WRITE)
	if err != nil {
		return "", nil
	}

	version, _, err := key.GetStringValue("CurrentVersion")
	if err != nil {
		return "", nil
	}

	return version, nil
}
