package parse

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
)

var cacheIndexName string = ""
var cacheIndex string = ""

// GetIndexContent retrieves the final index.html content to be returned to the user.
func GetIndexContent(path string) (string, error) {
	meta, err := GetMeta(path)
	if err != nil {
		if errors.Is(err, errMeta404) {
			return GetIndexContent("/")
		} else if errors.Is(err, errMetaFatal404) {
			log.Fatal("Cannot fetch any meta data")
		}
	}
	newHTML := strings.Replace(cacheIndex, "<!-- SSR-HEAD -->", meta, 1)
	return newHTML, nil
}

// CacheIndex caches the initial index.[hash].html every 5 minutes
func CacheIndex() {
	match, err := filepath.Glob("index.*.html")
	if err != nil {
		log.Fatal(err)
	}
	if len(match) < 1 {
		log.Fatal("No template index.[hash].html found")
	}
	if match[0] != cacheIndexName {
		html, err := ioutil.ReadFile(match[0])
		if err != nil {
			log.Fatal(err)
		}
		cacheIndexName = match[0]
		cacheIndex = string(html)
	}
	time.Sleep(300000 * time.Millisecond)
	CacheIndex()
}
