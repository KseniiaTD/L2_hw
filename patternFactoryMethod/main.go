package main

import "fmt"

func main() {
	excel1, err := getDocument("excel", 300, "loadcards")
	if err != nil {
		fmt.Println(err)
	} else {
		excel1.printDocument()
	}

	pdf1, err := getDocument("pdf", 100, "passport_scan")
	if err != nil {
		fmt.Println(err)
	} else {
		pdf1.printDocument()
	}

	image1, err := getDocument("image", 15, "photo")
	if err != nil {
		fmt.Println(err)
	} else {
		image1.printDocument()
	}

}
