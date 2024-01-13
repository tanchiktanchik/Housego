package devices

import (
	"fmt"
)

type Devices struct {
	Type   string  //название
	Brand  string  //Бренд
	Length float32 //Длина
	Width  float32 //Ширина
	Colour string  //Цвет
	Price  int32   //Цена
}

func (d Devices) Print() {
	fmt.Print("Название устройства: ", d.Type, "\nБренд устройства", "\nДлина устройства: ", d.Length, "\nШирина устройства: ", d.Width, "\nЦвет устройства: ", d.Colour, "\nЦена", d.Price, "\n")
}

func MakeDevices() []Devices {
	var devices []Devices
	telephone := Devices{
		Type:   "Телефон",
		Brand:  "Samsung",
		Length: 15,
		Width:  6,
		Colour: "Синий",
		Price:  15000,
	}
	laptop := Devices{
		Type:   "Ноутбук",
		Brand:  "DELL",
		Length: 35,
		Width:  25,
		Colour: "Черный",
		Price:  30000,
	}
	WashingMachine := Devices{
		Type:   "Стиральная машинка",
		Brand:  "Samsung",
		Length: 85,
		Width:  60,
		Colour: "Белая",
		Price:  40000,
	}
	TV := Devices{
		Type:   "Телевизор",
		Brand:  "LG",
		Length: 90,
		Width:  150,
		Colour: "Серый",
		Price:  75000,
	}
	devices = append(devices, telephone, laptop, WashingMachine, TV)
	return devices
}
