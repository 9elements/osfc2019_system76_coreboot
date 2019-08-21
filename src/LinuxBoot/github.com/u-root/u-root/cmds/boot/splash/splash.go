package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	go func() {
		baseCmd, err := exec.LookPath("mplayer")
		if err != nil {
			log.Panicln("Can't find mplayer")
			os.Exit(2)
		}
		var loadCommands []string
		loadCommands = append(loadCommands, "-fs")
		loadCommands = append(loadCommands, "-zoom")
		loadCommands = append(loadCommands, "-xy")
		loadCommands = append(loadCommands, "1920")
		loadCommands = append(loadCommands, "-vo")
		loadCommands = append(loadCommands, "fbdev2")
		loadCommands = append(loadCommands, "/video.webm")

		cmdLoad := exec.Command(baseCmd, loadCommands...)
		if err := cmdLoad.Run(); err != nil {
			log.Panicln(err.Error())
			os.Exit(3)
		}
	        var clearCommands []string
	        clearCommands = append(clearCommands, "if=/dev/zero")
	        clearCommands = append(clearCommands, "of=/dev/fb0")
	        clearCommands = append(clearCommands, "bs=1024")
	        clearCommands = append(clearCommands, "count=10000")

	        baseCmd, err = exec.LookPath("dd")
	        cmdLoad = exec.Command(baseCmd, clearCommands...)
		cmdLoad.Run()
	}()
	baseCmd, err := exec.LookPath("/bbin/uinit")
	if err != nil {
		baseCmd, err = exec.LookPath("/bin/unit")
	}

	var bootCommands []string
	bootCommands = append(bootCommands, "-q")

	cmdLoad := exec.Command(baseCmd, bootCommands...)
	cmdLoad.Run()

	os.Exit(1)
}
