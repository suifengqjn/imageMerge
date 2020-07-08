package common

import (
	"myTool/fmg"
	"path/filepath"
	"strconv"
	"strings"
)

func FloatValue(v string) float64  {
	value, err := strconv.ParseFloat(v, 10)
	if err != nil {
		return 0
	}

	return value
}

func ExtractVideos(files []string) []string  {

	var res []string
	for _, f := range files {
		baseName := filepath.Base(f)

		suf := strings.Split(baseName, ".")[1]
		if Contains(fmg.VideoSufs(), strings.ToLower(suf)) {
			res = append(res, f)
		}
	}
	return res
}