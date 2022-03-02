package viewdata

import (
	"admin/webapp/session"
	"net/http"
)

type ViewData struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (vd ViewData) HasPermission(name string) bool {
	err := session.CheckRigths(vd.ResponseWriter, vd.Request, name)
	return err == nil
}
