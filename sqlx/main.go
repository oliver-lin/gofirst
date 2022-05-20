package main

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"title"`
}

var Db *sqlx.DB

func init() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/mylaravel"
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func (post *Post) GetPost(id int) (*Post, error) {
	post = &Post{}
	sql := "select id,content,title from posts where id= ?"
	err := Db.QueryRowx(sql, id).StructScan(post)
	if err != nil {
		return nil, nil
	}
	return nil, nil

}

func (post *Post) Create() error {
	sql := "insert into posts(content,title) values (?,?)"
	err := Db.QueryRow(sql, post.Content, post.AuthorName).Scan(post.Id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	nano := strconv.FormatInt(time.Now().UnixNano(), 10)
	content := "content" + nano
	title := "title" + nano
	post := &Post{Content: content, AuthorName: title}
	post.Create()
	p, _ := post.GetPost(31)
	fmt.Println(post, p.)
}
