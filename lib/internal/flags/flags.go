// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ShowVersion bool
	Log         string
	Debug       bool
	Quiet       bool
	intHttpPort int
	HttpPort    string
	Webapp      string
	Assetsdir   string
	Datadir     string
)

var sprintf = fmt.Sprintf

func init() {
	intHttpPort = 6080
	Webapp = os.Getenv("JCMS_WEBAPP")
	if Webapp == "" {
		Webapp = "default"
	}
	Quiet = false
	Debug = false
	Log = os.Getenv("JCMS_LOG")
	if Log == "" {
		Log = "error"
	}
	if Log == "quiet" {
		Quiet = true
	}
	if Log == "debug" {
		Debug = true
	}
	Assetsdir = os.Getenv("JCMS_ASSETSDIR")
	if Assetsdir == "" {
		Assetsdir = filepath.FromSlash("/srv/jcms/assets")
	}
	Datadir = os.Getenv("JCMS_DATADIR")
	if Datadir == "" {
		Datadir = filepath.FromSlash("/srv/jcms/data")
	}
	flag.BoolVar(&ShowVersion, "version", false, "show version")
	flag.BoolVar(&Debug, "debug", Debug, "enable debug")
	flag.BoolVar(&Quiet, "quiet", Quiet, "quiet mode")
	flag.IntVar(&intHttpPort, "port", intHttpPort, "http port `number`")
	flag.StringVar(&Webapp, "webapp", Webapp, "`webapp` name")
	flag.StringVar(&Assetsdir, "assets", Assetsdir, "assets `directory`")
	flag.StringVar(&Datadir, "data", Datadir, "data `directory`")
}

func Parse() {
	flag.Parse()
	if Quiet {
		Log = "quiet"
	}
	if Debug {
		Log = "debug"
	}
	HttpPort = sprintf("%d", intHttpPort)
}
