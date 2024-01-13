package family

import (
	"fmt"
)

type FamilyMembers struct {
	Name    string //Имя
	Surname string //Фамилия
	Gender  bool   //Пол
	Zodiac  string //Знак зодиака
	Age     int8   //Возраст
	Job     bool   //Наличие работы
}

func (f FamilyMembers) Print() {
	fmt.Print("Имя: ", f.Name, "\nФамилия: ", f.Surname, "\nПол: ", f.Gender, "\nЗнак зодиака: ", f.Zodiac, "\nВозраст: ", f.Age, "\nРаботает?", f.Job, "\n")

	if f.Gender {
		fmt.Print("Пол: Мужской\n")
	} else {
		fmt.Print("Пол: Женский\n")
	}
	if f.Job {
		fmt.Print("Работает? Да\n")
	} else {
		fmt.Print("Работает? Учится\n")
	}
}
func MakeFamilyMembers() []FamilyMembers {
	var family_members []FamilyMembers
	mother := FamilyMembers{
		Name:    "Татьяна",
		Surname: "Чанкаева",
		Gender:  false,
		Zodiac:  "Рыбы",
		Age:     42,
		Job:     true,
	}
	father := FamilyMembers{
		Name:    "Василий",
		Surname: "Чанкаев",
		Gender:  true,
		Zodiac:  "Телец",
		Age:     48,
		Job:     true,
	}
	brother := FamilyMembers{
		Name:    "Руслан",
		Surname: "Чанкаев",
		Gender:  true,
		Zodiac:  "Козерог",
		Age:     26,
		Job:     true,
	}
	sister := FamilyMembers{
		Name:    "Алтана",
		Surname: "Чанкаева",
		Gender:  false,
		Zodiac:  "Дева",
		Age:     15,
		Job:     false,
	}
	creator := FamilyMembers{
		Name:    "Татьяна",
		Surname: "Чанкаева",
		Gender:  false,
		Zodiac:  "Водолей",
		Age:     20,
		Job:     false,
	}
	family_members = append(family_members, mother, father, brother, sister, creator)
	return family_members
}
