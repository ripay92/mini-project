package router

import (
	"GO-AUTH/controller"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r:= mux.NewRouter()
	    // User routes
    r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	// admin route
	r.HandleFunc("/admin/quiz", controller.GetQuiz).Methods("GET")
	r.HandleFunc("/admin/create-quiz", controller.CreateQuiz).Methods("POST")
	r.HandleFunc("/admin/update-quiz", controller.UpdateQuiz).Methods("POST")
	r.HandleFunc("/admin/delete-quiz", controller.DeleteQuiz).Methods("GET")

	r.HandleFunc("/admin/create-pertanyaan", controller.CreatePertanyaan).Methods("POST")
	// r.HandleFunc("/admin/update-pertanyaan", controller.UpdateQuiz).Methods("POST")
	// r.HandleFunc("/admin/delete-pertanyaan", controller.DeleteQuiz).Methods("GET")

	r.HandleFunc("/admin/jawaban-peserta", controller.CreateJawabanPeserta).Methods("POST")
    return r
}