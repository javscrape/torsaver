package torsaver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var DefaultNYAAURL = `https://sukebei.nyaa.si`
var DefaultNYAAUser = `offkab`

type nyaa struct {
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

func (n nyaa) SaveAll(path string) (e error) {
	for i := range n.torrents {
		if int64(i) >= n.limit {
			return nil
		}
		e := n.Save(i, path)
		if e != nil {
			return e
		}
	}
	return nil
}

func (n nyaa) Save(idx int, path string) (e error) {
	size := len(n.torrents)
	if idx >= size {
		return fmt.Errorf("index(%d) is over than size(%d)", idx, size)
	}

	t := n.torrents[idx]
	get, e := cli.Get(DefaultNYAAURL + t.File)
	if e != nil {
		return Wrap(e, "cli get")
	}
	all, e := ioutil.ReadAll(get.Body)
	if e != nil {
		return Wrap(e, "read all")
	}
	e = ioutil.WriteFile(filepath.Join(path, t.ID+".torrent"), all, 0755)
	if e != nil {
		return Wrap(e, "write file")
	}
	marshal, e := json.MarshalIndent(t, "", " ")
	if e != nil {
		return Wrap(e, "save")
	}

	return ioutil.WriteFile(filepath.Join(path, t.ID+".json"), marshal, 0755)
}

func (n nyaa) List() (l []string) {
	for _, t := range n.torrents {
		l = append(l, t.Name)
	}
	return
}

func (n *nyaa) Limit(i int64) {
	n.limit = i
}

func (n *nyaa) Find(name string) error {
	n.Name = name
	get, err := Get(n.URL())
	if err != nil {
		return err
	}
	reader, err := goquery.NewDocumentFromReader(bytes.NewReader(get))
	if err != nil {
		return err
	}
	n.torrents = findSub(reader, n.limit)

	return nil
}

func findSub(document *goquery.Document, limit int64) []*NyaaTorrent {
	var result []*NyaaTorrent
	document.Find("table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		if int64(i) >= limit {
			return
		}
		result = append(result, decodeNyaa(selection))
	})

	return result
}

func decodeNyaa(sel *goquery.Selection) *NyaaTorrent {
	var tor NyaaTorrent
	sel.Find("td").Each(func(i int, selection *goquery.Selection) {
		switch i {
		case 0:
			tor.Category = selection.Find("a").AttrOr("title", "")
		case 1:
			tor.Name = selection.Find("a").AttrOr("title", "")
			tor.Link = selection.Find("a").AttrOr("href", "")
			l := strings.Split(tor.Link, "/")
			for _, l1 := range l {
				tor.ID = l1
			}

		case 2:
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tor.File = selection.AttrOr("href", "")
				} else {
					tor.Magnet = selection.AttrOr("href", "")
				}
			})
		case 3:
			tor.Size = selection.Text()
		case 4:
			tor.Date = selection.Text()
		case 5:
			tor.Seeders = selection.Text()
		case 6:
			tor.Leechers = selection.Text()
		case 7:
			tor.Downloads = selection.Text()
		default:
			fmt.Println("idx", i, "text", selection.Text())
		}
	})
	return &tor
}

func NewNyaa() Saver {
	return &nyaa{
		torrents: nil,
		limit:    50,
		Name:     "",
		User:     "",
		F:        "",
		C:        "",
		Q:        "",
		S:        "",
		O:        "",
		P:        "1",
	}
}

func (n nyaa) URL() string {
	url := strings.Join([]string{DefaultNYAAURL, "user", DefaultNYAAUser}, "/")
	args := fmt.Sprintf("f=0&c=0_0&q=%s&s=id&o=desc&p=%s", n.Name, n.P)

	return strings.Join([]string{url, args}, "?")
	//todo:
	//args := strings.Join([]string{}, "&")
}

type NyaaTorrent struct {
	ID        string
	Link      string
	Category  string
	Name      string
	File      string
	Magnet    string
	Size      string
	Date      string
	Seeders   string
	Leechers  string
	Downloads string
}
