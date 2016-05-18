package main

import "github.com/codeskyblue/go-sh"

//Idea: Break the bulk of this into .sh files, and execute those, to make it more user-modifiable

func FolderStructure(){
session := sh.NewSession()
session.ShowCMD = true
session.Command("mkdir /romcooker && echo [ OK ] || echo '[ OK }' ")
session.Command("mkdir /romcooker/boards || echo '/romcooker/boards already present [ OK ]' " )
session.Command("mkdir /romcooker/images || echo '[ OK ]' ")
session.Command("mkdir /romcooker/kernels")
session.Command("mkdir /romcooker/u-boots")
session.Command("mkdir /romcooker/rootfses")
session.Command("mkdir /buidls/u-boots")
session.SetDir("/romcooker/kernels")
session.Command("git clone https://github.com/torvalds/linux")
session.Command("cp -R linux linux-next")
session.Command("cp -R linux linux-mainline")
session.SetDir("/romcooker/linux-next")
session.Command("git remote add linux-next https://git.kernel.org/pub/scm/linux/kernel/git/next/linux-next.git")
session.Command("git fetch linux-next")
session.Command("git fetch --tags linux-next")
session.SetDir("/romcooker")
session.Command("cp -R linux-next linux-next-wens-h3")
session.SetDir("/romcooker/linux-next-wens-h3")
session.Command("git remote add wens-linux https://github.com/wens/linux")
session.Command("git fetch h3-emac")
session.Command("git fetch --tags h3-emac")
session.SetDir("/romcooker/u-boots")
session.Command("git clone https://github.com/u-boot/u-boot")
if err != nil {
    log.Fatal(err)
}
}
