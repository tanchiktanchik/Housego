package Pets

import (
	"fmt"
)

type Pets struct {
	Type   string //Вид питомца
	Name   string //Кличка питомца
	Colour string //Окрас питомца
	Gender bool   //Пол питомца
	Age    int8   //Возраст питомца
}

func (p Pets) Print() {
	fmt.Print("Вид: ", p.Type, "\nКличка: ", p.Name, "\nОкрас: ", p.Colour, "\nПол: ", p.Gender, "\nВозраст: ", p.Age)

	if p.Gender {
		fmt.Print("Пол: Мужской\n")
	} else {
		fmt.Print("Пол: Женский\n")
	}
}
func MakePets() []Pets {
	var pets []Pets
	Pet_1 := Pets{
		Type:   "Кот",
		Name:   "Тигра",
		Colour: "Серый с белыми полосками",
		Gender: true,
		Age:    1,
	}
	Pet_2 := Pets{
		Type:   "Кот",
		Name:   "Нира",
		Colour: "Черный",
		Gender: false,
		Age:    1,
	}
	pets = append(pets, Pet_1, Pet_2)
	return pets
}
