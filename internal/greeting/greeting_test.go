package greeting

import "testing"

func TestMessage(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "default", in: "", want: "Hello, world!"},
		{name: "named", in: "gopher", want: "Hello, gopher!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.in); got != tt.want {
				t.Errorf("Message(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
