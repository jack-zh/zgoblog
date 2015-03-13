package app

import (
	"fmt"
	"github.com/jack-zh/zGoBlog/app/cmd"
	"github.com/jack-zh/zGoBlog/app/handler"
	"os"
)

const (
	__version__ = "zGoBlog version 0.3 build in (2015 03 14)"
	__helpStr__ = `
COMMANDS:
	zGoBlog install    -- Install zGoBlog
	zGoBlog backup     -- Backup data
	zGoBlog update     -- update zip

GLOBAL OPTIONS:
	zGoBlog version    -- Show zGoBlog version
	zGoBlog help       -- Show usage
`
)

// Cmd starts command line application.
// It captures command line arguments and executes proper operation.
// Some operations will exit application when finished.
func Cmd() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "install":
			cmd.DoInstall()
		case "update":
			file, _ := cmd.DoBackup(App, false)
			cmd.DoUpdateZipBytes(file)
		case "backup":
			cmd.DoBackup(App, true)
		case "help":
			fmt.Println(__helpStr__)
		case "version":
			fmt.Println(__version__)
		default:
			fmt.Println(__helpStr__)
		}
		os.Exit(0)
	}
	// do install and run server together
	if !cmd.CheckInstall() {
		cmd.DoInstall()
		return
	}
	// check app version
	if cmd.CheckUpgrade(VERSION, true) {
		os.Exit(1)
		return
	}

	// begin cmd init
	cmd.Init(App)
}

func registerCmdHandler() {
	App.Route("GET,POST,DELETE", "/cmd/backup/", handler.Auth, handler.CmdBackup)
	App.Get("/cmd/backup/file/", handler.Auth, handler.CmdBackupFile)

	App.Route("GET,POST,DELETE", "/cmd/message/", handler.Auth, handler.CmdMessage)
	App.Route("GET,DELETE", "/cmd/logs/", handler.Auth, handler.CmdLogs)
	App.Get("/cmd/monitor/", handler.Auth, handler.CmdMonitor)
	App.Route("GET,POST", "/cmd/theme/", handler.Auth, handler.CmdTheme)
	App.Route("GET,POST", "/cmd/reader/", handler.Auth, handler.CmdReader)
}
