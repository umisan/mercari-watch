package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/umisan/mercari-watch/config"
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
	for {
		result := search()
		fmt.Println(result)
		time.Sleep(config.CRAWL_DURATION * time.Minute)
	}
}
