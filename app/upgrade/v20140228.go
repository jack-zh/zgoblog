package upgrade

import (
	"github.com/fuxiaohei/GoInk"
	"github.com/jack-zh/zGoBlog/app/cmd"
	"github.com/jack-zh/zGoBlog/app/model"
)

func init() {
	cmd.SetUpgradeScript(20140228, upgrade_20140228)
}

func upgrade_20140228(_ *GoInk.App) bool {

	// change settings
	model.LoadSettings()
	model.SetSetting("popular_size", "4")
	model.SetSetting("recent_comment_size", "3")
	model.SetSetting("theme_cache", "false")
	model.SyncSettings()

	// overwrite zip bundle bytes
	cmd.ExtractBundleBytes()
	return true
}
