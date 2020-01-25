package parse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var errMeta404 = errors.New("Meta data not found")
var errMetaFatal404 = errors.New("None meta data found")

// GetMeta returns metatags based on given path.
//
// Path is expected to match the directory and file structure. For example:
// Path: /          = pages/meta
// Path: /page-one  = pages/page-one/meta
// Path: /blog/post = pages/blog/post/meta
func GetMeta(path string) (string, error) {
	var builder strings.Builder
	shortPath := path[1:]
	builder.WriteString("pages/")
	builder.WriteString(shortPath)
	if path[len(path)-1:] != "/" {
		builder.WriteString("/")
	}
	builder.WriteString("meta")
	finalPath := builder.String()
	metafile, err := ioutil.ReadFile(builder.String())
	if err != nil {
		if path == "/" {
			return "", fmt.Errorf("No file found: %v: %w", finalPath, errMetaFatal404)
		}
		return "", fmt.Errorf("File not found: %v: %w", finalPath, errMeta404)
	}
	return string(metafile), nil
}
