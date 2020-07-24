package GUI

import (
	"github.com/icza/gowut/gwu"
	"myProject/imageMerge/common"
	"myProject/imageMerge/engine"
	"strings"
)

var GRAY = "#808080"
var showLabel gwu.Label
var defaultShowString = "图片放在images目录,支持两级目录"

func buildHomeUI(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()
	showLabel = gwu.NewLabel(defaultShowString)
	p.Add(showLabel)

	//p.AddVSpace(20)
	//buildLoadBtn(p)

	p.AddVSpace(10)
	buildPreSize(p)

	p.AddVSpace(20)
	buildPreText(p)

	p.AddVSpace(20)
	buildComposeStyle(p)

	p.AddVSpace(20)
	buildAfterText(p)

	p.AddVSpace(20)
	buildLogo(p)

	p.AddVSpace(20)
	buildAfterSize(p)

	p.AddVSpace(10)
	buildActionBtn(p)

	return p
}

func buildLoadBtn(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	bt1 := gwu.NewButton("从配置1加载")
	bt1.AddEHandlerFunc(func(e gwu.Event) {
		con := common.LoadConfig1()
		if con != nil {

		}
	}, gwu.ETypeClick)
	row.Add(bt1)

	row.AddHSpace(200)
	bt2 := gwu.NewButton("从配置2加载")
	bt2.AddEHandlerFunc(func(e gwu.Event) {

		con := common.LoadConfig2()
		if con != nil {

		}
	}, gwu.ETypeClick)
	row.Add(bt2)

	row.AddHSpace(200)
	bt3 := gwu.NewButton("从配置3加载")
	bt3.AddEHandlerFunc(func(e gwu.Event) {
		con := common.LoadConfig3()
		if con != nil {

		}
	}, gwu.ETypeClick)
	row.Add(bt3)
	p.Add(row)
}

func buildPreSize(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	PreSizeCb = gwu.NewCheckBox("统一单张尺寸")
	PreSizeCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreSize.Switch = PreSizeCb.State()
	}, gwu.ETypeClick)

	row.Add(PreSizeCb)

	row.AddHSpace(18)
	content := gwu.NewLabel("宽:")
	row.Add(content)
	PreSizeW = gwu.NewTextBox("")
	PreSizeW.SetMaxLength(200)
	PreSizeW.Style().SetWidthPx(50)
	PreSizeW.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreSizeW.AddEHandlerFunc(func(e gwu.Event) {

		common.VideoCon.PreSize.W = intValue(PreSizeW.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PreSizeW)

	row.AddHSpace(18)
	content = gwu.NewLabel("高:")
	row.Add(content)
	PreSizeH = gwu.NewTextBox("")
	PreSizeH.SetMaxLength(500)
	PreSizeH.Style().SetWidthPx(50)
	PreSizeH.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreSizeH.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreSize.H = intValue(PreSizeH.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PreSizeH)
	p.Add(row)

}


/*
	PreTextCb            gwu.CheckBox
	PreTextSizeTb        gwu.TextBox
	PreTextBoldCb        gwu.CheckBox
	PreTextColorTb       gwu.TextBox
	PreTextStrokeColorTb gwu.TextBox
	PreTextPreTb         gwu.TextBox
	PreTextSufTb         gwu.TextBox
	PreTextFontLb        gwu.TextBox
	PreTextStyleLb       gwu.ListBox
	PreTextHSpanTb       gwu.TextBox
	PreTextVSpanTb       gwu.TextBox
*/

func buildPreText(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	PreTextCb = gwu.NewCheckBox("单张文字")
	PreTextCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.Switch = PreTextCb.State()
	}, gwu.ETypeClick)
	row.Add(PreTextCb)

	row.AddHSpace(18)
	content := gwu.NewLabel("字体大小:")
	row.Add(content)
	PreTextSizeTb = gwu.NewTextBox("")
	PreTextSizeTb.SetMaxLength(500)
	PreTextSizeTb.Style().SetWidthPx(50)
	PreTextSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextSizeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.Size = floatValue(PreTextSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PreTextSizeTb)

	row.AddHSpace(10)
	PreTextBoldCb = gwu.NewCheckBox("粗体")
	PreTextBoldCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.Bold = PreTextBoldCb.State()
	}, gwu.ETypeClick)
	row.Add(PreTextBoldCb)

	row.AddHSpace(10)
	content = gwu.NewLabel("字体颜色")
	row.Add(content)
	PreTextColorTb = gwu.NewTextBox("")
	PreTextColorTb.SetMaxLength(7)
	PreTextColorTb.Style().SetWidthPx(50)
	PreTextColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.TextColor = PreTextColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PreTextColorTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("描边颜色")
	row.Add(content)
	PreTextStrokeColorTb = gwu.NewTextBox("")
	PreTextStrokeColorTb.SetMaxLength(7)
	PreTextStrokeColorTb.Style().SetWidthPx(50)
	PreTextStrokeColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextStrokeColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.StrokeColor = PreTextStrokeColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PreTextStrokeColorTb)

	p.Add(row)

	//-------------------//

	row1 := gwu.NewHorizontalPanel()
	row1.AddHSpace(10)
	content = gwu.NewLabel("前缀")
	row1.Add(content)
	PreTextTb = gwu.NewTextBox("")
	PreTextTb.SetMaxLength(7)
	PreTextTb.Style().SetWidthPx(50)
	PreTextTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.Pre = PreTextTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(PreTextTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("后缀")
	row1.Add(content)
	PreTextSufTb = gwu.NewTextBox("")
	PreTextSufTb.SetMaxLength(7)
	PreTextSufTb.Style().SetWidthPx(50)
	PreTextSufTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextSufTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.Suf = PreTextSufTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(PreTextSufTb)

	row1.AddHSpace(10)
	label := gwu.NewLabel("选择字体:")
	row1.Add(label)
	fonts, keys := common.LoadFonts()
	arr := []string{"默认"}
	arr = append(arr, keys...)
	PreTextFontLb = gwu.NewListBox(arr)
	common.VideoCon.PreText.FontPath = common.DefaultFont()
	PreTextFontLb.AddEHandlerFunc(func(e gwu.Event) {
		if PreTextFontLb.SelectedIdx() > 0 {
			common.VideoCon.PreText.FontPath = fonts[PreTextFontLb.SelectedValue()]
		}

	}, gwu.ETypeChange)
	row1.Add(PreTextFontLb)

	row1.AddHSpace(10)
	PreTextStyleLb = gwu.NewListBox(common.WaterStyle)
	PreTextStyleLb.AddEHandlerFunc(func(e gwu.Event) {

		//fmt.Println(WaterTextStyleLb.SelectedIdx(),WaterTextStyleLb.SelectedValue())
		for k, v := range common.WaterStyleMap {
			if v == strings.TrimSpace(PreTextStyleLb.SelectedValue()) {
				common.VideoCon.PreText.LayoutStyle = k
			}
		}

	}, gwu.ETypeChange)
	row1.Add(PreTextStyleLb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("水平间距")
	row1.Add(content)
	PreTextHSpanTb = gwu.NewTextBox("")
	PreTextHSpanTb.SetMaxLength(7)
	PreTextHSpanTb.Style().SetWidthPx(50)
	PreTextHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextHSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.HSpan = intValue(PreTextHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(PreTextHSpanTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("垂直间距")
	row1.Add(content)
	PreTextVSpanTb = gwu.NewTextBox("")
	PreTextVSpanTb.SetMaxLength(7)
	PreTextVSpanTb.Style().SetWidthPx(50)
	PreTextVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	PreTextVSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.PreText.VSpan = intValue(PreTextVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(PreTextVSpanTb)

	p.Add(row1)
}

func buildComposeStyle(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	ComposeCb = gwu.NewCheckBox("合成模式")
	ComposeCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.Switch = ComposeCb.State()
	}, gwu.ETypeClick)
	row.Add(ComposeCb)

	ComposeStyleLb = gwu.NewListBox(common.ComposeStyles)
	ComposeStyleLb.AddEHandlerFunc(func(e gwu.Event) {
		if ComposeStyleLb.SelectedIdx() > 0 {
			index := ComposeStyleLb.SelectedIdx()
			if index <= 3 {
				common.VideoCon.Compose.Style = 3
			} else {
				common.VideoCon.Compose.Style = 1
			}
			common.VideoCon.Compose.Count = common.GetComposeCount(ComposeStyleLb.SelectedIdx())
		} else {
			common.VideoCon.Compose.Switch = false
		}

	}, gwu.ETypeChange)
	row.Add(ComposeStyleLb)

	row.AddHSpace(10)
	content := gwu.NewLabel("背景颜色")
	row.Add(content)
	ComposeBjColorTb = gwu.NewTextBox("")
	ComposeBjColorTb.SetMaxLength(7)
	ComposeBjColorTb.Style().SetWidthPx(50)
	ComposeBjColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeBjColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.BjColor = ComposeBjColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeBjColorTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("左侧间距")
	row.Add(content)
	ComposeLSpanTb = gwu.NewTextBox("")
	ComposeLSpanTb.SetMaxLength(7)
	ComposeLSpanTb.Style().SetWidthPx(50)
	ComposeLSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeLSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.LSpan = intValue(ComposeLSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeLSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("右侧间距")
	row.Add(content)
	ComposeRSpanTb = gwu.NewTextBox("")
	ComposeRSpanTb.SetMaxLength(7)
	ComposeRSpanTb.Style().SetWidthPx(50)
	ComposeRSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeRSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.RSpan = intValue(ComposeRSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeRSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("顶部间距")
	row.Add(content)
	ComposeTSpanTb = gwu.NewTextBox("")
	ComposeTSpanTb.SetMaxLength(7)
	ComposeTSpanTb.Style().SetWidthPx(50)
	ComposeTSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeTSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.TSpan = intValue(ComposeTSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeTSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("底部间距")
	row.Add(content)
	ComposeBSpanTb = gwu.NewTextBox("")
	ComposeBSpanTb.SetMaxLength(7)
	ComposeBSpanTb.Style().SetWidthPx(50)
	ComposeBSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeBSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.BSpan = intValue(ComposeBSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeBSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("水平间距")
	row.Add(content)
	ComposeHSpanTb = gwu.NewTextBox("")
	ComposeHSpanTb.SetMaxLength(7)
	ComposeHSpanTb.Style().SetWidthPx(50)
	ComposeHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeHSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.HSpan = intValue(ComposeHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeHSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("垂直间距")
	row.Add(content)
	ComposeVSpanTb = gwu.NewTextBox("")
	ComposeVSpanTb.SetMaxLength(7)
	ComposeVSpanTb.Style().SetWidthPx(50)
	ComposeVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	ComposeVSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.Compose.VSpan = intValue(ComposeVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ComposeVSpanTb)

	p.Add(row)
}



func buildAfterText(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	AfterTextCb = gwu.NewCheckBox("合成后文字")
	AfterTextCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.Switch = AfterTextCb.State()
	}, gwu.ETypeClick)
	row.Add(AfterTextCb)

	row.AddHSpace(18)
	content := gwu.NewLabel("字体大小:")
	row.Add(content)
	AfterTextSizeTb = gwu.NewTextBox("")
	AfterTextSizeTb.SetMaxLength(500)
	AfterTextSizeTb.Style().SetWidthPx(50)
	AfterTextSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextSizeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.Size = floatValue(AfterTextSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AfterTextSizeTb)

	row.AddHSpace(10)
	AfterTextBoldCb = gwu.NewCheckBox("粗体")
	AfterTextBoldCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.Bold = AfterTextBoldCb.State()
	}, gwu.ETypeClick)
	row.Add(AfterTextBoldCb)

	row.AddHSpace(10)
	content = gwu.NewLabel("字体颜色")
	row.Add(content)
	AfterTextColorTb = gwu.NewTextBox("")
	AfterTextColorTb.SetMaxLength(7)
	AfterTextColorTb.Style().SetWidthPx(50)
	AfterTextColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.TextColor = AfterTextColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AfterTextColorTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("描边颜色")
	row.Add(content)
	AfterTextStrokeColorTb = gwu.NewTextBox("")
	AfterTextStrokeColorTb.SetMaxLength(7)
	AfterTextStrokeColorTb.Style().SetWidthPx(50)
	AfterTextStrokeColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextStrokeColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.StrokeColor = AfterTextStrokeColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AfterTextStrokeColorTb)

	p.Add(row)

	//-------------------//

	row1 := gwu.NewHorizontalPanel()
	row1.AddHSpace(10)
	content = gwu.NewLabel("前缀")
	row1.Add(content)
	AfterTextTb = gwu.NewTextBox("")
	AfterTextTb.SetMaxLength(7)
	AfterTextTb.Style().SetWidthPx(50)
	AfterTextTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.Pre = AfterTextTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(AfterTextTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("后缀")
	row1.Add(content)
	AfterTextSufTb = gwu.NewTextBox("")
	AfterTextSufTb.SetMaxLength(7)
	AfterTextSufTb.Style().SetWidthPx(50)
	AfterTextSufTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextSufTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.Suf = AfterTextSufTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(AfterTextSufTb)

	row1.AddHSpace(10)
	label := gwu.NewLabel("选择字体:")
	row1.Add(label)
	fonts, keys := common.LoadFonts()
	arr := []string{"默认"}
	arr = append(arr, keys...)
	AfterTextFontLb = gwu.NewListBox(arr)
	common.VideoCon.AfterText.FontPath = common.DefaultFont()
	AfterTextFontLb.AddEHandlerFunc(func(e gwu.Event) {
		if AfterTextFontLb.SelectedIdx() > 0 {
			common.VideoCon.AfterText.FontPath = fonts[AfterTextFontLb.SelectedValue()]
		}

	}, gwu.ETypeChange)
	row1.Add(AfterTextFontLb)

	row1.AddHSpace(10)
	AfterTextStyleLb = gwu.NewListBox(common.WaterStyle)
	AfterTextStyleLb.AddEHandlerFunc(func(e gwu.Event) {

		//fmt.Println(WaterTextStyleLb.SelectedIdx(),WaterTextStyleLb.SelectedValue())
		for k, v := range common.WaterStyleMap {
			if v == strings.TrimSpace(AfterTextStyleLb.SelectedValue()) {
				common.VideoCon.AfterText.LayoutStyle = k
			}
		}

	}, gwu.ETypeChange)
	row1.Add(AfterTextStyleLb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("水平间距")
	row1.Add(content)
	AfterTextHSpanTb = gwu.NewTextBox("")
	AfterTextHSpanTb.SetMaxLength(7)
	AfterTextHSpanTb.Style().SetWidthPx(50)
	AfterTextHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextHSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.HSpan = intValue(AfterTextHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(AfterTextHSpanTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("垂直间距")
	row1.Add(content)
	AfterTextVSpanTb = gwu.NewTextBox("")
	AfterTextVSpanTb.SetMaxLength(7)
	AfterTextVSpanTb.Style().SetWidthPx(50)
	AfterTextVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterTextVSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterText.VSpan = intValue(AfterTextVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(AfterTextVSpanTb)

	p.Add(row1)
}




func buildLogo(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	LogoCb = gwu.NewCheckBox("图片logo")
	LogoCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterImage.Switch = LogoCb.State()
	}, gwu.ETypeClick)
	row.Add(LogoCb)

	LogoStyleLb = gwu.NewListBox(common.WaterStyle)
	LogoStyleLb.AddEHandlerFunc(func(e gwu.Event) {
		for k, v := range common.WaterStyleMap {
			if v == strings.TrimSpace(LogoStyleLb.SelectedValue()) {
				common.VideoCon.AfterImage.LayoutStyle = k
			}
		}

	}, gwu.ETypeChange)
	row.Add(LogoStyleLb)

	row.AddHSpace(10)
	content := gwu.NewLabel("水平间距")
	row.Add(content)
	LogoHSpanTb = gwu.NewTextBox("")
	LogoHSpanTb.SetMaxLength(7)
	LogoHSpanTb.Style().SetWidthPx(50)
	LogoHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	LogoHSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterImage.HSpan = intValue(LogoHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(LogoHSpanTb)

	row.AddHSpace(10)
	content = gwu.NewLabel("垂直间距")
	row.Add(content)
	LogoVSpanTb = gwu.NewTextBox("")
	LogoVSpanTb.SetMaxLength(7)
	LogoVSpanTb.Style().SetWidthPx(50)
	LogoVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	LogoVSpanTb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterImage.VSpan = intValue(LogoVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(LogoVSpanTb)

	row.AddHSpace(10)
	label := gwu.NewLabel("选择logo:")
	row.Add(label)
	images, keys := common.LoadImages()
	LogoPathLb = gwu.NewListBox(keys)
	LogoPathLb.AddEHandlerFunc(func(e gwu.Event) {
		if LogoPathLb.SelectedIdx() > 0 {
			common.VideoCon.AfterImage.LogoPath = images[LogoPathLb.SelectedValue()]
		} else {
			common.VideoCon.AfterImage.LogoPath = ""
		}
	}, gwu.ETypeChange)
	row.Add(LogoPathLb)

	p.Add(row)
}


func buildAfterSize(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	AfterSizeCb = gwu.NewCheckBox("合成尺寸")
	AfterSizeCb.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterSize.Switch = AfterSizeCb.State()
	}, gwu.ETypeClick)

	row.Add(AfterSizeCb)

	row.AddHSpace(18)
	content := gwu.NewLabel("宽:")
	row.Add(content)
	AfterSizeW = gwu.NewTextBox("")
	AfterSizeW.SetMaxLength(500)
	AfterSizeW.Style().SetWidthPx(50)
	AfterSizeW.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterSizeW.AddEHandlerFunc(func(e gwu.Event) {

		common.VideoCon.AfterSize.W = intValue(AfterSizeW.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AfterSizeW)

	row.AddHSpace(18)
	content = gwu.NewLabel("高:")
	row.Add(content)
	AfterSizeH = gwu.NewTextBox("")
	AfterSizeH.SetMaxLength(500)
	AfterSizeH.Style().SetWidthPx(50)
	AfterSizeH.AddSyncOnETypes(gwu.ETypeKeyUp)
	AfterSizeH.AddEHandlerFunc(func(e gwu.Event) {
		common.VideoCon.AfterSize.H = intValue(AfterSizeH.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AfterSizeH)
	p.Add(row)

}


func buildSaveBtn(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	bt1 := gwu.NewButton("保存到配置1")
	bt1.AddEHandlerFunc(func(e gwu.Event) {

	}, gwu.ETypeClick)
	row.Add(bt1)

	row.AddHSpace(200)
	bt2 := gwu.NewButton("保存到配置2")
	bt2.AddEHandlerFunc(func(e gwu.Event) {

	}, gwu.ETypeClick)
	row.Add(bt2)

	row.AddHSpace(200)
	bt3 := gwu.NewButton("保存到配置3")
	bt3.AddEHandlerFunc(func(e gwu.Event) {

	}, gwu.ETypeClick)
	row.Add(bt3)
	row.AddVSpace(50)
	p.Add(row)

}

func buildActionBtn(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	row.AddHSpace(300)
	startBtn := gwu.NewButton("开始处理")
	startBtn.Style().SetFontSize("18")
	startBtn.Style().SetColor(gwu.ClrBlue)

	startBtn.AddEHandlerFunc(func(e gwu.Event) {
		go engine.CliEngine.DoEdit("./images", *common.VideoCon)
	}, gwu.ETypeClick)

	row.Add(startBtn)
	//
	//row.AddHSpace(100)
	//clearBtn := gwu.NewButton("清空配置")
	//clearBtn.Style().SetFontSize("18")
	//clearBtn.Style().SetColor(gwu.ClrBlue)
	//clearBtn.AddEHandlerFunc(func(e gwu.Event) {
	//
	//	common.VideoCon = common.InitHVConf()
	//	fillWithConfig(common.VideoCon, e)
	//
	//}, gwu.ETypeClick)
	//
	//row.Add(clearBtn)
	p.Add(row)

}
