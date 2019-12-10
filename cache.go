package torsaver

import (
	"crypto/sha256"
	"fmt"
	cache "github.com/gocacher/badger-cache/v2"
	"github.com/gocacher/cacher"
	"io/ioutil"
	"net/http"
)

var useCache bool = false

func RegisterCache() {
	cacher.Register(cache.New())
	useCache = true
}

func Get(url string) (data []byte, e error) {
	name := Hash(url)
	get, e := http.Get(url)
	if e != nil {
		return nil, Wrap(e, "httget")
	}
	bys, e := ioutil.ReadAll(get.Body)
	if e != nil {
		return nil, Wrap(e, "readall")
	}
	e = cacher.Set(name, bys)
	if e != nil {
		return nil, Wrap(e, "cache")
	}
	return bys, nil
}

// Hash ...
func Hash(url string) string {
	sum256 := sha256.Sum256([]byte(url))
	return fmt.Sprintf("%x", sum256)
}
