package release

import (
	"strconv"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	GetNativeSystemInfo = kernel32.NewProc("GetNativeSystemInfo")
)

type SYSTEM_INFO struct {
	wProcessorArchitecture uint16
}

func Arch() (string, error) {
	var sys SYSTEM_INFO
	GetNativeSystemInfo.Call(
		uintptr(unsafe.Pointer(&sys)),
	)

	processArch := strconv.Itoa(int(sys.wProcessorArchitecture))
	switch processArch {
	case "9":
		return "x64", nil
	case "5":
		return "arm", nil
	case "12":
		return "arm64", nil
	case "6":
		return "ia64", nil
	case "0":
		return "x86", nil
	default:
		return "unknown arch", nil
	}
}
