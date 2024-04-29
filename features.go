package mcdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Wooderan/mcdata/generated/types"
	"log"
	"os"
	"path/filepath"
)

func GetAllFeatures(MCDataPath string, platform string) (types.Features, error) {
	isAvailable := IsAvailablePlatform(MCDataPath, platform)
	if !isAvailable {
		return nil, errors.New(fmt.Sprintf("there is no such platform like %s\n", platform))
	}

	featuresJsonPath := filepath.Join(MCDataPath, "data", platform, "common", "features.json")
	featuresBytes, err := os.ReadFile(featuresJsonPath)
	if err != nil {
		return nil, err
	}

	var features types.Features
	if err = json.Unmarshal(featuresBytes, &features); err != nil {
		return nil, err
	}
	return features, nil
}

func GetAvailableFeatures(MCDataPath string, platform string, version string) ([]string, error) {
	allFeatures, err := GetAllFeatures(MCDataPath, platform)
	if err != nil {
		return nil, err
	}

	availableVersions, err := GetAvailableVersions(MCDataPath, platform)
	if err != nil {
		return nil, err
	}
	latestVersion := availableVersions[len(availableVersions)-1]

	var availableFeatures []string
	for _, feature := range allFeatures {
		if feature.Versions != nil {
			if len(feature.Versions) == 1 {
				if compareVersions(version, feature.Versions[0]) == 0 {
					availableFeatures = append(availableFeatures, *feature.Name)
				}
			} else if len(feature.Versions) == 2 {
				var fromVersion string
				var toVersion string
				fromVersion = feature.Versions[0]
				if feature.Versions[1] == "latest" {
					toVersion = latestVersion
				} else {
					toVersion = feature.Versions[1]
				}
				if compareVersions(version, fromVersion) >= 0 && compareVersions(version, toVersion) <= 0 {
					availableFeatures = append(availableFeatures, *feature.Name)
				}
			} else {
				log.Panicln("Unsupported set of versions in features")
			}
		}
	}

	return availableFeatures, err
}
