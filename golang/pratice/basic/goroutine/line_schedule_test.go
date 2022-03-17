package goroutine

import "testing"

func Test_line_schedule(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "line_schedule"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line_schedule()
		})
	}
}
