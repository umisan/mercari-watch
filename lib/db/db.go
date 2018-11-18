package db

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/umisan/mercari-watch/config"
)

var database *sql.DB = nil

func Init() {
	var err error
	database, err = sql.Open("postgres", config.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	if database == nil {
		Init()
	}
	return database
}

func Close() {
	if database != nil {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

//util
//////////////////////////////////////

//Modelインターフェース
type Model interface {
	To_Slice() []interface{}
}

//複数行を一括で追加するためのSQL生成関数
//引数:追加するデータの配列(二重配列)
//戻り値:SQL用の引数とSQL文
func GenerateSqlForMultiRows(datas []Model) ([]interface{}, string) {
	vals := []interface{}{}
	sqlStr := ""
	count := 1
	for _, data := range datas {
		sliced_data := data.To_Slice()
		temp := "("
		length := len(sliced_data)
		for i := 0; i < length-1; i++ {
			temp += "$" + strconv.Itoa(count) + ","
			vals = append(vals, sliced_data[i])
			count++
		}
		temp += "$" + strconv.Itoa(count) + "),"
		vals = append(vals, sliced_data[length-1])
		count++
		sqlStr += temp
	}
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	return vals, sqlStr
}
