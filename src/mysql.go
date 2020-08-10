package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:test@/test")
	if err != nil {
		fmt.Println("failed open", err)

		return
	}

	//insert
	stmt, err := db.Prepare("INSERT userinfo set username=?, departname=?, created=?")
	if err != nil {
		fmt.Println("failed Prepare", err)
		return
	}

	res, err := stmt.Exec("qiyue", "security", "2012-06-26")
	if err != nil {
		fmt.Println("failed exec", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("failed get lastid", err)
		return
	}

	fmt.Println(id)

	//update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	if err != nil {
		fmt.Println("failed prepare update", err)
		return
	}

	res, err = stmt.Exec("sophiehu", id)

	if err != nil {
		fmt.Println("failed exec update", err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println("failed get rows affected")
		return
	}

	fmt.Println(affect)

	//query
	rows, err := db.Query("select * from userinfo")
	if err != nil {
		fmt.Println("failed to prepare select * from userinfo")
		return
	}

	defer rows.Close()
	for rows.Next() {
		var uin int
		var username string
		var department string
		var created string

		err := rows.Scan(&uin, &username, &department, &created)
		if err != nil {
			fmt.Println("failed rows.Scan")
			return
		}

		fmt.Printf("uin:%d, username:%s\n", uin, username)

	}

	db.Close()
}
