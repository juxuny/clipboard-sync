package lib

import (
	"github.com/hashicorp/go-version"
)

// appVersion是否大于等于minVersion
func VersionGreatThanOrEqual(appVersion, minVersion string) bool {
	v1, err := version.NewVersion(minVersion)
	kit := false
	if err == nil {
		v2, err := version.NewVersion(appVersion)
		if err == nil {
			if v2.Equal(v1) || v2.GreaterThan(v1) { //等于或者大于
				kit = true
			}
		}
	}
	return kit
}
