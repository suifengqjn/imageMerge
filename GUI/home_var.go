package GUI

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	cm "myProject/videoCli/common"
	"strconv"
	"strings"
)

var (
	PreSizeCb gwu.CheckBox
	PreSizeW  gwu.TextBox
	PreSizeH  gwu.TextBox
)

/*
	Size        float64      // 字号
	Bold        bool         // 是否加粗
	TextColor   string       // 文字颜色
	StrokeColor string       // 文字描边颜色
	Pre         string       // 文字前缀
	Suf         string       // 文字后缀
	FontPath    string       // 字体路径
	LayoutStyle int          // 1 靠上居中  2 靠右居中  3 靠下居中  4 靠坐居中  5 左上角 6 右上角 7左下角 8右下角
	LayoutSpan  int          // 与 上右下 左的间距
	VSpan       int          // 垂直间距
	HSpan       int          // 水平间距
*/
var (
	PreTextCb            gwu.CheckBox
	PreTextSizeTb        gwu.TextBox
	PreTextBoldCb        gwu.CheckBox
	PreTextColorTb       gwu.TextBox
	PreTextStrokeColorTb gwu.TextBox
	PreTextTb         gwu.TextBox
	PreTextSufTb         gwu.TextBox
	PreTextFontLb        gwu.ListBox
	PreTextStyleLb       gwu.ListBox
	PreTextHSpanTb       gwu.TextBox
	PreTextVSpanTb       gwu.TextBox
)

/*
	Switch bool
	Style       int    //拼接样式
	Count       int    //几张图片拼在一起
	BjColor string //背景颜色 16进制
	LSpan int
	RSpan int
	TSpan int
	BSpan int
	HSpan int //图片与图片横向间距
	VSpan int //图片与图片纵向间距

*/

var (
	ComposeCb        gwu.CheckBox
	ComposeStyleLb   gwu.ListBox
	ComposeBjColorTb gwu.TextBox
	ComposeLSpanTb   gwu.TextBox
	ComposeRSpanTb   gwu.TextBox
	ComposeTSpanTb   gwu.TextBox
	ComposeBSpanTb   gwu.TextBox
	ComposeHSpanTb   gwu.TextBox
	ComposeVSpanTb   gwu.TextBox
)

var (
	AfterTextCb            gwu.CheckBox
	AfterTextSizeTb        gwu.TextBox
	AfterTextBoldCb        gwu.CheckBox
	AfterTextColorTb       gwu.TextBox
	AfterTextStrokeColorTb gwu.TextBox
	AfterTextTb            gwu.TextBox
	AfterTextSufTb         gwu.TextBox
	AfterTextFontLb        gwu.ListBox
	AfterTextStyleLb       gwu.ListBox
	AfterTextHSpanTb       gwu.TextBox
	AfterTextVSpanTb       gwu.TextBox
)

/*
	Switch      bool
	LayoutStyle int    //
	VSpan       int    // 垂直间距
	HSpan       int    // 水平间距
	LogoPath    string //logo路径
*/

var (
	LogoCb      gwu.CheckBox
	LogoStyleLb gwu.ListBox
	LogoHSpanTb gwu.TextBox
	LogoVSpanTb gwu.TextBox
	LogoPathLb  gwu.ListBox
)

var (
	AfterSizeCb gwu.CheckBox
	AfterSizeW  gwu.TextBox
	AfterSizeH  gwu.TextBox
)

func fillWithConfig(con *cm.HVConfig, e gwu.Event) {

	//fonts, keys := common.LoadFonts()
	//// 文字描边
	//StrokeTextCb.SetState(con.StrokeText.Switch)
	//StrokeTextTb.SetText(con.StrokeText.Content)
	//StrokeTextSizeTb.SetText(str(con.StrokeText.Size))
	//if _, ok := common.TextColorsMap[con.StrokeText.Color]; ok {
	//	index := common.TextColorsMap[con.StrokeText.Color]
	//	StrokeTextColorLb.SetSelected(index, true)
	//} else {
	//	StrokeTextColorLb.ClearSelected()
	//}
	//
	//StrokeTextFontLb.ClearSelected()
	//for _, v := range fonts {
	//	if v == con.StrokeText.Path {
	//		for i := 0; i < len(keys); i++ {
	//			StrokeTextFontLb.SetSelected(i+1, true)
	//			break
	//		}
	//	}
	//}
	//
	//if _, ok := common.StrokeStyleMap[con.StrokeText.Alignment]; ok {
	//	StrokeTextStyleLb.SetSelected(con.StrokeText.Alignment, true)
	//} else {
	//	StrokeTextStyleLb.ClearSelected()
	//}
	//StrokeTextStrokeWidthTb.SetText(str(con.StrokeText.StrokeWidth))
	//StrokeTextStrokeColorTb.SetText(str(con.StrokeText.StrokeColor))
	//e.MarkDirty(StrokeTextCb)
	//e.MarkDirty(StrokeTextTb)
	//e.MarkDirty(StrokeTextSizeTb)
	//e.MarkDirty(StrokeTextColorLb)
	//e.MarkDirty(StrokeTextFontLb)
	//e.MarkDirty(StrokeTextStyleLb)
	//e.MarkDirty(StrokeTextStrokeWidthTb)
	//e.MarkDirty(StrokeTextStrokeColorTb)
	//
	//// 16 bgm
	//BgmCb.SetState(con.AddBgm.Switch)
	//BgmFrontVolumeTb.SetText(fmt.Sprintf("%v", con.AddBgm.FrontVolume))
	//BgmBackVolumeTb.SetText(fmt.Sprintf("%v", con.AddBgm.BackVolume))
	//BgmCoverCb.SetState(con.AddBgm.Cover)
	//e.MarkDirty(BgmCb)
	//e.MarkDirty(BgmFrontVolumeTb)
	//e.MarkDirty(BgmBackVolumeTb)
	//e.MarkDirty(BgmCoverCb)
	//
	//films, keys := common.LoadFilms()
	//
	//// 17
	//filmHeadCb.SetState(con.FilmHead.Switch)
	//filmHeadLb.ClearSelected()
	//for _, v := range films {
	//	if v == con.FilmHead.Path {
	//		for i := 0; i < len(keys); i++ {
	//			filmHeadLb.SetSelected(i+1, true)
	//			break
	//		}
	//	}
	//}
	//e.MarkDirty(filmHeadCb)
	//e.MarkDirty(filmHeadLb)
	//
	//filmFootCb.SetState(con.FilmFoot.Switch)
	//filmFootLb.ClearSelected()
	//for _, v := range films {
	//	if v == con.FilmFoot.Path {
	//		for i := 0; i < len(keys); i++ {
	//			filmFootLb.SetSelected(i+1, true)
	//			break
	//		}
	//	}
	//}
	//e.MarkDirty(filmFootCb)
	//e.MarkDirty(filmFootLb)

}

func str(v interface{}) string {

	if fmt.Sprintf("%v", v) == "0" {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func intValue(v string) int {
	v = strings.TrimSpace(v)
	value, err := strconv.Atoi(v)
	if err != nil {
		valueF, err := strconv.ParseFloat(v, 10)
		if err != nil {
			return 0
		} else {
			return int(valueF)
		}
	}
	return value
}

func floatValue(v string) float64 {
	valueF, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0
	}
	return valueF
}
