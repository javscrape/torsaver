package torsaver

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var DefaultNYAAURL = `https://sukebei.nyaa.si`
var DefaultNYAAUser = `offkab`

type nyaa struct {
	Name string
	User string
	F    string
	C    string
	Q    string
	S    string
	O    string
	P    string
}

func (n nyaa) Limit(i int64) {
	panic("implement me")
}

func (n nyaa) Find(name string) error {
	get, err := Get(n.URL(name))
	if err != nil {
		return err
	}
	reader, err := goquery.NewDocumentFromReader(bytes.NewReader(get))
	if err != nil {
		return err
	}
	findSub(reader)

	return nil
}

func findSub(document *goquery.Document) {
	document.Find("table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("idx", i, "text", selection.Text())
	})
}

func NewNyaa() Saver {
	return &nyaa{
		User: "",
		F:    "",
		C:    "",
		Q:    "",
		S:    "",
		O:    "",
		P:    "",
	}
}

func (n nyaa) URL(name string) string {
	url := strings.Join([]string{DefaultNYAAURL, "user", DefaultNYAAUser}, "/")
	args := fmt.Sprintf("f=0&c=0_0&q=%s&s=id&o=desc&p=1", name)

	return strings.Join([]string{url, args}, "?")
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
