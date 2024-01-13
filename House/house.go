package House

import (
	"House/Structures/Devices"
	"House/Structures/Family"
	"House/Structures/Pets"
	"House/Structures/Furniture"
	"House/Structures/Rooms"
)

type House struct {
	Devices 	devices.Devices
	Family  	family.Family
	Pets    	pets.Pets
	Furniture	furniture.Furniture
	Rooms 		rooms.Rooms
}

func MakeHouse() House {
	return House {
		Devices: 	Devices.MakeDevices(),
		Family: 	Family.MakeFamily(),
		Pets: 		Pets.MakePets(),
		Furniture:  Furniture.MakeFurniture(),
		Rooms: 		Rooms.MakeRooms(),
	}
}