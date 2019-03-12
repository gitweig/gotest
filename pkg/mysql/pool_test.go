package mysql

import (
	"testing"
)

//User
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Score int    `db:"score"`
}

//Test_Insert
func Test_Insert(t *testing.T) {
	pool, _ := Pool("127.0.0.1",
		3306,
		"root",
		"123456",
		"test",
		"utf8mb4",
		10,
		5)
	if pool == nil {
		t.Skip("Init mysql error!")
		return
	}

	res, err := pool.Exec("Insert into User(id, name, score) values(?,?,?)", 1, "jack", 99)
	if err != nil {
		t.Errorf("insert error %s", err.Error())
	} else {
		lastid, err := res.LastInsertId()
		if err != nil {
			t.Errorf("get last id err %s ", err.Error())
		} else {
			t.Logf("insert row last id %d", lastid)
		}

	}
}

//Test_Select
func Test_Select(t *testing.T) {
	pool, _ := Pool("127.0.0.1",
		3306,
		"root",
		"123456",
		"test",
		"utf8mb4",
		10,
		5)
	if pool == nil {
		t.Skip("Init mysql error!")
		return
	}

	var users []User
	err := pool.Select(&users,
		"SELECT id, name, score from user limit ?",
		10)
	if err != nil {
		t.Errorf("select fail %s", err.Error())
	} else {
		t.Logf("users length %d", len(users))
	}
}
