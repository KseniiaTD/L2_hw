package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

var urlMy string
var pathName string

func download(bytes []byte, filename string) error {
	err := os.MkdirAll(pathName+filepath.Dir(filename), 0755)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(pathName+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)

	if err != nil {
		return err
	}

	return err
}

func main() {
	flag.StringVar(&urlMy, "u", "", "'url' - ссылка на html-страницу")
	flag.Parse()

	if len(urlMy) == 0 {
		log.Fatal("url is empty")
	}

	urlPath, err := url.Parse(urlMy)
	if err != nil {
		log.Fatal(err)
	}

	pathName, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathName = pathName + "/" + urlPath.Host
	err = os.MkdirAll(pathName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pathName)

	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		link, err := url.Parse(href)
		if err != nil {
			log.Fatal(err)
		}
		host := link.Hostname()
		if (host == urlPath.Host || len(host) == 0) && (len(link.Scheme) == 0 || link.Scheme == "http" || link.Scheme == "https") { //!strings.HasPrefix(href, "mailto") {
			e.Request.Visit(href)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		path := r.Request.URL.Path
		download(r.Body, path)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(urlMy)
}
