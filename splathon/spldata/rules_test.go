package spldata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetRuleByKey(t *testing.T) {
	tests := []struct {
		in   string
		want *Rule
	}{
		{in: "turf_war", want: &turfWar},
		{in: "splat_zones", want: &splatZones},
		{in: "tower_control", want: &towerControl},
		{in: "rainmaker", want: &rainmaker},
		{in: "clam_blitz", want: &clamBlitz},
	}
	for _, tt := range tests {
		got, ok := GetRuleByKey(tt.in)
		if !ok {
			t.Errorf("%q not found", tt.in)
			continue
		}
		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Errorf("result has diff:\n%s", diff)
		}
	}
}
