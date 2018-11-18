package main

import (
	"github.com/umisan/mercari-watch/lib/db"
	"github.com/umisan/mercari-watch/lib/model/item"
)

func main() {
	db.Init()
	items := []item.Item{
		{Name: "test", Price: 600},
		{Name: "test2", Price: 700},
		{Name: "test3", Price: 800},
	}
	item.AddManyItems(items)
}
