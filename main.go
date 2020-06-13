package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	service:= NewNgApp()
	service.mainWin.Show()
	app.Exec()
}
