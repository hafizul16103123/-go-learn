package main

import "fmt"

func main() {
	pk := User{name: "piyash", email: "p@gmail.com", status: false}
	fmt.Println(pk)

	isOnline := pk.isOnline()
	newEmail := pk.newEmail()
	fmt.Println(isOnline)
	fmt.Println(newEmail)
	fmt.Println(pk)
}

type User struct {
	name   string
	email  string
	status bool
}

func (u User) isOnline() string {
	if u.status {
		return "Online"
	} else {
		return "Offline"
	}
}
func (u *User) newEmail() string {
	u.email = "pk@gmail.com"
	return u.email

}
