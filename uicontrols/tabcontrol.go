package uicontrols

import (
	"github.com/gazercloud/gazerui/uievents"
	"github.com/gazercloud/gazerui/uiinterfaces"
	"image"
)

type TabControl struct {
	Container

	headerButtons         []*Button
	pages                 []*TabPage
	visiblePages          []int
	currentPageIndex      int
	closeButtonIndexHover int

	hoverTabIndex int

	header     *Panel
	pagesPanel *Panel

	OnNeedClose func(index int)
}

type TabPage struct {
	Panel
	ShowImage       bool
	ShowText        bool
	Img             image.Image
	text            string
	ShowCloseButton bool
	tabControl      *TabControl
	headerButton    *Button
}

func NewTabControl(parent uiinterfaces.Widget) *TabControl {
	var c TabControl
	c.InitControl(parent, &c)
	c.pages = make([]*TabPage, 0)

	c.header = NewPanel(&c)
	c.AddWidgetOnGrid(c.header, 0, 0)
	c.header.SetCellPadding(0)
	c.header.SetPanelPadding(0)
	//c.header.AddHSpacerOnGrid(100, 0)

	c.pagesPanel = NewPanel(&c)
	c.AddWidgetOnGrid(c.pagesPanel, 0, 1)
	c.pagesPanel.SetCellPadding(0)
	c.pagesPanel.SetPanelPadding(0)

	c.visiblePages = make([]int, 0)

	c.xExpandable = true
	c.yExpandable = true

	return &c
}

func (c *TabPage) Dispose() {
	c.Panel.Dispose()
}

func (c *TabPage) ControlType() string {
	return "TabPage"
}

func (c *TabPage) SetText(text string) {
	c.text = text
	c.headerButton.SetText(text)
	c.Update("TabPage")
}

func (c *TabPage) SetVisible(visible bool) {
	c.headerButton.SetVisible(visible)
	c.Panel.SetVisible(visible)
}

func (c *TabPage) MouseMove(event *uievents.MouseMoveEvent) {
	c.Panel.MouseMove(event)
}

func (c *TabControl) Dispose() {
	for _, p := range c.pages {
		p.Dispose()
	}

	c.Control.Dispose()

	c.pages = nil
	c.OnNeedClose = nil
}

func (c *TabControl) ControlType() string {
	return "TabControl"
}

func (c *TabControl) AddPage() *TabPage {
	var t TabPage
	t.InitControl(c, &t)
	t.SetWindow(c.OwnWindow)
	pageIndex := len(c.pages)
	c.pagesPanel.AddWidgetOnGrid(&t, pageIndex, 0)
	btn := c.header.AddButtonOnGrid(pageIndex, 0, "TabName", func(event *uievents.Event) {
		c.SetCurrentPage(event.Sender.(*Button).UserData("index").(int))
	})
	btn.SetUserData("index", len(c.pages))
	c.headerButtons = append(c.headerButtons, btn)

	c.pages = append(c.pages, &t)

	t.ShowText = true
	t.ShowImage = false
	t.tabControl = c
	t.headerButton = btn

	if len(c.pages) == 1 {
		c.SetCurrentPage(0)
	} else {
		c.SetCurrentPage(c.currentPageIndex)
	}

	t.SetYExpandable(true)

	return &t
}

func (c *TabControl) RemovePage(index int) {

	/*c.pages[index].Dispose()

	c.pages = append(c.pages[:index], c.pages[index+1:]...)
	if c.currentPageIndex >= len(c.pages) {
		c.currentPageIndex = len(c.pages) - 1
	}

	c.Update("TabControl")*/
}

func (c *TabControl) Page(index int) *TabPage {
	return c.pages[index]
}

func (c *TabControl) SetCurrentPage(index int) {
	if index >= 0 && index < len(c.pages) {
		for i, page := range c.pages {
			if index == i {
				page.Panel.SetVisible(true)
			} else {
				page.Panel.SetVisible(false)
			}
		}
		c.currentPageIndex = index
		c.Update("TabControl")
	}
}

func (c *TabControl) PagesCount() int {
	return len(c.pages)
}

func (c *TabControl) Tooltip() string {
	if c.hoverTabIndex > -1 {
		if c.hoverTabIndex < len(c.pages) {
			return c.pages[c.hoverTabIndex].text
		}
	}
	return ""
}
