package furniture

import (
	"fmt"
)

type Furniture struct {
	Type     string  //Вид мебели
	Brand    string  //Бренд мебели
	Material string  //Материал мебели
	Colour   string  //Цвет мебели
	Length   float32 //Длина мебели
	Height   float32 //Высота мебели
	Width    float32 //Ширина мебели
	Price    int32   //Цена мебели
}

func (f Furniture) Print() {
	fmt.Print("Вид мебели: ", f.Type, "\nБренд мебели: ", f.Brand, "\nМатериал мебели: ", f.Material, "\nЦвет мебели: ", f.Colour, "\nДлина мебели: ", f.Length, "\nВысота мебели: ", f.Height, "\nШирина мебели: ", f.Width, "\nЦена мебели: ", f.Price, "\n")
}

func MakeFurniture() []Furniture {
	var furniture []Furniture
	table := Furniture{
		Type:     "Стол",
		Brand:    "Dario",
		Material: "Ясень",
		Colour:   "Черный",
		Length:   180,
		Height:   80,
		Width:    100,
		Price:    12000,
	}
	chair := Furniture{
		Type:     "Стул",
		Brand:    "Flor",
		Material: "Экокожа",
		Colour:   "Серый",
		Length:   50,
		Height:   80,
		Width:    40,
		Price:    3000,
	}
	sofa := Furniture{
		Type:     "Диван",
		Brand:    "Malta",
		Material: "Ткань",
		Colour:   "Молочный",
		Length:   215,
		Height:   100,
		Width:    100,
		Price:    45000,
	}
	cabinet := Furniture{
		Type:     "Шкаф",
		Brand:    "Шведский стандарт",
		Material: "ДСП, стекло",
		Colour:   "Белый",
		Length:   150,
		Height:   200,
		Width:    50,
		Price:    20000,
	}
	furniture = append(furniture, table, chair, sofa, cabinet)
	return furniture
}
