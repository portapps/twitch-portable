//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "twitch-portable"
	Papp.Name = "Twitch"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")
	Papp.Process = PathJoin(Papp.AppPath, "Bin", "Twitch.exe")
	Papp.Args = nil
	Papp.WorkingDir = Papp.AppPath

	roamingPath := CreateFolder(PathJoin(Papp.DataPath, "AppData", "Roaming", "Twitch"))
	Log.Infof("Roaming path: %s", roamingPath)

	OverrideEnv("USERPROFILE", Papp.DataPath)
	Launch()
}
