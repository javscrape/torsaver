package torsaver

import (
	"testing"
)

func init() {
	err := RegisterProxy("socks5://127.0.0.1:10808")
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
			name: "newnyaa",
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

func Test_nyaa_Save(t *testing.T) {
	type fields struct {
		torrents []*NyaaTorrent
		limit    int64
		Name     string
		User     string
		F        string
		C        string
		Q        string
		S        string
		O        string
		P        string
	}
	type args struct {
		idx  int
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "save",
			fields: fields{},
			args: args{
				idx:  0,
				path: "d:\\torrent\\temp\\",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNyaa()
			err := n.Find("FC2")
			if err != nil {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := n.Save(tt.args.idx, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nyaa_SaveAll(t *testing.T) {
	type fields struct {
		torrents []*NyaaTorrent
		limit    int64
		Name     string
		User     string
		F        string
		C        string
		Q        string
		S        string
		O        string
		P        string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "saveall",
			fields: fields{},
			args: args{
				path: "D:\\torrent\\temp",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNyaa()
			err := n.Find("FC2")
			if err != nil {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := n.SaveAll(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("SaveAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
