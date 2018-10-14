package utils

import (
	"testing"
)

func TestGetRegionByIp(t *testing.T) {
	region := GetRegionFromCache("61.135.169.121")
	region1 := GetRegionFromCache("61.135.169.120")
	t.Logf("%v, %v", region.Province, region.City)
	t.Logf("%v, %v", region1.Province, region1.City)
}