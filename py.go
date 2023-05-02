package py

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Option int

const (
	Normal Option = iota
	Tone
	NoTone
	Init
)

func Py(str, sep string, options ...Option) string {
	return strings.Join(Pinyin(str, options...), sep)
}

// 不支持多音字，返回一维字符串数组
func Pinyin(str string, options ...Option) []string {
	runes := []rune(str)
	ret := make([]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = RuneOption(han, style, i, runes)[0]
	}

	return ret
}

// 支持多音字，返回二维字符串数组
func Pinyins(str string, options ...Option) [][]string {
	runes := []rune(str)
	ret := make([][]string, len(runes))
	style := Normal

	if len(options) > 0 {
		style = options[0]
	}

	for i, han := range runes {
		ret[i] = RuneOption(han, style, i, runes)
	}

	return ret
}

func Initials(str string, Heteronym ...bool) string {
	bHeteronym := false
	if len(Heteronym) > 0 {
		bHeteronym = Heteronym[0]
	}
	if bHeteronym {
		return strings.Join(Flatten(Pinyins(str, Init)), "")
	} else {
		return strings.Join(Pinyin(str, Init), "")
	}
}

func ToneIndex(tone string) int {
	// return slices.Index(tones, FromStyle(tone))
	return slices.Index(tones, tone)
}

func Flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}
