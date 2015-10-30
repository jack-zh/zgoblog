package cmd

import "github.com/jack-zh/zgoblog/fweb"

func Init(app *fweb.App) {
	StartBackupTimer(app, 24)
}
