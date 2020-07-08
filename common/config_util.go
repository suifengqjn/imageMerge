package common

import (
	"encoding/json"
	"io/ioutil"
	cm "myProject/videoCli/common"
	"myTool/ffmpeg"
	"myTool/file"
	"myTool/sys"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	config1Path = "./source/files/config1.json"
	config2Path = "./source/files/config2.json"
	config3Path = "./source/files/config3.json"

	imagePath = "./material"
	filmPath  = "./material"
	fontPath  = "./source/font"
)

var (
	imageSufs = []string{"jpg", "jpeg", "png"}
	fontSufs  = []string{"ttc", "ttf"}
	system    = sys.GetSysInfo()
)

var (
	TextColors    = []string{"未选择", "red", "yellow", "green", "black", "white", "blue", "gray", "purple"}
	TextColorsMap = map[string]int{
		"red":    1,
		"yellow": 2,
		"green":  3,
		"black":  4,
		"white":  5,
		"blue":   6,
		"gray":   7,
		"purple": 8,
	}

	ComposeStyles    = []string{"未选择", "2x2", "2x3", "2x4", "1x2", "1x3", "1xn"}
	ComposeStylesMap = map[int]string{
		0: "未选择",
		1: "2x2",
		2: "2x3",
		3: "2x4",
		4: "1x2",
		5: "1x3",
		6: "1xn",
	}

	WaterStyle    = []string{"未选择", "左上角", "右上角", "右下角", "左下角", "正中间", "顶部居中", "右侧居中", "底部居中", "左侧居中"}
	WaterStyleMap = map[int]string{
		1: "左上角",
		2: "右上角",
		3: "右下角",
		4: "左下角",
		5: "正中间",
		6: "顶部居中",
		7: "右侧居中",
		8: "底部居中",
		9: "左侧居中"}

	/*
		        1:"左下角",
				2:"底部居中",
				3:"右下角",
				5:"左上角",
				6:"顶部居中",
				7:"右上角",
				9:"左侧居中",
				10:"正中间",
				11:"右侧居中",
	*/
	StrokeStyle    = []string{"未选择", "左下角", "底部居中", "右下角", "左上角", "顶部居中", "右上角", "左侧居中", "正中间", "右侧居中"}
	StrokeStyleMap = map[int]string{
		1: "左下角",
		2: "底部居中",
		3: "右下角",
		4: "左上角",
		5: "顶部居中",
		6: "右上角",
		7: "左侧居中",
		8: "正中间",
		9: "右侧居中",
	}

	PinPStyle    = []string{"未选择", "左上角", "右上角", "右下角", "左下角"}
	PinPStyleMap = map[int]string{
		1: "左上角",
		2: "右上角",
		3: "右下角",
		4: "左下角"}

	BjImageStyle    = []string{"未选择", "自动", "左右", "上下"}
	BjImageStyleMap = map[int]string{
		1: "自动",
		2: "左右",
		3: "上下",
	}
)

func LoadConfig1() *cm.ImageConf {
	buf, err := ioutil.ReadFile(config1Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &VideoCon)
	if err != nil {
		return nil
	}

	return VideoCon

}

func LoadConfig2() *cm.ImageConf {
	buf, err := ioutil.ReadFile(config2Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &VideoCon)
	if err != nil {
		return nil
	}

	return VideoCon

}

func LoadConfig3() *cm.ImageConf {
	buf, err := ioutil.ReadFile(config3Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &VideoCon)
	if err != nil {
		return nil
	}

	return VideoCon

}

func SaveConfig1(con *cm.ImageConf) {

	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config1Path, buf)
	}

}

func SaveConfig2(con *cm.ImageConf) {
	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config2Path, buf)
	}
}

func SaveConfig3(con *cm.ImageConf) {
	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config3Path, buf)
	}
}

func writeToFile(filepath string, buf []byte) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err == nil {
		file.Write(buf)
	}
}

func LoadFonts() (map[string]string, []string) {
	var files []string
	var err error
	if system.PlatForm == sys.MacOS {
		files, err = file.GetAllFiles(fontPath)
	} else {
		files, err = file.GetAllFiles("./source/font/")
	}
	if err != nil {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		f = formatPath(f)
		baseName := file.GetFileBaseName(f)
		suf := file.GetFileSuf(f)
		if Contains(fontSufs, strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	return res, keys
}

// fileName : filePath
func LoadImages() (map[string]string, []string) {
	files, err := file.GetAllFiles(imagePath)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		f = formatPath(f)
		baseName := filepath.Base(f)
		suf := strings.Split(baseName, ".")[1]
		if Contains(imageSufs, strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	sort.Strings(keys)
	return res, keys
}

func LoadFilms() (map[string]string, []string) {
	files, err := file.GetAllFiles(filmPath)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		f = formatPath(f)
		baseName := filepath.Base(f)
		suf := strings.Split(baseName, ".")[1]
		if Contains(ffmpeg.VideoSufs(), strings.ToLower(suf)) || Contains(ffmpeg.ImageSufs(), strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	sort.Strings(keys)
	return res, keys
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func formatPath(path string) string {
	if system.PlatForm >= sys.Win32 {
		return strings.ReplaceAll(path, `\`, `/`)
	} else {
		return path
	}
}

func DefaultFont() string {
	return "./source/simsun.ttc"
}

func GetComposeCount(style int) int {
	switch style {
	case 1:
		return 4
	case 2:
		return 6
	case 3:
		return 8
	case 4:
		return 2
	case 5:
		return 3
	case 6:
		return -1
	default:
		return 1
	}
	return 1
}
