package engine

import (
	"fmt"
	cm "myProject/videoCli/common"
	"myProject/videoCli/imageCli"
	"myTool/file"
	"os"
)

var CliEngine *imageCli.Engine

var isDealing bool

func init() {
	f := func(msg string) {
		fmt.Println(msg)
	}
	CliEngine = imageCli.NewEngine("./source/tempImages", "./output", f)
}

func DoEdit(con cm.ImageConf) {
	if isDealing {
		fmt.Println("正在处理中...")
		return
	}

	if file.PathExist(CliEngine.OutPutDir) == false {
		os.MkdirAll(CliEngine.OutPutDir, os.ModePerm)
	}
	isDealing = true
	CliEngine.DoEdit("./images", con)
	isDealing = false
}
