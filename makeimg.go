package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/olebedev/config"

)

func main() {
	//YAML IS EATEN HERE
	cfg, err := config.ParseYaml(yamlString)
	host, err := cfg.String("development.database.host")
	var t int = 0
	fmt.Println("Enter the number of the board you want to build for")
	files, err := ioutil.ReadDir("./boards")
	if err != nil {
	log.Fatal(err)
	return err
	}


	//LISTS  THE CONFIG FILES AND PROMPTS USER TO CHOOSE A CONFIG NUMBER
	fmt.Println("TYPE THE NUMBER OF THE ROOTFS YOU WANT TO USE. This will be made dynamic soon. ")
	for _, e := {
		fmt.Println(t, e)
	}

	for _, file := range files {
		t = t + 1
		fmt.Println(t, " ", file.Name())
		var e []t
	}

	//BUILDS OUT THE FOLDER STRUCTURE NEEDED FOR BUILDING
	session := sh.NewSession()
	session.ShowCMD = true
	session.Command("mkdir /builds && echo [ OK ] || 'echo [ OK }' ")
	session.Command("mkdir /builds/boards || echo '/builds/boards already present [ OK ]' )
	session.Command("mkdir /builds/images || echo '[ OK ]' ")
	session.Command("mkdir /builds/kernels")
	session.Command("mkdir /builds/u-boots")
	session.Command("mkdir /builds/rootfses")
	session.Command("mkdir /buidls/u-boots")
	session.SetDir("/builds")
	session.Command("echo hello")
	session.Command("git clone ")
	return err

	//Goroutines fetch the neeeded files
	go Ubuntu()
	go Debian()
	go Arch()
	go Pecovnik()


	}

	app.Run(os.Args)
}

//Downloads an ubuntu rootfs
func Ubuntu()
	{
	sh.Command("wget http://cdimage.ubuntu.com/ubuntu-core/xenial/daily-preinstalled/current/xenial-preinstalled-core-armhf.device.tar.gz")
	}


//Downloads a Debian rootfs
func Debian()
	 {
	sh.Command("wget https://rcn-ee.com/rootfs/2016-04-07/elinux/debian-8.4-console-armhf-2016-04-07.tar.xz")
	}


//Downloads an Arch rootfs
func Arch() {
	sh.Command("wget http://sg2.mirror.archlinuxarm.org/os/ArchLinuxARM-armv7-latest.tar.gz")
}

//Downloads one of Uncle Igor's ROMs to use the structure he's so kindly built for us, and mounts the ROM Image
func Pecovnik(){
	sh.Command("wget http://mirror.igorpecovnik.com/Armbian_5.10_Orangepih3_Debian_jessie_3.4.112.7z")
}


