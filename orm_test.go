package ToyORM

import (
	"ToyORM/log"
	"ToyORM/session"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"testing"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	e, err := NewEngine("sqlite3", "Test.db")
	if err != nil {
		t.Fatal("Failed to connect db.", err)
	}
	return e
}

func TestNewEngine(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
}

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (interface{}, error) {
		_ = s.Model(&User{}).CreateTable()
		_, _ = s.Insert(&User{"Tom", 18})
		return nil, errors.New("Error ")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	_, _ = s.Raw("INSERT INTO User(`Name`) values (?),(?)", "Tom", "Sam").Exec()
	err := engine.Migrate(&User{})
	if err != nil {
		log.Error(err)
		return
	}
	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	cols, _ := rows.Columns()
	if !reflect.DeepEqual(cols, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", cols)
	}
}
