package main

import (
	"fmt"
)

// Account структура к которой присвоили структуру User
type Account struct {
	id   uint
	Name string
	User
}

type User struct {
	Name string
}

func (u *User) writeName() {
	fmt.Printf("My name is %s\n", u.Name)
}

func main() {
	// При объявлении структуры в которую мы встроили ещё одну структуру,
	// мы должны как-бы внутри неё объявить эту встроенную структуру, для
	// установки значений.
	a := Account{
		id:   1,
		Name: "Monster",
		User: User{
			Name: "Jojo",
		},
	}
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.User.Name)
	a.writeName()

}
