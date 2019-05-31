package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// testGin()
	testMysql()
}

func testMysql() {
	dsn := "root:!ZhangLei123@tcp(cdb-6mzei5eu.bj.tencentcdb.com:10011)/zj"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db connect failed!", dsn, err)
		return
	}
	rows, err := db.Query("show tables")
	if err != nil {
		fmt.Println("db query error", err)
		return
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(columns)
	count := len(columns)

	row := make([]interface{}, count)
	for rows.Next() {
		err := rows.Scan(&row)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(row)
	}
}

func testGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
