package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/burntsushi/toml"
	"time"
)







viper.SetConfigName("") // name of config file (without extension)
viper.AddConfigPath(".")               // optionally look for config in the working directory
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
panic(fmt.Errorf("Fatal error config file: %s \n", err))
}

description: "a small computer for $15"
config: "https://github.com/faddat/linux"
tags: [ "mainline, sunxi, allwinner, h3" ]
kerneldefconfig: "sunxi_defconfig"
ubootconfig: "sun8i_h3_orangepi_pc_defconfig"
categories:
- "hasmainlinegpu"
- "VIM"
- "cheap"
- "allwinner"
env:
- "CROSS_COMPILE=arm-linux-gnueabihf-"
- "ARCH=arm"


session := sh.NewSession()
session.ShowCMD = true
session.Command("mkdir /builds")
session.Command("mkdir /builds/boards")
session.Command("mkdir /builds/images")
session.Command("mkdir /builds/kernels")
session.Command("mkdir /builds/u-boots")
session.Command("mkdir /builds/rootfses")
session.Command("mkdir /buidls/u-boots")
session.SetDir("/builds")
session.Command("echo", "hello").Run()
session.Command("git clone ")

