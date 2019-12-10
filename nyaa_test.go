package torsaver

import (
	"reflect"
	"testing"
)

func TestNewNyaa(t *testing.T) {
	tests := []struct {
		name string
		want Saver
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNyaa()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNyaa() = %v, want %v", got, tt.want)
			}
			err := got.Find("FC2")
			if err != nil {
				t.Errorf("NewNyaa() = %v, want %v", err, nil)
			}
		})
	}
}
