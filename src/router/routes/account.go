package routes

import (
	"myBank/src/controllers"
	"net/http"
)

var accountRoutes = []Route{
	{
		URI:          "/accounts",
		Method:       http.MethodPost,
		Function:     controllers.CreateAccount,
		AuthRequired: false,
	},
	// {
	// 	URI:          "/accounts",
	// 	Method:       http.MethodGet,
	// 	Function:     controllers.GetAccounts,
	// 	AuthRequired: true,
	// },
	{
		URI:          "/accounts/{accountID}",
		Method:       http.MethodGet,
		Function:     controllers.GetAccount,
		AuthRequired: true,
	},
	// {
	// 	URI:          "/accounts/{accountID}",
	// 	Method:       http.MethodDelete,
	// 	Function:     controllers.DeleteAccount,
	// 	AuthRequired: true,
	// },
	// {
	// 	URI:          "/accounts/{accountID}/changePassword",
	// 	Method:       http.MethodPost,
	// 	Function:     controllers.ChangePassword,
	// 	AuthRequired: true,
	// },
}