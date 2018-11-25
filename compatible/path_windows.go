// +build windows

package compatible

import (
	"os"
	"strings"
)

var IsMinGW64 = false

func init() {
	msystem := os.Getenv("MSYSTEM")
	if msystem == "MINGW64" {
		IsMinGW64 = true
	}
}

//path fix for mingw64 zsh
func PathFix(path string) string {
	if IsMinGW64 {
		path = strings.Replace(path, "\\", "/", -1)
		if path[1] == ':' {
			path = "/" + strings.ToLower(path[0:1]) + path[2:]
		}
	}
	return path
}
