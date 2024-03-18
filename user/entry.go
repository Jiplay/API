package user

import "fmt"
import "api/jwt"

type User struct {
	Name string
	Age  int
}

func (u User) Create() string {
	fmt.Println(u.Name, "has been created")
	token, err := jwt.CreateToken(u.Name)
	if err != nil {
		return ""
	}

	return token
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
