package torsaver

import (
	"testing"
)

func init() {
	err := RegisterProxy("socks5://127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
}

func TestNewNyaa(t *testing.T) {
	tests := []struct {
		name string
		want Saver
	}{
		{
			name: "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNyaa()
			if got == nil {
				t.Errorf("NewNyaa() = %v, want %v", got, tt.want)
				return
			}
			err := got.Find("FC2")
			if err != nil {
				t.Errorf("NewNyaa() = %v, want %v", err, nil)
			}
		})
	}
}
