package tests

import (
	"github.com/Wooderan/mcdata"
	"reflect"
	"testing"
)

var (
	versionFeaturesMap = map[string]map[string][]string{
		"bedrock": map[string][]string{
			"1.20.80": []string{
				"usesPalettedChunks", "newRecipeSchema", "usesBlockStates",
				"tallWorld", "spawnEggsUseInternalIdInNbt", "spawnEggsHaveSpawnedEntityInName",
				"explicitMaxDurability", "compressorInPacketHeader", "blockHashes",
			},
		},
	}
)

func TestGetAvailableFeatures(t *testing.T) {
	for platform, versions := range versionFeaturesMap {
		for version, expectedFeatures := range versions {
			features, err := mcdata.GetAvailableFeatures(MCDataPath, platform, version)
			if err != nil {
				t.Errorf("GetAvailableFeatures() error = %v", err)
			}

			if !reflect.DeepEqual(features, expectedFeatures) {
				t.Errorf("GetAvailableFeatures() = %v, want %v", features, expectedFeatures)
			}
		}
	}
}
