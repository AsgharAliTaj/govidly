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

func (u *UserHandler) UserGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value("id").(uuid.UUID)
	if !ok {
		http.Error(w, "id is invalid", http.StatusBadRequest)
		return
	}
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userbytes, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userbytes)
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
	ss, err := user.GenerateAuthToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("x-auth-token", ss)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created!"))
}
