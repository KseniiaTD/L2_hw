package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagram(strInput []string) map[string][]string {
	mapTemp := make(map[string][]string, len(strInput))
	mapOutput := make(map[string][]string)
	for _, elem := range strInput {
		newElem := strings.Split(strings.ToLower(elem), "")
		sort.Strings(newElem)
		newStr := strings.Join(newElem, "")
		mapTemp[newStr] = append(mapTemp[newStr], strings.ToLower(elem))
	}
	for _, val := range mapTemp {
		mapElem := make(map[string]struct{})
		sliceElem := make([]string, 0, len(val))
		for _, el := range val {
			mapElem[el] = struct{}{}
		}
		for keyMap := range mapElem {
			sliceElem = append(sliceElem, keyMap)
		}
		if len(sliceElem) > 1 {
			sort.Strings(sliceElem)
			mapOutput[val[0]] = sliceElem
		}

	}
	return mapOutput
}

func main() {
	mapRes := Anagram([]string{"мааал", "пятак", "НоГа", "слиток", "пятка", "тяпка", "листок", "ламаа", "столик", "рука", "нога", "столик"})
	fmt.Println(mapRes)

}
