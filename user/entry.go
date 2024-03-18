package user

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Create() bool {
	fmt.Println(u.Name, "has been created")
	return true
}

func (u User) Update() {
	fmt.Println(u.Name, "has been updated")
}

func (u User) Delete() {
	fmt.Println(u.Name, "has been deleted")
}

func Greetings(name string) {
	fmt.Println("Hello ", name)
}

func Get() {
	fmt.Println("All users are here")
}
