package controllers

import (
 	"net/http"
	"encoding/json"
	"io/ioutil"
	"myBank/src/responses"
	"myBank/src/models"
	"myBank/src/database"
	"myBank/src/repositories"
	"myBank/src/utils"
	"github.com/gorilla/mux"
	"strconv"
	"errors"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}	
	
	
	var operation models.Operation
	if err = json.Unmarshal(requestBody, &operation); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = operation.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
	defer db.Close()

	repository := repositories.NewOperationsRepository(db)
	operation.ID, err = repository.Deposit(operation)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, operation)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}	
	
	var operation models.Operation
	if err = json.Unmarshal(requestBody, &operation); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = operation.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
	defer db.Close()

	repositoryAccount := repositories.NewAccountsRepository(db)
	originAccount, err := repositoryAccount.SearchByID(operation.Origin)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}


	if originAccount.Amount < operation.ExchangeCurrency {
		responses.Error(w, http.StatusUnauthorized, errors.New("não é possível transferir dinheiro a mais do que o total da conta."))
		return
	}

	destinyAccount, err := repositoryAccount.SearchByID(operation.Destiny)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if originAccount.Currency != "BRL" && destinyAccount.Currency != "BRL" {
		responses.Error(w, http.StatusUnauthorized, errors.New("não é possível transferir dinheiro entre contas estrangeiras"))
		return
	}

	operation.OriginCurrency = originAccount.Currency
	operation.DestinyCurrency = destinyAccount.Currency

	if originAccount.Currency != destinyAccount.Currency {
		operation.CurrencyTransformed = utils.CurrencyConversion(originAccount.Currency, destinyAccount.Currency, operation.ExchangeCurrency)
	} else {
		operation.CurrencyTransformed = operation.ExchangeCurrency
	}

	repository := repositories.NewOperationsRepository(db)
	operation.ID, err = repository.Transfer(operation)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, operation)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}	
	
	var operation models.Operation
	if err = json.Unmarshal(requestBody, &operation); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = operation.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
	defer db.Close()

	repositoryAccount := repositories.NewAccountsRepository(db)
	account, err := repositoryAccount.SearchByID(operation.Origin)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if account.Amount < operation.ExchangeCurrency {
		responses.Error(w, http.StatusUnauthorized, errors.New("não é possível retirar dinheiro a mais do que o total da conta.") )
	}

	repository := repositories.NewOperationsRepository(db)
	operation.ID, err = repository.Withdraw(operation)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, operation)
}

func GetOperations(w http.ResponseWriter, r *http.Request) { 
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

	repository := repositories.NewOperationsRepository(db)

	operations, err := repository.Search(accountID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, operations)
}
