package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func TestMySQLAPI() {
	// ...
	dsn := "root:Aa12345600@tcp(172.30.189.56:3306)/demo"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

	type user struct {
		id       int
		name     string
		password string
	}

	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.password)
		if err != nil {
			fmt.Printf("scan data failed, err:%v\n", err)
			return
		}
		fmt.Printf("user name:%v, password:%v\n", u.name, u.password)
	}
	insertRowDemo(db)
	updateRowDemo(db)
	deleteRowDemo(db)
}

// 增加一行数据
func insertRowDemo(db *sql.DB) {
	sqlStr := "INSERT INTO user(name, password) VALUES(?, ?)"
	result, err := db.Exec(sqlStr, "小羽", 1235555)
	if err != nil {
		fmt.Printf("insert data failed, err:%v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert lastInsertId failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, id:%d\n", id)

}

// 更新一组数据
func updateRowDemo(db *sql.DB) {
	sqlStr := "UPDATE user SET password = ? WHERE id = ?"
	result, err := db.Exec(sqlStr, 9999999, 3)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除一行数据
func deleteRowDemo(db *sql.DB) {
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete data failed, err:%d\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
