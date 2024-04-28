package mcdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Wooderan/mcdata/generated/types"
)

func GetAvailableVersions(MCDataPath string, platform string) ([]string, error) {
	isAvailable := IsAvailablePlatform(MCDataPath, platform)
	if !isAvailable {
		return nil, errors.New(fmt.Sprintf("there is no such platform like %s\n", platform))
	}

	versionsJsonPath := filepath.Join(MCDataPath, "data", platform, "common", "versions.json")
	versionsBytes, err := os.ReadFile(versionsJsonPath)
	if err != nil {
		return nil, err
	}

	var versions []string
	if err = json.Unmarshal(versionsBytes, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

func GetAvailableProtocolVersions(MCDataPath string, platform string) ([]types.Version, error) {
	isAvailable := IsAvailablePlatform(MCDataPath, platform)
	if !isAvailable {
		return nil, errors.New(fmt.Sprintf("there is no such platform like %s\n", platform))
	}

	versionsJsonPath := filepath.Join(MCDataPath, "data", platform, "common", "protocolVersions.json")
	versionsBytes, err := os.ReadFile(versionsJsonPath)
	if err != nil {
		return nil, err
	}

	var versions []types.Version
	if err = json.Unmarshal(versionsBytes, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}
