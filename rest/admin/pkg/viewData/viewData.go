package viewdata

import (
	"admin/domain/model"
	"admin/webapp/session"
	"net/http"
)

//ViewData struct needs to show permitted data in templates
type ViewData struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

//HasPermission determines have employee has access to certain data
func (vd ViewData) HasPermission(name string) bool {
	err := session.CheckRigths(vd.ResponseWriter, vd.Request, model.PermissionName(name))
	return err == nil
}
