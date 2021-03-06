package server

import (
	"belanjayukid_go/http/server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func buildRouteHandler() http.Handler {
	router := mux.NewRouter()
	routePost(router)

	return router
}

func routePost(router *mux.Router) {

	//auth
	router.HandleFunc("/v1/auth/login", handlers.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/register", handlers.HandleRegister).Methods(http.MethodPost)
	router.HandleFunc("/v1/me", handlers.Auth(handlers.HandleGetMe)).Methods(http.MethodGet)

	//transaction
	router.HandleFunc("/v1/transactions", handlers.HandleGetTransactionList).Methods(http.MethodGet)
	router.HandleFunc("/v1/transactions/{transactionID}", handlers.HandleGetTransactionDetail).Methods(http.MethodGet)
	router.HandleFunc("/v1/transactions", handlers.HandleInitTransaction).Methods(http.MethodPost)
	router.HandleFunc("/v1/transactions", handlers.HandleAddToCartTransaction).Methods(http.MethodPatch)
	router.HandleFunc("/v1/transactions/{transactionID}/finish", handlers.HandleFinishTransaction).Methods(http.MethodPost)

	//category
	router.HandleFunc("/v1/categories", handlers.HandleCreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/v1/categories", handlers.HandleGetCategoryList).Methods(http.MethodGet)

	//productunit
	router.HandleFunc("/v1/productunits", handlers.HandleCreateProductUnit).Methods(http.MethodPost)
	router.HandleFunc("/v1/productunits", handlers.HandleGetProductUnitList).Methods(http.MethodGet)

	//product
	router.HandleFunc("/v1/products", handlers.HandleGetProductList).Methods(http.MethodGet)
}
