package zorm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)


type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	engine, _ := NewEngine("sqlite3", "zorm.db")
	s := engine.NewSession().Model(&User{})
	//log.Warnf("s.RefTable().Tostring():%s",s.RefTable().Tostring())
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}