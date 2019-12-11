package torsaver

import (
	"crypto/sha256"
	"fmt"
	cache "github.com/gocacher/badger-cache/v2"
	"github.com/gocacher/cacher"
	"github.com/goextension/log"
	"io/ioutil"
)

var useCache bool = false

func init() {
	RegisterCache()
}

func RegisterCache() {
	cacher.Register(cache.New())
	useCache = true
}

func Get(url string) (data []byte, e error) {
	log.Infow("get", "url", url)
	name := Hash(url)
	if useCache {
		data, e = cacher.Get(name)
		if e == nil {
			return
		}
	}
	get, e := cli.Get(url)
	if e != nil {
		return nil, Wrap(e, "httget")
	}
	data, e = ioutil.ReadAll(get.Body)
	if e != nil {
		return nil, Wrap(e, "readall")
	}
	e = cacher.Set(name, data)
	if e != nil {
		return nil, Wrap(e, "cache")
	}
	return
}

// Hash ...
func Hash(url string) string {
	sum256 := sha256.Sum256([]byte(url))
	return fmt.Sprintf("%x", sum256)
}
