package fenci

import (
	"github.com/go-ego/gse"
	"strings"
)

var (
	Seg gse.Segmenter
)

type Term struct {
	Text  string // 词
	Freq  float64
	Start int // 开始段
	End   int // 结束段
	Pos   string
}

// TermExtract 提取索引词
// 除了标点符号，助词，语气词，形容词，叹词, 副词 其他都被分出来
func TermExtract(str string) []*Term {
	segments := Seg.Segment([]byte(str))
	termList := make([]*Term, 0)
	for _, v := range segments {
		t := v.Token()
		p := t.Pos()
		txt := t.Text()
		end := v.End()
		start := v.Start()
		//logger.Info("txt = ", txt, p)

		if p == "w" || p == "u" || p == "uj" || p == "y" || p == "a" || p == "e" || p == "d" {
			continue
		}

		if p == "x" && !ContainsEnglishAndNumber(txt) {
			continue
		}

		termList = append(termList, &Term{
			Text:  txt,
			Freq:  t.Freq(),
			End:   end,
			Start: start,
			Pos:   p,
		})
	}
	return termList
}

func ContainsEnglishAndNumber(str string) bool {
	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for _, v := range str {
		if strings.Contains(dictionary, string(v)) {
			return true
		}
	}
	return false
}
