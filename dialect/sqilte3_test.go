package dialect

import (
	"reflect"
	"testing"
)

func TestDataTypeof(t *testing.T) {
	dial := &Sqlite3{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Tom", "text"},
		{123, "integer"},
		{1.2, "real"},
		{[]int{1, 2, 3}, "blob"},
	}
	for _, c := range cases {
		if tp := dial.DataTypeof(reflect.ValueOf(c.Value)); tp != c.Type {
			t.Fatalf("Expect %s, but got %s", c.Type, tp)
		}
	}
}
