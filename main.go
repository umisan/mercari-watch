package main

import (
	"io/ioutil"

	"github.com/umisan/mercari-watch/lib/parser"
)

func main() {
	html_string, _ := ioutil.ReadFile("/home/umino/go/src/github.com/umisan/mercari-watch/lib/test.html")
	parser.Get_item_list(string(html_string))
}
