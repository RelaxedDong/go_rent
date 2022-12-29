package main

import "fmt"

type User struct {
	// 驼峰命名，大写开头
	Id   int
	Name string
}

func main() {
	//var user = User{Id:   1, Name: "donghao"}
	//user := User{Id: 1, Name: "donghao"}
	var user User
	user.Id = 1
	user.Name = "donghao"
	fmt.Println("user is ", user)
}
