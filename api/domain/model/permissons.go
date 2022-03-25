package model

//Permission refers to what an authorized worker can do ,  like creat_user , delete_hotel
type Permission struct {
	PermissionID int
	Name         PermissionName
	Descriptoin  string
}

// PermissionName ...
type PermissionName string

//Permission name constant. Each model must have 4 permissions (СRUD),
//the name of the permission in the lover camel case and in the singular
const (
	ReadUser       PermissionName = "read_user"
	CreatUser      PermissionName = "creat_user"
	DeleteUser     PermissionName = "delete_user"
	UpdateUser     PermissionName = "update_user"
	ReadHotel      PermissionName = "read_hotel"
	CreatHotel     PermissionName = "creat_hotel"
	DeleteHotel    PermissionName = "delete_hotel"
	UpdateHotel    PermissionName = "update_hotel"
	ReadBooking    PermissionName = "read_booking"
	CreatBooking   PermissionName = "creat_booking"
	DeleteBooking  PermissionName = "delete_booking"
	UpdateBooking  PermissionName = "update_booking"
	ReadEmployee   PermissionName = "read_employee"
	CreatEmployee  PermissionName = "creat_employee"
	DeleteEmployee PermissionName = "delete_employee"
	UpdateEmployee PermissionName = "update_employee"
	ReadPet        PermissionName = "read_pet"
	CreatPet       PermissionName = "creat_pet"
	DeletePet      PermissionName = "delete_pet"
	UpdatePet      PermissionName = "update_pet"
	ReadRoom       PermissionName = "read_room"
	CreatRoom      PermissionName = "creat_room"
	DeleteRoom     PermissionName = "delete_room"
	UpdateRoom     PermissionName = "update_room"
	ReadSeat       PermissionName = "read_seat"
	CreatSeat      PermissionName = "creat_seat"
	DeleteSeat     PermissionName = "delete_seat"
	UpdateSeat     PermissionName = "update_seat"
)
