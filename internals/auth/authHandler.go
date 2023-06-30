package auth

import (
	"encoding/json"
	"net/http"

	passwordHash "github.com/asgharalitaj/govidly/utils"
)

type AuthHandler struct {
	authRepository Auther
}

func (a *AuthHandler) UserGet(w http.ResponseWriter, r *http.Request) {
	var userReqData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&userReqData)
	user, err := a.authRepository.GetUser(userReqData.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	match, err := passwordHash.ComparePasswordAndHash(userReqData.Password, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if match == true {
		ss, err := user.GenerateAuthToken()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ss))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Password didn't match\n"))
}
