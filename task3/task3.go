package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Row []string

var clmnSort int
var isReverse bool
var isNotDouble bool
var isNumeric bool

type Tab struct {
	rows []Row
}

func (t Tab) Len() int {
	return len(t.rows)
}

func (t Tab) Swap(i int, j int) {
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

func (t Tab) Less(i int, j int) bool {
	if len(t.rows[i]) <= clmnSort || len(t.rows[j]) <= clmnSort {
		for l := 0; l < min(len(t.rows[i]), len(t.rows[j])); l++ {
			if t.rows[i][l] != t.rows[j][l] {
				return check(t.rows[i][l], t.rows[j][l])
			}
		}
		return len(t.rows[i]) < len(t.rows[j])
	} else {

		if t.rows[i][clmnSort] != t.rows[j][clmnSort] {
			return check(t.rows[i][clmnSort], t.rows[j][clmnSort])
		} else {
			for l := 0; l < min(len(t.rows[i]), len(t.rows[j])); l++ {
				if l != clmnSort {
					if t.rows[i][l] != t.rows[j][l] {
						return check(t.rows[i][l], t.rows[j][l])
					}

				}
			}
		}
		return len(t.rows[i]) < len(t.rows[j])
	}
}

func check(el1 string, el2 string) bool {
	if !isNumeric {
		return el1 < el2
	} else {
		x, errX := strconv.ParseInt(el1, 10, 64)
		y, errY := strconv.ParseInt(el2, 10, 64)

		if errX != nil && errY != nil {
			return el1 < el2
		} else if errX != nil {
			return true
		} else if errY != nil {
			return false
		}

		return x < y
	}
}

func main() {

	var filename string
	flag.IntVar(&clmnSort, "k", 0, "number of column for sort")
	flag.BoolVar(&isReverse, "r", false, "is reverse sort")
	flag.BoolVar(&isNotDouble, "u", false, "sort with unique rows")
	flag.BoolVar(&isNumeric, "n", false, "compare according to string numerical value")
	flag.StringVar(&filename, "f", "test.txt", "filename for read")
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tab := Tab{}
	m := make(map[string]struct{})
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		_, found := m[sc.Text()]
		if !(isNotDouble && found) {
			tab.rows = append(tab.rows, strings.Split(sc.Text(), " "))
			m[sc.Text()] = struct{}{}
		}
	}

	if isReverse {
		sort.Sort(sort.Reverse(tab))
	} else {
		sort.Sort(tab)
	}

	filenameOutArr := strings.Split(filename, ".")
	filenameOut := strings.Join([]string{filenameOutArr[0], "_out.", filenameOutArr[1]}, "")

	f, err = os.Create(filenameOut)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i, line := range tab.rows {
		if i > 0 {
			_, err := f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := f.WriteString(strings.Join(line[:], " "))
		if err != nil {
			log.Fatal(err)
		}
	}
}
