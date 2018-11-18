package db

import (
	"testing"
)

type test struct {
	name  string
	price int
}

func (t *test) To_Slice() []interface{} {
	return []interface{}{t.name, t.price}
}

func TestGenerateSqlForMultiRows(t *testing.T) {
	datas := []test{
		{"test1", 600},
		{"test2", 700},
		{"test3", 800},
	}
	exp_vals := []interface{}{"test1", 600, "test2", 700, "test3", 800}
	exp_sqlStr := "($1,$2),($3,$4),($5,$6)"
	temp := make([]Model, len(datas))
	for i := 0; i < len(datas); i++ {
		temp[i] = &datas[i]
	}
	vals, sqlStr := GenerateSqlForMultiRows(temp)
	//valsのチェック
	if len(vals) == 0 {
		t.Errorf("expecting 6, got %d", len(vals))
	}
	for i := 0; i < len(vals); i++ {
		if vals[i] != exp_vals[i] {
			t.Errorf("expecting %v, got %v", exp_vals[i], vals[i])
		}
	}
	//sqlStrのチェック
	if sqlStr != exp_sqlStr {
		t.Errorf("expecting %s, got %s", exp_sqlStr, sqlStr)
	}
}
