package tests

import (
	"github.com/Wooderan/mcdata"
	"reflect"
	"testing"
)

var (
	expectedPlatforms = []string{"bedrock", "pc"}
)

func TestAvailablePlatforms(t *testing.T) {
	platforms, err := mcdata.GetAvailablePlatforms(MCDataPath)
	if err != nil {
		t.Errorf("GetAvailablePlatforms() error = %v", err)
		return
	}
	if !reflect.DeepEqual(platforms, expectedPlatforms) {
		t.Errorf("GetAvailablePlatforms() = %v, want %v", platforms, expectedPlatforms)
	}
}
