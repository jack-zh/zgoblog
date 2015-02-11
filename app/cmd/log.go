package cmd

import (
	"github.com/jack-zh/zGoBlog/fweb"
	"io/ioutil"
	"os"
	"path/filepath"
)

type logItem struct {
	Name       string
	CreateTime int64
	Text       string
}

func GetLogs(app *fweb.App) []*logItem {
	dir := app.Get("log_dir")
	logs := make([]*logItem, 0)
	filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err == nil {
			if info.IsDir() {
				return nil
			}
			ext := filepath.Ext(info.Name())
			if ext != ".log" {
				return nil
			}
			bytes, e := ioutil.ReadFile(filepath.Join(dir, info.Name()))
			if e != nil {
				return nil
			}
			l := new(logItem)
			l.Name = info.Name()
			l.CreateTime = info.ModTime().Unix()
			l.Text = string(bytes)
			logs = append([]*logItem{l}, logs...)
		}
		return nil
	})
	return logs
}

func RemoveLogFile(app *fweb.App, file string) {
	f := filepath.Join(app.Get("log_dir"), file)
	os.Remove(f)
}
