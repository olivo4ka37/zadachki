package main

import "fmt"

type UnmatchedHero struct {
	hp int
}

func main() {
	// Присвоили структуру
	medusa := UnmatchedHero{
		hp: 16,
	}

	// присвоили адрес структуры которая хранит в себе hp 18
	achilles := &UnmatchedHero{
		hp: 18,
	}

	// Получаем идентичную форму вывода, скорей всего из-за внутреннего устройства fmt.Println
	fmt.Printf("Medusa hp is: %d, Achilles hp is: %d\n", medusa.hp, achilles.hp)

	// Смотри на разницу вывода, когда мы присваиваем адрес переменной типа, и переменной структуры.
	var a int = 14
	b := &a
	c := &UnmatchedHero{}
	fmt.Printf("a is: %v, b is: %v, c is: %v\n", a, b, c)

}
