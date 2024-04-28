package mcdata

import (
	"path/filepath"
)

func GetAvailablePlatforms() ([]string, error) {
	dataPath := filepath.Join(SubmodulePath, "data")

	dataSubItems, err := embeddedMinecraftData.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	var availablePlatforms []string
	for _, item := range dataSubItems {
		if item.IsDir() {
			availablePlatforms = append(availablePlatforms, item.Name())
		}
	}

	return availablePlatforms, nil
}

func IsAvailablePlatform(platform string) bool {
	availablePlatforms, err := GetAvailablePlatforms()
	if err != nil {
		return false
	}

	isAvailable := false
	for _, availableProtocolVersion := range availablePlatforms {
		if availableProtocolVersion == platform {
			isAvailable = true
			break
		}
	}

	return isAvailable
}
