package controllers

import (
	"myBank/src/auth"
	"myBank/src/database"
	"myBank/src/models"
	"myBank/src/repositories"
	"myBank/src/responses"
	"myBank/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account
	if err = json.Unmarshal(requestBody, &account); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountsRepository(db)
	databaseAccount, err := repository.SearchByName(account.Name)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(databaseAccount.Password, account.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(databaseAccount.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	accountID := strconv.FormatUint(databaseAccount.ID, 10)

	responses.JSON(w, http.StatusOK, models.Auth{ID: accountID, Token: token})
}