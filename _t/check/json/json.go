// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package json

import (
	j "encoding/json"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
)

func NotEqual(t *testing.T, blob []byte, key string, expect interface{}, desc string) bool {
	t.Helper()
	var d map[string]interface{}
	err := j.Unmarshal(blob, &d)
	if err != nil {
		t.Fatalf("%s: %s", desc, err)
	}
	//~ t.Log(d)
	v, ok := d[key]
	if !ok {
		t.Logf("%s %s key not found", desc, key)
		return true
	}
	return check.NotEqual(t, v, expect, "JSON " + desc + " " + key)
}
