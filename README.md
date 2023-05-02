# go-py

[![GoCI](https://github.com/xuender/go-py/workflows/Go/badge.svg)](https://github.com/xuender/go-py/actions)
[![codecov](https://codecov.io/gh/xuender/go-py/branch/main/graph/badge.svg?token=1QQNBH82CM)](https://codecov.io/gh/xuender/go-py)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-py)](https://goreportcard.com/report/github.com/xuender/go-py)
[![GoDoc](https://godoc.org/github.com/xuender/go-py?status.svg)](https://pkg.go.dev/github.com/xuender/go-py)
[![Gitter](https://badges.gitter.im/xuender-go-py/community.svg)](https://gitter.im/xuender-go-py/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GitHub license](https://img.shields.io/github/license/xuender/go-py)](https://github.com/xuender/go-py/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/go-py)](https://github.com/xuender/go-py/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/go-py)](https://github.com/xuender/go-py/stargazers)

中文转拼音.

## use

```shell
go get github.com/xuender/go-py
```

```go
package main

import (
	"fmt"
	py "github.com/xuender/go-py"
)

func main() {
	hans := "曾某曾是异乡人！"
	fmt.Println(py.Pinyin(hans))
	fmt.Println(py.Pinyin(hans, py.Tone))
	fmt.Println(py.Pinyins(hans, py.Tone))
	fmt.Println(py.Pinyin(hans, py.NoTone))
	fmt.Println(py.Pinyins(hans, py.NoTone))
	fmt.Println(py.Pinyin(hans, py.Init))
	fmt.Println(py.Pinyins(hans, py.Init))
	fmt.Println(py.Initials(hans))
	fmt.Println(py.Initials(hans, true))
	// output:
	// [ce2ng mo3u ce2ng shi4 yi4 xia1ng re2n ！]
	// [céng mǒu céng shì yì xiāng rén ！]
	// [[céng zēng] [mǒu méi] [céng zēng] [shì tí] [yì yí] [xiāng] [rén] [！]]
	// [ceng mou ceng shi yi xiang ren ！]
	// [[ceng zeng] [mou mei] [ceng zeng] [shi ti] [yi] [xiang] [ren] [！]]
	// [c m c s y x r ！]
	// [[c z] [m] [c z] [s t] [y] [x] [r] [！]]
	// cmcsyxr！
	// czmczstyxr！
}
```

## cmd

```shell
py --help
```

### install

```shell
go install github.com/xuender/go-py/cmd/py@latest
```

### examples

```shell
py 阿弥陀佛
py -i -s= 阿弥陀佛
py -t -h -hs=\; 阿弥陀佛
py -n 阿弥陀佛
```
