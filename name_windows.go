package distroyer

func PrettyName() (string, error) {
	name, err := Codename()
	if err != nil {
		return "", err
	}

	return name, nil
}
