package uicontrols

import (
	"allece.com/system/core/canvas"
	"allece.com/system/core/go-gl/glfw/v3.3/glfw"
	"allece.com/system/core/uievents"
	"allece.com/system/core/uiinterfaces"
	"allece.com/system/core/uiresources"
	"allece.com/system/core/uistyles"
	"golang.org/x/image/colornames"
	"log"
)

type Dialog struct {
	Panel

	headerPanel  *dialogHeader
	contentPanel *Panel

	title      string
	menuWidth  int
	menuHeight int
	CloseEvent func()

	closed bool

	acceptButton *Button
	rejectButton *Button
	OnAccept     func()
	OnReject     func()

	TryAccept func() bool
	OnShow    func()
}

type dialogHeader struct {
	Panel
	dialog     *Dialog
	headerText *TextBlock
	btnClose   *Button

	mousePointerInRect   bool
	lastMouseDownX       int
	lastMouseDownY       int
	lastMouseDownDialogX int
	lastMouseDownDialogY int
	pressed              bool
}

func NewDialogHeader(parent uiinterfaces.Widget) *dialogHeader {
	var c dialogHeader
	c.InitControl(parent, &c)
	c.headerText = c.AddTextBlockOnGrid(0, 0, "Dialog")
	c.headerText.SetMinHeight(30)
	c.headerText.TextHAlign = canvas.HAlignCenter
	c.SetBorderBottom(1, c.ForeColor())

	c.btnClose = c.AddButtonOnGrid(1, 0, "Close", func(event *uievents.Event) {
		c.dialog.Reject()
	})
	c.btnClose.SetBorders(0, colornames.Red) // icons\material\navigation\drawable-hdpi\ic_close_black_48dp.png
	c.btnClose.SetImage(uiresources.ResImageAdjusted("icons/material/navigation/drawable-hdpi/ic_close_black_48dp.png", c.ForeColor()))
	c.btnClose.imageForeColor = true
	c.btnClose.showText = false
	c.btnClose.rebuildContent()
	c.btnClose.SetMaxWidth(30)
	c.btnClose.onMouseEnter = func() {
		c.btnClose.SetBackColor(colornames.Red)
		c.btnClose.SetForeColor(colornames.White)
	}
	c.btnClose.onMouseLeave = func() {
		c.btnClose.SetBackColor(c.BackColor())
		c.btnClose.SetForeColor(c.ForeColor())
	}

	uistyles.InitDefaultStyle(&c)
	return &c
}

func (c *dialogHeader) Dispose() {
	c.dialog = nil
	c.headerText = nil
	c.btnClose = nil
	c.Panel.Dispose()
}

func (c *dialogHeader) setTitle(title string) {
	c.headerText.SetText(title)
}

func (c *dialogHeader) MouseEnter() {
	c.mousePointerInRect = true
}

func (c *dialogHeader) MouseLeave() {
	c.mousePointerInRect = false
}

func (c *dialogHeader) MouseDown(event *uievents.MouseDownEvent) {
	w := c.Panel.FindWidgetUnderPointer(event.X, event.Y)
	if w == c.btnClose.widget {
		c.Panel.MouseDown(event)
		return
	}

	c.pressed = true

	ev := *event
	posHeaderTextX, posHeaderTextY := c.RectClientAreaOnWindow()
	ev.X += posHeaderTextX
	ev.Y += posHeaderTextY

	c.lastMouseDownX = ev.X
	c.lastMouseDownY = ev.Y
	// Position of dialog in mouse down moment
	c.lastMouseDownDialogX, c.lastMouseDownDialogY = c.dialog.RectClientAreaOnWindow()
	c.lastMouseDownDialogX -= c.dialog.LeftBorderWidth()
	c.lastMouseDownDialogY -= c.dialog.TopBorderWidth()
	//fmt.Println("Dialog Mouse Down: ", ev.X, ev.Y, c.lastMouseDownDialogX, c.lastMouseDownDialogY)
}

func (c *dialogHeader) MouseMove(event *uievents.MouseMoveEvent) {
	w := c.Panel.FindWidgetUnderPointer(event.X, event.Y)
	if w == c.btnClose.widget {
		c.Panel.MouseMove(event)
		return
	}

	if c.pressed {
		ev := *event
		posHeaderTextX, posHeaderTextY := c.RectClientAreaOnWindow()
		ev.X += posHeaderTextX
		ev.Y += posHeaderTextY

		deltaX := ev.X - c.lastMouseDownX
		deltaY := ev.Y - c.lastMouseDownY
		c.dialog.SetX(c.lastMouseDownDialogX + deltaX)
		c.dialog.SetY(c.lastMouseDownDialogY + deltaY)
		//fmt.Println("Dialog Mouse Move: ", deltaX, deltaY, c.lastMouseDownDialogX, c.lastMouseDownDialogY)
	}
}

func (c *dialogHeader) MouseUp(event *uievents.MouseUpEvent) {
	w := c.Panel.FindWidgetUnderPointer(event.X, event.Y)
	if w == c.btnClose.widget {
		c.Panel.MouseUp(event)
		return
	}

	if c.pressed {
		c.pressed = false
	}
}

func (c *dialogHeader) FindWidgetUnderPointer(x int, y int) uiinterfaces.Widget {
	w := c.Panel.FindWidgetUnderPointer(x, y)
	if w == c.btnClose.widget {
		return w
	}

	return nil // Event filter
}

func (c *Dialog) OnInit() {
	c.SetParent(c.parent.Window().CentralWidget())
	c.headerPanel = NewDialogHeader(c)
	c.headerPanel.dialog = c
	c.headerPanel.SetPanelPadding(0)
	c.headerPanel.SetCellPadding(0)
	c.AddWidgetOnGrid(c.headerPanel, 0, 0)
	p := c.AddPanelOnGrid(0, 1)
	p.SetCellPadding(0)
	p.SetPanelPadding(0)
	p.AddVSpacerOnGrid(0, 0)
	c.contentPanel = p.AddPanelOnGrid(1, 0)

	c.SetBorders(2, c.ForeColor())
	c.SetName("Dialog")

	c.SetPanelPadding(0)
	c.SetCellPadding(0)
}

func (c *Dialog) Dispose() {
	c.headerPanel = nil
	c.contentPanel = nil
	c.CloseEvent = nil
	c.acceptButton = nil
	c.rejectButton = nil
	c.OnAccept = nil
	c.OnReject = nil
	c.TryAccept = nil
	c.OnShow = nil
	c.Panel.Dispose()
}

func (c *Dialog) Close() {
	c.Reject()
}

func (c *Dialog) ControlType() string {
	return "Dialog"
}

func (c *Dialog) ShowDialog() {
	x := (c.Window().Width() - c.Width()) / 2
	y := (c.Window().Height() - c.Height()) / 2
	c.ShowDialogAtPos(x, y)
}

func (c *Dialog) ShowDialogAtPos(x, y int) {
	c.SetX(x)
	c.SetY(y)
	c.Window().AppendPopup(c.widget)
	c.Focus()

	if c.OnShow != nil {
		c.OnShow()
	}
}

func (c *Dialog) ContentPanel() *Panel {
	return c.contentPanel
}

func (c *Dialog) Resize(w, h int) {
	c.SetWidth(w)
	c.SetHeight(h)
}

func (c *Dialog) SetAcceptButton(acceptButton *Button) {
	c.acceptButton = acceptButton
	acceptButton.onMouseClick = func(event *uievents.MouseClickEvent) {
		c.Accept()
	}
}

func (c *Dialog) SetRejectButton(rejectButton *Button) {
	c.rejectButton = rejectButton
	rejectButton.onMouseClick = func(event *uievents.MouseClickEvent) {
		c.Reject()
	}
}

func (c *Dialog) SetTitle(title string) {
	c.title = title
	c.headerPanel.setTitle(title)
}

func (c *Dialog) ClosePopup() {
	if c.CloseEvent != nil {
		c.CloseEvent()
	}
}

func (c *Dialog) Accept() {
	log.Println("Dialog::Accept 1 ", c.title)

	if c.title == "Edit sensor" {
		log.Println("Dialog::Accept 1 ", c.title)
	}

	if c.closed {
		return
	}

	if c.TryAccept != nil {
		if !c.TryAccept() {
			return
		}
	}

	log.Println("Dialog::Accept 2 ", c.title)

	onAccept := c.OnAccept
	c.Window().CloseTopPopup()
	c.closed = true
	if onAccept != nil {
		log.Println("Dialog::Accept 3 ", c.title)
		onAccept()
	}
	log.Println("Dialog::Accept 4 ", c.title)
}

func (c *Dialog) Reject() {
	if c.closed {
		return
	}

	onReject := c.OnReject

	c.Window().CloseTopPopup()
	c.closed = true
	if onReject != nil {
		onReject()
	}
}

func (c *Dialog) KeyDown(event *uievents.KeyDownEvent) bool {
	if event.Key == glfw.KeyEnter || event.Key == glfw.KeyKPEnter {
		if c.acceptButton != nil {
			c.acceptButton.Press()
		} else {
		}
		c.Accept()
		return true
	}
	if event.Key == glfw.KeyEscape {
		if c.rejectButton != nil {
			c.rejectButton.Press()
		} else {
		}
		c.Reject()
		return true
	}
	return false
}
