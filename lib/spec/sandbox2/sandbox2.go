package sandbox2

import (
	strutils "github.com/torden/go-strutil"
)

func HelloWorld() string {
	strutil := strutils.NewStringProc()
	return strutil.StripSlashes("Hello World")
}
