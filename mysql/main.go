package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	name     string
	password string
}

var db *sql.DB

func initDB() (err error) {
	// mysql 驱动是否存在
	dsn := "root:@tcp(127.0.0.1:3306)/mylaravel?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func insertRowDemo() {
	sqlstr := "INSERT INTO users(`name`,`email`,`password`,`regbyemail`) VALUES(?,?,?,?)"
	prefix := "robot"
	uxt := time.Now().Unix()
	sufix := fmt.Sprintf("%d", uxt)
	name := prefix + sufix
	email := prefix + sufix + "@gmail.com"
	password := prefix + sufix
	regbyemail := email
	res, err := db.Exec(sqlstr, name, email, password, regbyemail)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := res.LastInsertId()
	fmt.Println(id)
}
func updateRowDemo() {
	sqlstr := "UPDATE users SET remember_token = ? WHERE id= ?"
	uxt := time.Now().Unix()
	struxt := strconv.FormatInt(uxt, 10)
	remember_token := fmt.Sprintf("%x", md5.Sum([]byte(struxt)))
	id := 2
	stmt, err := db.Prepare(sqlstr)
	defer stmt.Close()
	res, err := stmt.Exec(remember_token, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	affectRows, err := res.RowsAffected()
	fmt.Println(affectRows)

}
func queryRowDemo() {
	id := 2
	sqlstr := "select id, name, password from users where id=?"
	var u user
	// QueryRow执行一次查询，并期望返回最多一行结果（即Row）。
	// QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
	err := db.QueryRow(sqlstr, id).Scan(&u.id, &u.name, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println(u.id, u.name, u.password)
}

func queryMulti1Demo() {
	sqlstr := "select id, name, password from users"
	var u user
	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&u.id, &u.name, &u.password); err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.id, u.name, u.password)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
func queryMulti2Demo() {
	sqlstr := "SELECT * FROM users"
	// 没有 where 则没有 db.Query 的第二个字段 args
	// Query执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}

	// 获取所有字段的名称，返回切片
	colName, _ := rows.Columns()

	userCol := make([][]byte, len(colName))
	scans := make([]interface{}, len(colName))

	// func (r *Row) Scan(dest ...interface{}) error 后续 Scan 需要使用 interface{} 切片。
	// 先把地址存入 []scans 切片中，
	for k := range colName {
		scans[k] = &userCol[k] //数据转移到scans，scans[0]="id",scans[0]
	}

	// i := 0
	// users := make(map[int]map[string]string) //定义一个map用于存储所有查询结果
	users := []map[string]string{}
	for rows.Next() {
		// Scan将 该行查询结果 各列 分别保存 进dest参数指定的值中。
		// 如果该查询匹配多行，Scan会使用 第一行结果 并 丢弃其余各行。如果没有匹配查询的行，Scan会返回ErrNoRows。
		rows.Scan(scans...) //使用scans切片接收填充数据

		/*
			// scan 以后，已经把数据保存到 userCol 的切片中
			for _, v := range userCol {
				fmt.Println(string(v))
			}
			os.Exit(0)
		*/

		theRow := make(map[string]string) //接收当前Scan到的1行数据到map切片
		for k, v := range userCol {
			theRow[colName[k]] = string(v) //这里把[]byte数据转成string
		}
		// users[i] = theRow //把格式化后的theRow放入users结果集
		// i++
		users = append(users, theRow)
	}
	for _, v := range users {
		fmt.Printf("ID:%v 名字:%v email:%v remember_token:%v \n", v["id"], v["name"], v["email"], v["remember_token"])
	}

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed")
	}
	// queryRowDemo()
	// queryMulti1Demo()
	// insertRowDemo()
	// updateRowDemo()
	queryMulti2Demo()

}
