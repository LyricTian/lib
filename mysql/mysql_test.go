package mysql

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Test struct {
	ID   int    `comment:"id,自增ID"`
	Code string `comment:"code,编号"`
	Name string `comment:"name,名称"`
}

func (t *Test) TableName() string {
	return "test"
}

func (t *Test) TableComment() string {
	return "测试表"
}

func TestAlterComment(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/myapp_test?charset=utf8")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	cdb := NewDB(db, "myapp_test", log.New(os.Stdout, "[db]", log.LstdFlags))
	err = cdb.AlterComment(new(Test))
	if err != nil {
		t.Error(err)
	}
}
