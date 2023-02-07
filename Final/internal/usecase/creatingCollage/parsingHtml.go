package creatingCollage

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Page структура web-страницы
type Page struct {
	URL       string
	ImagesUrl []string
	Links     []string
}

// pageCache кеш web-страниц
var pageCache = sync.Map{}

// ChanPics канал для передачи url-ссылок изображений
var ChanPics = make(chan string, 100_000)

// ChanLinks канал для передачи ссылок web-страниц, данный канал используется
// как очередь для корректного обхода в ширину
var ChanLinks = make(chan string, 100_000)

// getLinksAndImages возвращает все ссылки и url-адреса изображений со страницы
func getLinksAndImages(htmlDoc, host string) (links, images []string) {

	doc, err := html.Parse(strings.NewReader(htmlDoc))
	if err != nil {
		fmt.Println(err)
		return links, images
	}

	var extractLinkAndImage func(*html.Node)
	extractLinkAndImage = func(node *html.Node) {

		if node.Type == html.ElementNode && (node.Data == "a" || node.Data == "img") {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
				if attr.Key == "src" {
					images = append(images, host+attr.Val)
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			extractLinkAndImage(c)
		}
	}
	extractLinkAndImage(doc)

	return links, images
}

// processPage обрабатывает страницу, отправляя ссылки и url-адреса изображений
// в очередь и канал соответственно для последующей обработки
func processPage(page Page, host string) {
	if val, ok := pageCache.Load(page.URL); !ok {
		resp, err := http.Get(page.URL)
		if err != nil {
			log.Fatalf("URL page %s does not exist", page.URL)
		}
		defer func() {
			if err = resp.Body.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		page.Links, page.ImagesUrl = getLinksAndImages(string(body), host)
		pageCache.Store(page.URL, page)
	} else {
		page = val.(Page)
	}

	for _, imageUrl := range page.ImagesUrl {
		ChanPics <- imageUrl
	}

	for _, link := range page.Links {
		ChanLinks <- link
	}
}

// ProcessingAllPages считывает ссылки из очереди и запускает функцию processPage для каждой ссылки
func ProcessingAllPages(host string) {
	startPage := Page{URL: host}
	processPage(startPage, host)

	go func() {
		time.Sleep(time.Second)
		for {
			if len(ChanLinks) < 10 {
				close(ChanLinks)
				return
			}
		}
	}()

	for link := range ChanLinks {
		page := Page{URL: host + link}
		processPage(page, host)
	}
}
