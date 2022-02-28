package upload

import (
	"admin/domain/store"
	"admin/webapp/session"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func UploadFile(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
	
		if r.Method == "POST" {
			src, hdr, err := r.FormFile("file")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
				return
			}
			defer src.Close()

			dst, err := os.Create(filepath.Join(os.TempDir(), hdr.Filename))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Can not create file. Err msg:%v. ", err)
				return
			}
			defer dst.Close()

			io.Copy(dst, src)
		}

		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, `
          <form method="POST" enctype="multipart/form-data">
            <input type="file" name="file">
            <input type="submit">
          </form>
          `)
	}
}
