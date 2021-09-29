package session

import (
	"ToyORM/log"
	"testing"
)

type Account struct {
	ID       int `orm:"PRIMARY KEY"`
	Password string
}

func (a *Account) BeforeInsert() error {
	log.Info("Before insert", a)
	a.ID += 1000
	return nil
}

func (a *Account) AfterQuery() error {
	log.Info("After query", a)
	a.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})
	u := &Account{}
	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks")
	}
}
