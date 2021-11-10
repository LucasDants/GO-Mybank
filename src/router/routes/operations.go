package routes

import (
	"myBank/src/controllers"
	"net/http"
)

var operationsRoutes = []Route{
	{
		URI:          "/deposit",
		Method:       http.MethodPost,
		Function:     controllers.Deposit,
		AuthRequired: false,
	},
	{
		URI:          "/transfer",
		Method:       http.MethodPost,
		Function:     controllers.Transfer,
		AuthRequired: true,
	},
	{
		URI:          "/withdraw",
		Method:       http.MethodPost,
		Function:     controllers.Withdraw,
		AuthRequired: true,
	},
	{
		URI:          "/operations/{accountID}",
		Method:       http.MethodGet,
		Function:     controllers.GetOperations,
		AuthRequired: true,
	},

}