package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/codegangsta/cli"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "linuxbuild"
	var t int = 0
	app.Action = func(c *cli.Context) error {
		fmt.Println("TYPE THE NAME OF THE BOARD YOU WANT TO BUILD FOR")
		files, err := ioutil.ReadDir("./boards")
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, file := range files {
			t = t + 1
			fmt.Println(t, " ", file.Name())
		}
		fmt.Println("TYPE THE NUMBER OF THE ROOTFS YOU WANT TO USE")
		fmt.Println(""()
		fmt.Println(t, " ", file.Name())
		fmt.Println(t, " ", file.Name())
		fmt.Println(t, " ", file.Name())
		go StupidFile()


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
		session.Command("echo hello")
		session.Command("git clone ")
		return err
	}

	app.Run(os.Args)
}

func StupidFile() {
	for ;; {
	fmt.Println("GO routine")
	}

}