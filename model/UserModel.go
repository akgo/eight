package model

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
	Id   string
}

func NewUser (name string,age int,id string) *User {
	return &User{Name:name,Age:age,Id:id}
}

func (u *User) Login() {
	fmt.Println(" User.Login command hello ")
}

func (u *User) Logout() {
	fmt.Println(" User.Logout command Bye !")
}

func (u *User) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}
