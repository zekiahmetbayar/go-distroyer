package distroyer

import (
	"syscall"

	"golang.org/x/sys/windows/registry"
)

func Codename() (string, error) {
	key, err := registry.OpenKey(syscall.HKEY_LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE|registry.SET_VALUE|registry.WRITE)
	if err != nil {
		return "", nil
	}

	name, _, err := key.GetStringValue("ProductName")
	if err != nil {
		return "", nil
	}

	return name, nil
}
