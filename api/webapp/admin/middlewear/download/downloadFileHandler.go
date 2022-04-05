package download

import (
	"fmt"
	"goReact/domain/store"
	"goReact/webapp/admin/middlewear"
	"goReact/webapp/admin/session"
	"io"
	"net/http"
	"os"
	"strings"
)

// DownloadFileHandler ...
func DownloadFileHandler(s *store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session.CheckSession(w, r)

		path, err := middlewear.CtxFile(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while parsing file: %v", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing file: %v", err)
			return
		}

		file, err := os.Open(path)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while parsing file: %v", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while parsing file: %v", err)
			return
		}

		arr := strings.Split(file.Name(), "/")
		name := arr[len(arr)-1]

		w.Header().Set("Accept-ranges", "bytes")
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename="+name+"")
		w.WriteHeader(http.StatusOK)

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occured while sending file: %v", err), http.StatusInternalServerError)
			s.Logger.Errorf("Error occured while sending file: %v", err)
			return
		}

		s.Logger.Info("File sent")
	}
}
