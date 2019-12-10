package torsaver

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name: "get",
			args: args{
				url: "https://sukebei.nyaa.si/user/offkab?f=0&c=0_0&q=&s=id&o=desc",
			},
			wantData: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Get() gotData = %v, want %v", string(gotData), tt.wantData)
			}
		})
	}
}
