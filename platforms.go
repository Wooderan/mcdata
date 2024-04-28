package mcdata

import (
	"os"
	"path/filepath"
)

func GetAvailablePlatforms(MCDataPath string) ([]string, error) {
	dataPath := filepath.Join(MCDataPath, "data")

	dataSubItems, err := os.ReadDir(dataPath)
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

func IsAvailablePlatform(MCDataPath string, platform string) bool {
	availablePlatforms, err := GetAvailablePlatforms(MCDataPath)
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
