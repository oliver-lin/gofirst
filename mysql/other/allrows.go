package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//先把字段的值都当成字符串
func Rows2SliceMap(rows *sql.Rows) (list []map[string]string) {
	//字段名称
	columns, _ := rows.Columns()
	//多少个字段
	length := len(columns)
	//每一行字段的值
	values := make([]sql.RawBytes, length)
	//保存的是values的内存地址
	pointer := make([]interface{}, length)
	//
	for i := 0; i < length; i++ {
		pointer[i] = &values[i]
	}
	//
	for rows.Next() {
		//把参数展开，把每一行的值存到指定的内存地址去，循环覆盖，values也就跟着被赋值了
		rows.Scan(pointer...)
		//每一行
		row := make(map[string]string)
		for i := 0; i < length; i++ {
			row[columns[i]] = string(values[i])
		}
		list = append(list, row)
	}
	//
	return
}

func main() {
	// db, err := sql.Open("mysql", "root:1234@tcp(192.168.99.165:3306)/test?charset=utf8")
	// if err != nil {
	//     log.Println(err.Error())
	// }
	// err = db.Ping()
	// if err != nil {
	//     log.Println(err.Error())
	// }
	// rows, _ := db.Query("select * from cms")
	// defer rows.Close()
	// list := Rows2SliceMap(rows)
	// for k, v := range list {
	//     fmt.Println(k)
	//     fmt.Println(v["id"], v["name"], v["sex"], v["age"])
	// }
}
