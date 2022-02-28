package model

//Permissons like creat_user , delete_hotel
type Permission struct {
	PermissionID int
	Name         string
	Descriptoin  string
}

func Find(per *[]interface{}, perm Permission) bool {

	for _, i := range *per {
		if i == perm {
			return true
		}
	}
	return false
}
