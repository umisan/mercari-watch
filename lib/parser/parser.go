package parser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Entry struct {
	Name, Price string
}

//円マークと,を取り除いてintに変換
func stringToPrice(input string) string {
	input = strings.Replace(input, ",", "", -1)
	input = strings.Replace(input, "¥ ", "", -1)
	return input
}

//<section class="items-box">を検索
//section内の
//<div class="items-box-body">
//<h3 class="items-box-name font-2">ナイトアイボーテ</h3>が名前
//<div class="items-box-price font-5">¥ 3,700</div>が値段
func dfs(n *html.Node, list *[]Entry) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-body" {
				//探索対象
				var entry Entry
				get_name_price(n, &entry)
				*list = append(*list, entry)
			}
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, list)
	}
}

//items-box-body内の値段と名前をDFSでサーチ
func get_name_price(n *html.Node, entry *Entry) {
	if n.Type == html.ElementNode && n.Data == "h3" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-name font-2" {
				entry.Name = n.FirstChild.Data
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Val == "items-box-price font-5" {
				entry.Price = stringToPrice(n.FirstChild.Data)
			}
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		get_name_price(child, entry)
	}
}

func Get_item_list(html_string string) []Entry {
	html_reader := strings.NewReader(html_string)
	doc, err := html.Parse(html_reader)
	if err != nil {
		log.Fatal(err)
	}
	list := []Entry{}
	dfs(doc, &list)
	return list
}
