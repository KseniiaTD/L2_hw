package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type StrArr []string

var after int
var before int
var context int
var isCount bool
var isIgnore bool
var isInvert bool
var isFixed bool
var isLineNum bool
var grep string

func Grep(strArr []string) []string {
	var res []string
	lineNum := -1

	var minLine, maxLine int
	for i, str := range strArr {
		if isFixed {
			if isIgnore {
				if strings.EqualFold(str, grep) {
					lineNum, minLine, maxLine = i, i, i+1
					break
				}
			} else {
				if str == grep {
					lineNum, minLine, maxLine = i, i, i+1
					break
				}
			}
		} else {
			if isIgnore {
				if strings.Contains(strings.ToLower(str), strings.ToLower(grep)) {
					lineNum, minLine, maxLine = i, i, i+1
					break
				}
			} else {
				if strings.Contains(str, grep) {
					lineNum, minLine, maxLine = i, i, i+1
					break
				}
			}
		}
	}

	if isCount {
		res = append(res, "количество строк: ", fmt.Sprint(len(strArr)))
	}

	if isLineNum {
		res = append(res, "номер строки: ", fmt.Sprint(lineNum+1))
	}

	if before != 0 || context != 0 {
		if context != 0 {
			minLine = max(0, lineNum-context)
		} else {
			minLine = max(0, lineNum-before)
		}
	}

	if after != 0 || context != 0 {
		if context != 0 {
			maxLine = min(len(strArr), lineNum+context+1)
		} else {
			maxLine = min(len(strArr), lineNum+after+1)
		}
	}

	if isInvert {
		res = append(res, strArr[0:minLine]...)
		if maxLine < len(strArr) {
			res = append(res, strArr[maxLine:]...)
		}
	} else {
		res = append(res, strArr[minLine:maxLine]...)
	}

	return res
}

func main() {

	var filename string
	flag.IntVar(&after, "A", 0, "'after', печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "'before', печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "'context', (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&isCount, "c", false, "'count', количество строк")
	flag.BoolVar(&isIgnore, "i", false, "'ignore-case', игнорировать регистр")
	flag.BoolVar(&isInvert, "v", false, "'invert', вместо совпадения, исключать")
	flag.BoolVar(&isFixed, "F", false, "'fixed', точное совпадение со строкой, не паттерн")
	flag.BoolVar(&isLineNum, "n", false, "'line num', напечатать номер строки")
	flag.StringVar(&filename, "f", "test.txt", "файл с текстом")
	flag.StringVar(&grep, "g", "", "искомый текст")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var strArrIn []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		strArrIn = append(strArrIn, sc.Text())
	}

	strArrOut := Grep(strArrIn)

	filenameOutArr := strings.Split(filename, ".")
	filenameOut := strings.Join([]string{filenameOutArr[0], "_out.", filenameOutArr[1]}, "")

	f, err = os.Create(filenameOut)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i, line := range strArrOut {
		if i > 0 {
			_, err := f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := f.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}
