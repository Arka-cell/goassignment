package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Arka-cell/goassignment/api/auth"
	"github.com/Arka-cell/goassignment/api/models"
	"github.com/Arka-cell/goassignment/api/responses"
	"github.com/Arka-cell/goassignment/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	shop := models.Shop{}
	err = json.Unmarshal(body, &shop)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	shop.Prepare()
	err = shop.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(shop.Email, shop.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	shop := models.Shop{}

	err = server.DB.Debug().Model(models.Shop{}).Where("email = ?", email).Take(&shop).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(shop.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(shop.ID)
}
