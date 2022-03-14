package timer

import (
	"testing"
)

func Test_timer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timer()
		})
	}
}
