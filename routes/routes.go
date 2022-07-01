package routes

import (
	"github.com/Threx-code/go-api/controllers"
	"github.com/Threx-code/go-api/middlewares"
	"github.com/gorilla/mux"
)

func RequestHandler(router *mux.Router) {

	router.HandleFunc("/table", controllers.CreateTable).Methods("POST")
	router.HandleFunc("/database", controllers.CreateDatabase).Methods("POST")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	// router.HandleFunc("/user", middlewares.IsAuthorized(controllers.CreateBook)).Methods("POST")
	router.HandleFunc("image/upload", middlewares.IsAuthorized(controllers.UploadImage)).Methods("POST")
}
