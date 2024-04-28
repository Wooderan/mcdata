package tests

import (
	"github.com/Wooderan/mcdata"
	"reflect"
	"testing"
)

var (
	platformVersionsMap = map[string][]string{
		"bedrock": []string{
			"0.14", "0.15", "1.0", "1.16.201", "1.16.210", "1.16.220", "1.17.0", "1.17.10",
			"1.17.30", "1.17.40", "1.18.0", "1.18.11", "1.18.30", "1.19.1", "1.19.10", "1.19.20",
			"1.19.21", "1.19.30", "1.19.40", "1.19.50", "1.19.60", "1.19.62", "1.19.63", "1.19.70",
			"1.19.80", "1.20.0", "1.20.10", "1.20.15", "1.20.30", "1.20.40", "1.20.50", "1.20.61",
			"1.20.71", "1.20.80",
		},
		"pc": []string{
			"0.30c", "1.7", "1.8", "15w40b", "1.9", "1.9.1-pre2", "1.9.2", "1.9.4", "16w20a", "1.10-pre1",
			"1.10", "1.10.1", "1.10.2", "16w35a", "1.11", "1.11.2", "17w15a", "17w18b", "1.12-pre4", "1.12",
			"1.12.1", "1.12.2", "17w50a", "1.13", "1.13.1", "1.13.2-pre1", "1.13.2-pre2", "1.13.2", "1.14",
			"1.14.1", "1.14.3", "1.14.4", "1.15", "1.15.1", "1.15.2", "20w13b", "20w14a", "1.16-rc1", "1.16",
			"1.16.1", "1.16.2", "1.16.3", "1.16.4", "21w07a", "1.17", "1.17.1", "1.18", "1.18.1", "1.18.2",
			"1.19", "1.19.2", "1.19.3", "1.19.4", "1.20", "1.20.1", "1.20.2", "1.20.3", "1.20.4",
		},
	}
)

func TestAvailableVersions(t *testing.T) {
	for _, platform := range expectedPlatforms {
		versions, err := mcdata.GetAvailableVersions(platform)
		if err != nil {
			t.Errorf("GetAvailableVersions() error = %v", err)
			return
		}

		if !reflect.DeepEqual(versions, platformVersionsMap[platform]) {
			t.Errorf("GetAvailableVersions() = %v, want %v", versions, platformVersionsMap[platform])
		}
	}
}

func TestAvailableProtocolVersions(t *testing.T) {
	for _, platform := range expectedPlatforms {
		protocolVersions, err := mcdata.GetAvailableProtocolVersions(platform)
		if err != nil {
			t.Errorf("GetAvailableProtocolVersions() error = %v", err)
			return
		}

		if protocolVersions == nil || len(protocolVersions) == 0 {
			t.Errorf("GetAvailableProtocolVersions().Length - not enough length")
		}
	}
}
