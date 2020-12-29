package main

import "os"
import "net/http"

import "golang.org/x/net/html"

func panil(err error) {
	if err != nil {
		panic(err)
	}
}

func updateBookmarks(title string) {
	home, err := os.UserHomeDir()
	panil(err)

	home += "/.bookmarks"
	f, err := os.OpenFile(home, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	panil(err)
	defer f.Close()

	_, err = f.WriteString(os.Args[1] + "," + title + "\n")
	panil(err)
}

func GetHtmlTitle(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "title" {
		updateBookmarks(n.FirstChild.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		GetHtmlTitle(c)
	}
}

func main() {
	page, err := http.Get(os.Args[1])
	panil(err)
	doc, err := html.Parse(page.Body)
	panil(err)
	GetHtmlTitle(doc)
}
