package item

import (
	"fmt"
	"log"
	"time"

	"github.com/umisan/mercari-watch/lib/db"
)

type Item struct {
	Id         int
	Name       string
	Price      int
	Created_at time.Time
}

//テーブル内のいくつかの行を取得する
func (item *Item) List() {

}

//テーブル内の特定の行を取得する
func (item *Item) Get() {

}

//テーブルに新しい行を追加する
func (item *Item) Add() {
	dbc := db.GetDB()
	result, err := dbc.Exec("INSERT INTO items (name, price) VALUES ($1, $2)", item.Name, item.Price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

//複数行を一括で追加する
func AddManyItems(items []Item) {
	dbc := db.GetDB()
	temp := make([]db.Model, len(items))
	for i := 0; i < len(items); i++ {
		temp[i] = &items[i]
	}
	vals, sqlStr := db.GenerateSqlForMultiRows(temp)
	sqlStr = "INSERT INTO items (name, price) VALUES " + sqlStr
	stmt, err := dbc.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(vals...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}

//テーブルの特定の行を削除する
func (item *Item) Delete() {

}

//データをスライスにして返す
func (item *Item) To_Slice() []interface{} {
	sliced_data := []interface{}{item.Name, item.Price}
	return sliced_data
}
