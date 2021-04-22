
package GUI

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"log"
)


var tempEvent gwu.Event
var AppName = "图片合成器"
type demo struct {
	link      gwu.Label
	buildFunc func(gwu.Event) gwu.Comp
	comp      gwu.Comp // Lazily initialized demo comp
}
type pdemo *demo

func buildHome(sess gwu.Session) {
	win := gwu.NewWindow("show", AppName)

	win.Style().SetFullSize()
	win.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETypeWinLoad:
			//log.Println("LOADING window:", e.Src().ID())
		case gwu.ETypeWinUnload:
			//log.Println("UNLOADING window:", e.Src().ID())
		}
	}, gwu.ETypeWinLoad, gwu.ETypeWinUnload)

	hiddenPan := gwu.NewNaturalPanel()
	sess.SetAttr("hiddenPan", hiddenPan)

	header := gwu.NewHorizontalPanel()
	header.Style().SetFullWidth().SetBorderBottom2(2, gwu.BrdStyleSolid, "#cccccc")
	title := gwu.NewLink(AppName, win.Name())
	title.SetTarget("")
	title.Style().SetColor(gwu.ClrBlue).SetFontWeight(gwu.FontWeightBold).SetFontSize("120%").Set("text-decoration", "none")
	header.Add(title)

	header.AddHConsumer()

	header.AddHSpace(100)
	setNoWrap(header)
	win.Add(header)

	content := gwu.NewHorizontalPanel()
	content.SetCellPadding(1)
	content.SetVAlign(gwu.VATop)
	content.Style().SetFullSize()

	demoWrapper := gwu.NewPanel()
	demoWrapper.Style().SetPaddingLeftPx(5)
	demoWrapper.AddVSpace(10)
	demoTitle := gwu.NewLabel("")
	demoTitle.Style().SetFontWeight(gwu.FontWeightBold).SetFontSize("100%")
	demoWrapper.Add(demoTitle)
	demoWrapper.AddVSpace(10)

	links := gwu.NewPanel()
	links.SetCellPadding(1)
	links.Style().SetPaddingRightPx(5)

	demos := make(map[string]pdemo)
	var selDemo pdemo

	selectDemo := func(d pdemo, e gwu.Event) {
		if selDemo != nil {
			selDemo.link.Style().SetBackground("")
			if e != nil {
				e.MarkDirty(selDemo.link)
			}
			demoWrapper.Remove(selDemo.comp)
		}
		selDemo = d
		d.link.Style().SetBackground("#88ff88")
		demoTitle.SetText(d.link.Text())
		if d.comp == nil {
			d.comp = d.buildFunc(e)
		}
		demoWrapper.Add(d.comp)
		if e != nil {
			e.MarkDirty(d.link, demoWrapper)
		}
	}

	createDemo := func(name string, buildFunc func(gwu.Event) gwu.Comp) pdemo {
		link := gwu.NewLabel(name)
		link.Style().SetFullWidth().SetCursor(gwu.CursorPointer).SetDisplay(gwu.DisplayBlock).SetColor(gwu.ClrBlue)
		demo := &demo{link: link, buildFunc: buildFunc}
		link.AddEHandlerFunc(func(e gwu.Event) {
			selectDemo(demo, e)
			if tempEvent == nil {
				tempEvent = e
			}
		}, gwu.ETypeClick)
		links.Add(link)
		demos[name] = demo
		return demo
	}

	links.Style().SetFullHeight().SetBorderRight2(2, gwu.BrdStyleSolid, "#cccccc")
	links.AddVSpace(10)
	createDemo("介绍", buildIntroduction)
	//selectDemo(homeDemo, nil)
	//links.AddVSpace(5)


	//=============================================//
	//l0 := gwu.NewLabel("账户")
	//l0.Style().SetFontWeight(gwu.FontWeightBold)
	//links.Add(l0)
	//links.AddVSpace(10)
	//createDemo("激活", buildActiveUI)
	//=============================================//


	//l := gwu.NewLabel("bilibili")
	//l.Style().SetFontWeight(gwu.FontWeightBold)
	//links.Add(l)
	links.AddVSpace(10)
	home := createDemo("配置", buildHomeUI)
	selectDemo(home, nil)
	//=============================================//


	links.AddVConsumer()
	setNoWrap(links)
	content.Add(links)
	content.Add(demoWrapper)
	content.CellFmt(demoWrapper).Style().SetFullWidth()

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

	footer := gwu.NewHorizontalPanel()
	footer.Style().SetFullWidth().SetBorderTop2(2, gwu.BrdStyleSolid, "#cccccc")
	footer.Add(hiddenPan)
	footer.AddHConsumer()
	l := gwu.NewLabel("自媒体软件工作室")
	l.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(l)
	footer.AddHSpace(10)
	link := gwu.NewLink("点击进入", "https://mediablog.github.io/")
	link.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(link)
	setNoWrap(footer)
	win.Add(footer)

	sess.AddWin(win)
}

// setNoWrap sets WhiteSpaceNowrap to all children of the specified panel.
func setNoWrap(panel gwu.Panel) {
	count := panel.CompsCount()
	for i := count - 1; i >= 0; i-- {
		panel.CompAt(i).Style().SetWhiteSpace(gwu.WhiteSpaceNowrap)
	}
}

// SessHandler is our session handler to build the showcases window.
type sessHandler struct{}

func (h sessHandler) Created(s gwu.Session) {
	buildHome(s)
}

func (h sessHandler) Removed(s gwu.Session) {}

// StartServer creates and starts the Gowut GUI server.
func StartServer(appName, addr string) {
	// Create GUI server
	server := gwu.NewServer(appName, addr)

	server.SetText(AppName)

	server.AddSessCreatorName("show", "点击进入")
	server.AddSHandler(sessHandler{})

	fmt.Println("软件已经启动，请不要退出，可以最小化本窗口")
	if err := server.Start("show"); err != nil {
		log.Println("Error: Cound not start GUI server:", err)
		return
	}
}
