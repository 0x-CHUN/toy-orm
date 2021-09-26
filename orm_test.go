package ToyORM

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestNewEngine(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	e, err := NewEngine("sqlite3", "Test.db")
	if err != nil {
		t.Fatal("Failed to connect db.", err)
	}
	return e
}
