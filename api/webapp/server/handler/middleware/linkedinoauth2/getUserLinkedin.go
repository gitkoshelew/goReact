package linkedinoauth2

import (
	"encoding/json"
	"fmt"
	reqandresp "goReact/domain/reqAndResp"
	"goReact/domain/store"
	"goReact/service"
	"goReact/webapp/server/handler"
	"goReact/webapp/server/handler/response"
	"net/http"
	"os"
	"reflect"

	"golang.org/x/oauth2"
)

var getUserURI = os.Getenv("LINKEDIN_GET_USER_DATA")

func GetUserLinkedIn(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Context().Value(handler.CTXKeyAccessTokenLinkedINOAuth).(*oauth2.Token)

		authHeader := fmt.Sprintf("Bearer %s", token.AccessToken)

		w.Header().Set("Accept", "application/json")

		result, err := service.GetUserData(getUserURI, authHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occured while getting resource . Err msg: %v", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Error occured while getting resource . Err msg: %v", err)})
			return
		}

		//	re := bytes.NewReader(*result)
		linkedInSSOUser := &reqandresp.LinkedInSSOUser{}

		/*err = binary.Read(re, binary.BigEndian, *linkedInSSOUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		*/
		err = json.Unmarshal(*result, linkedInSSOUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Error occurred  while unmarshaling bytes. Err msg:%v.", err)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}
		fmt.Println("linkedInSSOUser -   :", reflect.TypeOf((linkedInSSOUser.Photo)))
		fmt.Println("linkedInSSOUser -   :", reflect.ValueOf((linkedInSSOUser.Photo)))


		/*if err := json.NewDecoder(re).Decode(linkedInSSOUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.Body)
			json.NewEncoder(w).Encode(response.Error{Messsage: err.Error()})
			return
		}*/

		json.NewEncoder(w).Encode(linkedInSSOUser)
	}
}
