package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post *Post) {
	PostById[post.Id] = post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], post)
}

func main() {
	PostById1 := make(map[string]map[string]string)
	PostById1["a"] = map[string]string{}
	PostById1["a"]["b1"] = "c"
	fmt.Printf("--- \n %+v \n %+v \n", PostById1, len(PostById1))

	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	fmt.Printf("--- \n %#v \n %#v \n %+v,%+v --- \n", PostById, PostsByAuthor, len(PostById), len(PostsByAuthor))

	post1 := Post{Id: 1, Content: "hello post1", Author: "poster1"}
	post2 := Post{Id: 1, Content: "hello post2", Author: "poster2"}
	post3 := Post{Id: 1, Content: "hello post3", Author: "poster3"}

	store(&post1)
	store(&post2)
	store(&post3)
	fmt.Printf("%+v\n", PostById[2])
	fmt.Printf("%+v\n", PostsByAuthor["poster2"][0])
}
