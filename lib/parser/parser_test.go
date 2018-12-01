package parser

import (
	"testing"
)

func TestDivideEntryByName(t *testing.T) {
	//テストデータの生成
	test_entry1 := Entry{"aaaa bbbb cccc", "1000"}
	test_entry2 := Entry{"aa　aa bbbb cc　cc", "1000"}
	test_entry3 := Entry{"aaaa　bbbb　cccc", "1000"}
	result1 := divideEntryByName(test_entry1)
	result2 := divideEntryByName(test_entry2)
	result3 := divideEntryByName(test_entry3)
	//test_entry1のチェック
	{
		if len(result1) != 3 {
			t.Errorf("expecting 3, got %d", len(result1))
		}
		if result1[0].Name != "aaaa" || result1[0].Price != "1000" {
			t.Errorf("expecting aaaa, 1000, got %s, %s", result1[0].Name, result1[0].Price)
		}
		if result1[1].Name != "bbbb" || result1[1].Price != "1000" {
			t.Errorf("expecting bbbb, 1000, got %s, %s", result1[1].Name, result1[1].Price)
		}
		if result1[2].Name != "cccc" || result1[2].Price != "1000" {
			t.Errorf("expecting cccc, 1000, got %s, %s", result1[2].Name, result1[2].Price)
		}
	}
	//test_entry2のチェック
	{
		if len(result2) != 5 {
			t.Errorf("expecting 5, got %d", len(result2))
		}
		if result2[0].Name != "aa" || result2[0].Price != "1000" {
			t.Errorf("expecting aa, 1000, got %s, %s", result2[0].Name, result2[0].Price)
		}
		if result2[1].Name != "aa" || result2[1].Price != "1000" {
			t.Errorf("expecting aa, 1000, got %s, %s", result2[1].Name, result2[1].Price)
		}
		if result2[2].Name != "bbbb" || result2[2].Price != "1000" {
			t.Errorf("expecting bbbb, 1000, got %s, %s", result2[2].Name, result2[2].Price)
		}
		if result2[3].Name != "cc" || result2[3].Price != "1000" {
			t.Errorf("expecting cc, 1000, got %s, %s", result2[3].Name, result2[3].Price)
		}
		if result2[4].Name != "cc" || result2[4].Price != "1000" {
			t.Errorf("expecting cc, 1000, got %s, %s", result2[4].Name, result2[4].Price)
		}
	}
	//test_entry3のチェック
	{
		if len(result3) != 3 {
			t.Errorf("expecting 3, got %d", len(result3))
		}
		if result3[0].Name != "aaaa" || result3[0].Price != "1000" {
			t.Errorf("expecting aaaa, 1000, got %s, %s", result3[0].Name, result3[0].Price)
		}
		if result3[1].Name != "bbbb" || result3[1].Price != "1000" {
			t.Errorf("expecting bbbb, 1000, got %s, %s", result3[1].Name, result3[1].Price)
		}
		if result3[2].Name != "cccc" || result3[2].Price != "1000" {
			t.Errorf("expecting cccc, 1000, got %s, %s", result3[2].Name, result3[2].Price)
		}
	}
}
