package db

import (
	"database/sql"
	"fmt"

	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

//MysqlConn 数据库连接
func MysqlConn() {
	db, err := sql.Open("mysql", "root:126com@/pandora_package_manager?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT tbuser SET sUserName=?,sPhoneNum=?,sEmail=?,sPassword=?,tCreateTime=?")
	checkErr(err)

	res, err := stmt.Exec("haipi", "18128866737", "hlcaca@126.com", "126com", "2019-02-24 21:51:36")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("uid=", id)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//封装数据库功能

/*
id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("uid=", id)
	//更新数据
	stmt, err = db.Prepare("update tbuser set sUserName=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("haipiUpdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect num =", affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM tbuser")
	checkErr(err)

	for rows.Next() {
		var uid int
		var sUserName string
		var sPhoneNum string
		var sEmail string
		err = rows.Scan(&uid, &sUserName, &sPhoneNum, &sEmail)
		checkErr(err)
		fmt.Println("记录", uid, ":")
		fmt.Println(sUserName)
		fmt.Println(sPhoneNum)
		fmt.Println(sEmail)
	}

	//删除数据
	stmt, err = db.Prepare("delete from tbuser where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("删除数量：", affect)

*/
