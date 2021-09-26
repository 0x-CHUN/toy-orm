package dialect

import "reflect"

type Dialect interface {
	DataTypeof(tp reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

var dialectsMAP = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMAP[name] = dialect
}

func GetDialect(name string) (Dialect, bool) {
	dialect, ok := dialectsMAP[name]
	return dialect, ok
}
