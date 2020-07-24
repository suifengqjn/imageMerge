package GUI

import (
	"github.com/icza/gowut/gwu"
)

func buildIntroduction(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	p.AddVSpace(20)
	link := gwu.NewLink("软件地址", "https://mediablog.github.io/2020/06/13/videoeditorpro/")
	p.Add(link)

	p.AddVSpace(20)
	link = gwu.NewLink("官网", "https://mediablog.github.io/")
	p.Add(link)


	p.Add(gwu.NewLabel("关注微信公众号，可实时接收软件版本更新信息，后续也会不断在公众号内分享免费软件和自媒体资料"))

	p.AddVSpace(20)
	img := gwu.NewImage("公众号", "http://cdn.qiniu.freetop.ren/gzh.jpg")
	img.Style().SetSizePx(200, 200)
	p.Add(img)

	return p
}

