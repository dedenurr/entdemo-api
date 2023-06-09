package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"entdemo-api/model"
	"entdemo-api/service"
	"entdemo-api/utils"

	"github.com/go-chi/chi"
)

type userController struct {
	userService service.UserService
}

func UserNewController(userService service.UserService) *userController {
	return &userController{userService}
}



func (c *userController) UserGetAllController(w http.ResponseWriter, r *http.Request) {

	users, err := c.userService.FindAll(r.Context())
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, users)
}


func (c *userController) UserGetByIDController(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(ID)


	users, err := c.userService.FindByID(r.Context(),id)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, users)
}


func (c *userController) UserCreateController(w http.ResponseWriter, r *http.Request) {

	var newUser model.UserRequest
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	user, err :=  c.userService.UserCreate(r.Context(), newUser)
	if err != nil {
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, true, http.StatusOK, nil, user)
}
