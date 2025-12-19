package server

import (
	"net/url"
	"path"
	"strings"
)

// cleanPath 防止路径穿越，但允许目录访问
func cleanPath(p string) string {
	decoded, _ := url.PathUnescape(p)
	clean := path.Clean("/" + decoded)
	if strings.Contains(clean, "..") {
		return "/"
	}
	return clean
}
