// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
	"github.com/jrmsdev/jcms/internal/httpd/handler"
)

func TestMain(m *testing.M) {
	handler.TestingMode()
	test.Main(m, "testing")
}
