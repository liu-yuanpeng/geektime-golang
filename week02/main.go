package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

)
type User struct {
	id int16 `db:"id"`
	name string`db:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := querySQL(db); err != nil {
		fmt.Printf("fatal:%+v\n", err)
	}
}

func querySQL(DB *sql.DB) error {
	user := User{}
	rows, err := DB.Query("select id, name from users where ids = ?", 1)
	if err != nil {
		//需要使用wrap进行包装,这样当调用者执调用出错时，既可以添加自定义错误信息，又会保留真正错误的堆栈信息
		return errors.Wrap(err, "main: querySQLErr")
	}
	if rows.Next(){
		rows.Scan(&user.id,&user.name)
		print(user.id,user.name)
	}
	return nil
}