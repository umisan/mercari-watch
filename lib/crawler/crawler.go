package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/umisan/mercari-watch/config"
	"github.com/umisan/mercari-watch/lib/model/item"
	"github.com/umisan/mercari-watch/lib/parser"
)

//コスメ香水美容の売り切れ品検索
var target_url string = "https://www.mercari.com/jp/search/?sort_order=&keyword=&category_root=6&category_child=&brand_name=&brand_id=&size_group=&price_min=&price_max=&status_trading_sold_out=1"

func search() string {
	//メルカリにgetを投げる関数
	response, err := http.Get(target_url)
	if err != nil {
		log.Fatal(err)
	}
	response_body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	result := string(response_body[:])
	return result
}

func Start() {
	//クローラーの開始関数
	item_list := []parser.Entry{}      //パース結果を格納するスライス
	item_list_prev := []parser.Entry{} //一つ前のパース結果を格納するスライス
	for {
		result := search()
		item_list_prev = item_list
		item_list = parser.Get_item_list(result)
		//前回との差分を計算
		new_items := []item.Item{}
		if len(item_list_prev) == 0 {
			for _, v := range item_list {
				//fmt.Println(v)
				new_item := item.Item{}
				new_item.Name = v.Name
				new_price, err := strconv.Atoi(v.Price)
				if err != nil {
					log.Fatal(err)
				}
				new_item.Price = new_price
				new_items = append(new_items, new_item)
			}
		} else {
			for _, v := range item_list {
				if v == item_list_prev[0] {
					break
				} else {
					//fmt.Println(v)
					new_item := item.Item{}
					new_item.Name = v.Name
					new_price, err := strconv.Atoi(v.Price)
					if err != nil {
						log.Fatal(err)
					}
					new_item.Price = new_price
					new_items = append(new_items, new_item)
				}
			}
		}
		//fmt.Println()
		item.AddManyItems(new_items)
		time.Sleep(time.Duration(config.CRAWL_DURATION) * time.Minute)
	}
}
