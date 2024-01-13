package rooms

import (
	"fmt"
)

type Rooms struct {
	Type     string //тип комнаты
	Square   int8   //Площадь комнаты
	Walls    string //Цвет стен
	Ceilings bool   //Потолки натяжные или нет?
	Windows  int8   //Количество окон
}

func (r Rooms) Print() {
	fmt.Print("Тип: ", r.Type, "\nПлощадь: ", r.Square, "\nЦвет стен: ", r.Walls, "\nНатяжные потолки?: ", r.Ceilings, "\nОкна: ", r.Windows, "\n")

	if r.Ceilings {
		fmt.Print("Натяжные потолки?: Да\n")
	} else {
		fmt.Print("Натяжные потолки?: Обычные\n")
	}
}
func MakeRooms() []Rooms {
	var rooms []Rooms
	bedroom := Rooms{
		Type:     "Спальня",
		Square:   15,
		Walls:    "Серые",
		Ceilings: true,
		Windows:  2,
	}
	living_room := Rooms{
		Type:     "Гостина",
		Square:   35,
		Walls:    "Молочные",
		Ceilings: true,
		Windows:  2,
	}
	kitchen := Rooms{
		Type:     "Кухня",
		Square:   20,
		Walls:    "Белые",
		Ceilings: false,
		Windows:  2,
	}
	bathroom := Rooms{
		Type:     "Ванная комната",
		Square:   15,
		Walls:    "Голубые",
		Ceilings: false,
		Windows:  1,
	}
	rooms = append(rooms, bedroom, living_room, kitchen, bathroom)
	return rooms
}
