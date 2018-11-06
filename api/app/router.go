package app

import "github.com/gorilla/mux"

const (
	V1                string = "v1"
	AUTH                     = "auth"
	PRODUCT                  = "product"
	ORDER                    = "order"
	NOTIFICATION             = "notification"
	SEP                      = "/"
	AUTH_BASE                = SEP + V1 + SEP + AUTH
	PRODUCT_BASE             = SEP + V1 + SEP + PRODUCT
	NOTIFICATION_BASE        = SEP + V1 + SEP + NOTIFICATION
	ORDER_BASE               = SEP + V1 + SEP + ORDER
)

func registerAuthRoutes(router *mux.Router) {
	router.HandleFunc(AUTH_BASE+SEP+"signup", SignUp).Methods("POST")
}

func registerProductRoutes(router *mux.Router) {
	router.HandleFunc(PRODUCT_BASE, HelloWorld).Methods("GET")
}

func registerOrderRoutes(router *mux.Router) {
	//router.HandleFunc(AUTH_BASE+SEP+"signup", SignUp).Methods("GET")
}

func registerNotificationRoutes(router *mux.Router) {
	//	router.HandleFunc(AUTH_BASE+SEP+"signup", SignUp).Methods("POST")
}

// New Instance Router
func New() *mux.Router {
	router := mux.NewRouter()

	registerAuthRoutes(router)
	registerProductRoutes(router)
	registerOrderRoutes(router)
	registerNotificationRoutes(router)

	return router
}
