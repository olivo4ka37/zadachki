package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func main() {
	pers := Person{1, "Jojo"}
	fmt.Println(pers.Name)
	pers.SetName("Jojo Velikyi")
	fmt.Println(pers.Name)
	(&pers).UpdateName("Jojo Velikyi 2")
	fmt.Println(pers.Name)
	(&pers).SetName("Jojo Velikyi 3")
	fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "jojojojo",
		Person: Person{
			Id:   2,
			Name: "jojojojo 2",
		},
	}

	acc.SetName("jojo 3")
	fmt.Printf("%#v \n", acc)

}
