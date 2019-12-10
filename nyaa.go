package torsaver

import "strings"

var DefaultNYAAURL = `https://sukebei.nyaa.si`
var DefaultNYAAUser = `offkab`

type nyaa struct {
	User  string
	Limit int64
	F     string
	C     string
	Q     string
	S     string
	O     string
	P     string
}

func NewNyaa() Saver {

}

func (n nyaa) URL() string {
	url := strings.Join([]string{DefaultNYAAURL, "user", DefaultNYAAUser}, "/")
	return strings.Join([]string{url, `f=0&c=0_0&q=&s=id&o=desc&p=2`}, "?")
	//todo:
	//args := strings.Join([]string{}, "&")
}

type NyaaTorrent struct {
	Category  string
	Name      string
	File      string
	Magnet    string
	Size      string
	Data      string
	Seeders   string
	Leechers  string
	Downloads string
}
