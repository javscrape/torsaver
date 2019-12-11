package torsaver

import (
	"context"
	"github.com/zyxar/argo/rpc"
	"testing"
)

func init() {
	err := RegisterProxy("socks5://127.0.0.1:10808")
	if err != nil {
		panic(err)
	}
}

// TestNewNyaa ...
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

// Test_nyaa_Save ...
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
			if err := n.Save(tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test_nyaa_SaveAll ...
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
			n := NewNyaa(func(nyaa *Nyaa) {
				var e error
				nyaa.User = ""
				DefaultAriaRPC = "http://aria2rpc.y11e.com:5000/jsonrpc"
				nyaa.Aria, e = NewRPCClient(context.Background())
				if e != nil {
					panic(e)
				}
			})
			err := n.Find("")
			if err != nil {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := n.SaveAll(); (err != nil) != tt.wantErr {
				t.Errorf("SaveAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestNyaa_DownloadAll ...
func TestNyaa_DownloadAll(t *testing.T) {
	type fields struct {
		torrents []*NyaaTorrent
		limit    int64
		path     string
		Aria     rpc.Client
		Name     string
		User     string
		F        string
		C        string
		Q        string
		S        string
		O        string
		P        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNyaa(func(nyaa *Nyaa) {
				var e error
				nyaa.User = "offkab"
				DefaultAriaRPC = "http://aria2rpc.y11e.com:5000/jsonrpc"
				nyaa.Aria, e = NewRPCClient(context.Background())
				if e != nil {
					panic(e)
				}
			})
			err := n.Find("FHD")
			if err != nil {
				t.Errorf("DownloadAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := n.DownloadAll(); (err != nil) != tt.wantErr {
				t.Errorf("DownloadAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
