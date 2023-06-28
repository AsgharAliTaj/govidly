package user

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	passwordHash "github.com/asgharalitaj/govidly/utils"
)

type UserHandler struct {
	userRepository Userer
}

func (u *UserHandler) UserPost(w http.ResponseWriter, r *http.Request) {
	var userReqData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &userReqData)

	userUUID := uuid.New()
	user.ID = userUUID
	user.Name = userReqData.Name
	user.Email = userReqData.Email
	user.Password = userReqData.Password

	err := validateEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	passwordString, err := passwordHash.GenerateFromPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = passwordString

	err = u.userRepository.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user created!"))
}
