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
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&user)
	result, err := a.authRepository.GetUser(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	match, err := passwordHash.ComparePasswordAndHash(user.Password, result.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if match == true {
		w.Write([]byte("Password Matched\n"))
		return
	}
	w.Write([]byte("Password didn't match\n"))
}
