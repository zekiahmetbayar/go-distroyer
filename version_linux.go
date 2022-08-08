package release

import (
	"os/exec"
	"regexp"
)

func Version() (string, error) {
	var re = regexp.MustCompile(`(?m)^VERSION_ID="*(.*?)"*$`)
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
