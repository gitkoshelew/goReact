package model

//Permission refers to what an authorized worker can do ,  like creat_user , delete_hotel
type Permission struct {
	PermissionID int
	Name         string
	Descriptoin  string
}

type PermissionName string