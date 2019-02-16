// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	htpl "html/template"
	"strings"
)

var utils = htpl.FuncMap{
	"join": strings.Join,
}
