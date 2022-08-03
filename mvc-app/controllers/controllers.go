package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/username/reponame/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to home page")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	userIdString := r.URL.Query().Get("user_id")
	userid, err := strconv.ParseUint(userIdString, 10, 64)

	if err != nil {
		http.Error(w, "please provide valid id", http.StatusBadRequest)
		return
	}
	user, err := model.FetchUser(userid)
	if err != nil {
		http.Error(w, "please provide valid userid", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}
