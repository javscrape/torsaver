package torsaver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/goextension/log"
	"github.com/zyxar/argo/rpc"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// DefaultNYAAURL ...
var DefaultNYAAURL = `https://sukebei.nyaa.si`

// DefaultSavePath ...
var DefaultSavePath = "tmp"

// NyaaOption ...
type NyaaOption func(nyaa *Nyaa)

// NyaaTorrent ...
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

// Nyaa ...
type Nyaa struct {
	torrents []*NyaaTorrent
	limit    int
	path     string
	Aria     rpc.Client
	User     string //user:NewDragon,offkab
	F        string
	Category string
	Query    string
	Sort     string
	Order    string
	Page     int
}

func (n Nyaa) CurrentPage() int {
	return n.Page
}

func (n *Nyaa) SetPage(p int) {
	n.Page = p
}

// Download ...
func (n Nyaa) Download(idx int) (e error) {
	size := len(n.torrents)
	if idx >= size {
		return fmt.Errorf("index(%d) is over than size(%d)", idx, size)
	}

	t := n.torrents[idx]
	torrent, e := n.Aria.AddTorrent(filepath.Join(n.path, t.ID+".torrent"))
	if e != nil {
		return Wrap(e, "download add")
	}
	log.Infow("download", "name", t.Name, "tid", torrent)
	return nil
}

// DownloadAll ...
func (n Nyaa) DownloadAll() (e error) {
	for idx := range n.torrents {
		e = n.Save(idx)
		if e != nil {
			return Wrap(e, "download save")
		}
		e = n.Download(idx)
		if e != nil {
			return Wrap(e, "download down")
		}
	}
	return nil
}

// SaveAll ...
func (n Nyaa) SaveAll() (e error) {
	for i := range n.torrents {
		if i >= n.limit {
			return nil
		}
		e := n.Save(i)
		if e != nil {
			return e
		}
	}
	return nil
}

// Save ...
func (n Nyaa) Save(idx int) (e error) {
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
	e = ioutil.WriteFile(filepath.Join(n.path, t.ID+".torrent"), all, 0755)
	if e != nil {
		return Wrap(e, "write file")
	}
	marshal, e := json.MarshalIndent(t, "", " ")
	if e != nil {
		return Wrap(e, "save")
	}

	return ioutil.WriteFile(filepath.Join(n.path, t.ID+".json"), marshal, 0755)
}

// List ...
func (n Nyaa) List() (l []string) {
	for _, t := range n.torrents {
		l = append(l, t.Name)
	}
	return
}

// Limit ...
func (n *Nyaa) Limit(i int) {
	n.limit = i
}

// Find ...
func (n *Nyaa) Find(name string) error {
	n.Query = name
	for len(n.torrents) < n.limit {
		get, err := Get(n.url())
		if err != nil {
			return err
		}
		reader, err := goquery.NewDocumentFromReader(bytes.NewReader(get))
		if err != nil {
			return err
		}
		n.torrents = append(n.torrents, findSub(reader, n.limit-len(n.torrents))...)
		n.SetPage(n.Page + 1)
	}
	return nil
}

func findSub(document *goquery.Document, limit int) []*NyaaTorrent {
	var result []*NyaaTorrent
	document.Find("table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		if i >= limit {
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

func (n Nyaa) url() string {
	url := DefaultNYAAURL
	if n.User != "" {
		url = strings.Join([]string{DefaultNYAAURL, "user", n.User}, "/")
	}

	args := fmt.Sprintf("f=0&c=%v&q=%v&s=%v&o=%v&p=%v", n.Category, n.Query, n.Sort, n.Order, n.Page)

	return strings.Join([]string{url, args}, "?")
}

// NewNyaa ...
func NewNyaa(opts ...NyaaOption) Saver {
	n := &Nyaa{
		torrents: nil,
		limit:    50,
		path:     DefaultSavePath,
		Aria:     nil,
		User:     "",
		F:        "",
		Category: "2_2", //video(2_2) args
		Query:    "",
		Sort:     "seeders",
		Order:    "desc",
		Page:     1,
	}
	for _, opt := range opts {
		opt(n)
	}
	return n
}
