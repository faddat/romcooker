package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"io/ioutil"
	"log"
)

func main() {
	//Prompts the user to ask which configuration they'd like to use

	var Confignumber int = 0
	fmt.Println("Enter the number of the board you want to build for")
	Boardname, err := ioutil.ReadDir("./boards")
	if err != nil {
	log.Fatal(err)
	return err
	}

	//LISTS  THE CONFIG FILES AND PROMPTS USER TO CHOOSE A CONFIG NUMBER
	fmt.Println("TYPE THE NUMBER OF THE ROOTFS YOU WANT TO USE. This will be made dynamic soon.")
	for _, Confignumber := range Boardname {
		fmt.Println(Confignumber, Boardname)
	}
	for _, file := range Boardname {
		Confignumber = Confignumber + 1
		fmt.Println(Confignumber, " ", file.Name())
	}

	go FolderStructure()
//	go Parse()

	//Goroutines fetch the neeeded files
	go Ubuntu()
	go Debian()
	go Arch()
	go Pecovnik()
	}


//Downloads an ubuntu rootfs
func Ubuntu(){
	sh.Command("wget http://cdimage.ubuntu.com/ubuntu-core/xenial/daily-preinstalled/current/xenial-preinstalled-core-armhf.device.tar.gz")
}

//Downloads a Debian rootfs
func Debian(){
	sh.Command("wget https://rcn-ee.com/rootfs/2016-04-07/elinux/debian-8.4-console-armhf-2016-04-07.tar.xz")
	}


//Downloads an Arch rootfs
func Arch() {
	sh.Command("wget http://sg2.mirror.archlinuxarm.org/os/ArchLinuxARM-armv7-latest.tar.gz")
}

//Downloads and mounts one of Uncle Igor's ROMs to use the structure he's so kindly built for us, and mounts the ROM Image
func Pecovnik(){
	sh.Command("wget http://mirror.igorpecovnik.com/Armbian_5.10_Orangepih3_Debian_jessie_3.4.112.7z")
	sh.Command("")
}

//
//	config.ParseYamlFile()(filename
//	string) (*Config, error)
//}

//BUILDS OUT THE FOLDER STRUCTURE NEEDED FOR BUILDING
func FolderStructure(){
session := sh.NewSession()
session.ShowCMD = true
session.Command("mkdir /builds && echo [ OK ] || 'echo [ OK }' ")
session.Command("mkdir /builds/boards || echo '/builds/boards already present [ OK ]' " )
session.Command("mkdir /builds/images || echo '[ OK ]' ")
session.Command("mkdir /builds/kernels")
session.Command("mkdir /builds/u-boots")
session.Command("mkdir /builds/rootfses")
session.Command("mkdir /buidls/u-boots")
session.SetDir("/builds")
session.Command("echo hello")
//session.Command("git clone ")
if
log.Fatal(err)
return err
}


