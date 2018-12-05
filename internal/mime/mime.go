// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mime

import (
	"fmt"
	"strings"
	mimelib "mime"
)

func TypeByExtension(x string) string {
	var typ string
	typId, found := dbExts[x]
	if found {
		typ = dbTypes[typId]
	} else {
		typ = mimelib.TypeByExtension(x)
	}
	if typ == "" {
		return "application/octet-stream"
	}
	if strings.HasPrefix(typ, "text/") ||
		typ == "application/javascript" {
		typ = fmt.Sprintf("%s; charset=utf-8", typ)
	}
	return typ
}
