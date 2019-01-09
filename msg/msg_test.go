package msg

import "testing"

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "echo文字列を返却すること", want: "Go We are the world."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
