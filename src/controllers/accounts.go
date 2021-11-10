package controllers

import (
 	"net/http"
	"encoding/json"
	"io/ioutil"
	"myBank/src/responses"
	"myBank/src/models"
	"myBank/src/database"
	"myBank/src/repositories"
	"github.com/gorilla/mux"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
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

	if err = account.Prepare(); err != nil {
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
	account.ID, err = repository.Create(account)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	accountID, err := strconv.ParseUint(parameters["accountID"], 10, 64)
	if err != nil {
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
	account, err := repository.SearchByID(accountID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, account)
}

// func GetDeposits(w http.ResponseWriter, r *http.Request) {
// 	db, err := database.Connection()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repository := repositories.NewDepositsRepository(db)

// 	deposits, err := repository.Get()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	responses.JSON(w, http.StatusOK, deposits)
// }

// func GetAmount(w http.ResponseWriter, r *http.Request) {
// 	parameter := r.URL.Query().Get("currency")

// 	db, err := database.Connection()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repository := repositories.NewDepositsRepository(db)

// 	amount, err := repository.GetAmount()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}


// 	if parameter == "USD" {
// 		responses.JSON(w, http.StatusOK, amount/5.48)
// 		return
// 	}

// 	if parameter == "EUR" {
// 		responses.JSON(w, http.StatusOK, amount/6.35)
// 		return
// 	}

// 	if parameter == "GBP" {
// 		responses.JSON(w, http.StatusOK, amount/7.43)
// 		return
// 	}

// 	responses.JSON(w, http.StatusOK, amount)
// }