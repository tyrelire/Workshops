package routes

import (
    "poc-workshop-go/controllers"
    "poc-workshop-go/middlewares"

    "github.com/gorilla/mux"
)

func Init(r *mux.Router) {

    r.HandleFunc("/", middlewares.LogRequest(controllers.Welcome)).Methods("GET")

    r.HandleFunc("/hello", middlewares.LogRequest(controllers.Hello)).Methods("GET")

    r.HandleFunc("/auth/hello", middlewares.LogRequest(middlewares.AuthMiddleware(controllers.AuthHello))).Methods("GET")

    r.HandleFunc("/whoami/{user}", middlewares.LogRequest(controllers.WhoAmI)).Methods("GET")

    r.HandleFunc("/add", middlewares.LogRequest(controllers.AddUser)).Methods("POST")
    r.HandleFunc("/get/{id}", middlewares.LogRequest(controllers.GetUser)).Methods("GET")
    r.HandleFunc("/del/{id}", middlewares.LogRequest(controllers.DeleteUser)).Methods("DELETE")
}
