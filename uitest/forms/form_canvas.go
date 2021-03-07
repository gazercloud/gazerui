package forms

import (
	"github.com/gazercloud/gazerui/canvas"
	"github.com/gazercloud/gazerui/uicontrols"
	"github.com/gazercloud/gazerui/uiforms"
	"github.com/gazercloud/gazerui/uiinterfaces"
	"golang.org/x/image/colornames"
	"log"
	"time"
)

type FormCanvas struct {
	uiforms.Form
}

type CanvasCtrl struct {
	uicontrols.Control
}

func NewCanvasCtrl(parent uiinterfaces.Widget) *CanvasCtrl {
	var c CanvasCtrl
	//c.InitControl(parent, &c, 0, 0, 0, 0, 0)
	return &c
}

func (c *CanvasCtrl) Draw(ctx *canvas.CanvasDirect) {
	ctx.DrawRect(10, 10, 100, 100, colornames.Red, 1)
	dt1 := time.Now()
	//ctx.DrawString(0, 10, "S", "Roboto", 112, colornames.Black)
	dtSpan := time.Since(dt1)

	log.Print("Time 1111111:", dtSpan)
}

func (c *FormCanvas) OnInit() {
	c.Resize(600, 800)
	c.SetTitle("FormFont")

	//ctrl := NewCanvasCtrl(c.Panel())
	//c.Panel().AddWidgetOnGrid(ctrl, 0, 0)
}
