package run

import "testing"

func TestRun_Runner(t *testing.T) {
	tests := []struct {
		name string
		r    *Run
	}{
		{
			name: "none",
			r:    &Run{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Run{}
			r.Runner()
		})
	}
}
