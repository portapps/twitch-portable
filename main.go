//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
	"github.com/portapps/portapps/pkg/utl"
)

var (
	app *App
)

func init() {
	var err error

	// Init app
	if app, err = New("twitch-portable", "Twitch"); err != nil {
		Log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	app.AppPath = utl.PathJoin(app.RootPath, "Bin")
	app.DataPath = ""
	app.Process = utl.PathJoin(app.AppPath, "Twitch.exe")
	app.WorkingDir = app.AppPath

	// On exit
	defer func() {
		// Remove tmp path
		os.RemoveAll(utl.PathJoin(utl.RoamingPath(), "Twitch"))
	}()

	app.Launch(os.Args[1:])
}
