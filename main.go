package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/burntsushi/toml"
	"time"
)


type Config struct {
	Age int
	Cats []string
	Soc string
	GPUion []int
	Updated time.Time
}



session := sh.NewSession()
session.ShowCMD = true
session.Command("mkdir /builds")
session.Command("mkdir /builds/")
session.SetEnv("BUILD_ID", "123")
session.SetDir("/builds")
session.Command("echo", "hello").Run()
session.Command("git clone ")

