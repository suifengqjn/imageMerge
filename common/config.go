package common

import (
	"fmt"
	"myProject/videoCli/common"
	"myTool/imageSplice"
	"os"
	"time"
)

var VideoCon *common.ImageConf
var ProjectPath string
func init() {

	var err error
	ProjectPath, err = os.Getwd()
	if err != nil {
		fmt.Println("获取文件路径失败", err)
		time.Sleep(time.Minute)
		panic(err)
	}
	VideoCon = InitImageConf()
}

func InitImageConf()*common.ImageConf  {
	return &common.ImageConf{
		PreSize:    imageSplice.Size{},
		PreText:    imageSplice.WaterTextConfig{F:singleText, Dpi:100},
		Compose:    imageSplice.ComposeParam{},
		AfterText:  imageSplice.WaterTextConfig{F:composeText,Dpi:100},
		AfterImage: imageSplice.WaterImageConfig{},
		AfterSize:  imageSplice.Size{},
	}
}

func singleText(fileName string, pre string, suf string) string {
	return pre + fileName + suf
}

func composeText(fileName string, pre string, suf string) string {
	return pre  + suf
}