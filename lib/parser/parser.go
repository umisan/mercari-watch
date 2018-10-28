package parser

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

//<section class="items-box">を検索
//section内の
//<div class="items-box-body">
//<h3 class="items-box-name font-2">ナイトアイボーテ</h3>が名前
//<div class="items-box-price font-5">¥ 3,700</div>が値段
func dfs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-body" {
				//探索対象
				var name, price string
				get_name_price(n, &name, &price)
				fmt.Println(name, price)
			}
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		dfs(child)
	}
}

//items-box-body内の値段と名前をDFSでサーチ
func get_name_price(n *html.Node, name *string, price *string) {
	if n.Type == html.ElementNode && n.Data == "h3" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-name font-2" {
				*name = n.FirstChild.Data
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-price font-5" {
				*price = n.FirstChild.Data
			}
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		get_name_price(child, name, price)
	}
}

func Get_item_list(html_string string) {
	html_reader := strings.NewReader(html_string)
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatal(err)
	}
	dfs(doc)
}
