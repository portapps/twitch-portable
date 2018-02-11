//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "twitch-portable"
	Papp.Name = "Twitch"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("Bin")
	Papp.DataPath = ""
	Papp.Process = PathJoin(Papp.AppPath, "Twitch.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	Launch(os.Args[1:])
}
