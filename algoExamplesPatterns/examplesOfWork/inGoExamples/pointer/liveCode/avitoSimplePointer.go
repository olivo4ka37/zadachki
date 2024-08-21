package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	// Получили копию адреса, изменили значение адреса в копии, завершили работу функции
	// именно поэтому в main функции ничего не изменится.
	person = &Person{
		Name: "Alice",
	}
	//person.Name = "Alice"
}

func main() {
	// := &Person - означает что мы присваиваем не структуру к переменной, а адрес переменной, которая хранит в себе Person{Name: Bob}
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name) // В го по адресу структуры можно вызывать её поля, fmt.Println выведет тоже самое что если бы мы присвоили
	// переменную person, а не адрес.
	changeName(person)       //Передаем копию нашего адреса переменной
	fmt.Println(person.Name) // Обращаемся все к тому адресу переменной, пишем Bob как и в первом случае.
}
