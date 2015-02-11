package cmd

import "github.com/jack-zh/zGoBlog/fweb"

func Init(app *fweb.App) {
	StartBackupTimer(app, 24)
}
