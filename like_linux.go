package distroyer

import (
	"os/exec"
	"regexp"
)

func Like() (string, error) {
	var re = regexp.MustCompile(`(?m)^ID_LIKE="*(.*?)"*$`)
	out, err := exec.Command("bash", "-c", "cat /etc/os-release").CombinedOutput()
	if err != nil {
		return "", err
	}
	if len(re.FindStringSubmatchIndex(string(out))) > 0 {
		distro := re.FindStringSubmatch(string(out))
		return distro[1], nil
	}
	return "", nil
}
