package bump_version

import "testing"

func TestChangeVersion(t *testing.T) {
	testCases := []struct {
		in    string
		vtype VersionType
		out   string
	}{
		{"0.4", Major, "v1.0"},
		{"0.4.0", Major, "v1.0.0"},
		{"1.0", Major, "v2.0"},
		{"1", Major, "v2"},
		{"1.0.1", Minor, "v1.1.0"},

		{"v0.4", Major, "v1.0"},
		{"v0.4.0", Major, "v1.0.0"},
		{"v1.0", Major, "v2.0"},
		{"v1", Major, "v2"},
		{"v1.0.1", Minor, "v1.1.0"},
	}
	for _, tt := range testCases {
		v, err := changeVersion(tt.vtype, tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if v.String() != tt.out {
			t.Errorf("changeVersion(%s, %s): got %s, want %s", tt.vtype, tt.in, v.String(), tt.out)
		}
	}
}
