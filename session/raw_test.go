package session

import (
	"ToyORM/dialect"
	"ToyORM/log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var (
	TestDB         *sql.DB
	TestDialect, _ = dialect.GetDialect("sqlite3")
	err            error
)

func TestMain(m *testing.M) {
	TestDB, err = sql.Open("sqlite3", "../Test.db")
	if err != nil {
		log.Error(err)
	}
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDialect)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`name`) values (?), (?)", "TOM", "SAM").Exec()
	if cnt, err := result.RowsAffected(); err != nil || cnt != 2 {
		t.Fatal("Expected 2, but got ", cnt)
	}
}

func TestSession_QueryRow(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var cnt int
	if err := row.Scan(&cnt); err != nil || cnt != 0 {
		t.Fatal("Failed to query db", err)
	}
}
