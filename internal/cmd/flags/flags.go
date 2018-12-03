// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"

	"github.com/jrmsdev/jcms/webapp/config"
)

var ShowVersion bool
var Debug bool
var Quiet bool
var HttpPort int
var Webapp string

func init() {
	flag.BoolVar(&ShowVersion, "V", false, "show version")
	flag.BoolVar(&Debug, "d", false, "enable debug")
	flag.BoolVar(&Quiet, "q", false, "quiet mode")
	flag.IntVar(&HttpPort, "p", 0, "http `port` (default \"0\")")
	flag.StringVar(&Webapp, "n", "default", "`webapp` name")
}

func Parse() *config.Config {
	flag.Parse()
	cfg := config.New(Webapp)
	if Quiet {
		cfg.Log = "quiet"
	}
	if Debug {
		cfg.Log = "debug"
	}
	cfg.HttpPort = fmt.Sprintf("%d", HttpPort)
	return cfg
}
