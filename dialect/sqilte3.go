package dialect

import (
	"ToyORM/log"
	"fmt"
	"reflect"
	"time"
)

type Sqlite3 struct{}

var _ Dialect = (*Sqlite3)(nil)

func init() {
	RegisterDialect("sqlite3", &Sqlite3{})
	log.Info("Sqlite3 dialect register ")
}

func (s Sqlite3) DataTypeof(tp reflect.Value) string {
	switch tp.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := tp.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", tp.Type().Name(), tp.Kind()))
}

func (s Sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
