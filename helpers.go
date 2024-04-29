package mcdata

import (
	"strconv"
	"strings"
)

func compareVersions(version1, version2 string) int {
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		v1, _ := strconv.Atoi(parts1[i])
		v2, _ := strconv.Atoi(parts2[i])

		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	return 0
}
