package distroyer

import (
	"os/exec"
	"strings"
)

func Arch() (string, error) {
	out, err := exec.Command("bash", "-c", "uname -m").CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
