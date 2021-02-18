package main

import (
	"github.com/gazercloud/gazerui/ui"
	"github.com/gazercloud/gazerui/uiforms"
)

type MainForm struct {
	uiforms.Form
}

func (c *MainForm) OnInit() {
	c.SetTitle("Text UI application")
	c.Resize(400, 400)
}

func main() {
	ui.InitUISystem()
	var mainForm MainForm
	uiforms.StartMainForm(&mainForm)
}
