package router

import (
	"log"
	"net/http"

	"github.com/mizmorr/wallet/internal/controller"
)

func Handle() {
	http.HandleFunc("/postgr", func(w http.ResponseWriter, r *http.Request) {
		er := controller.ToPostgres()
		if er != nil {
			w.Write([]byte("database problem: " + er.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("everything is gOOD!"))
			w.WriteHeader(http.StatusOK)
		}
	})
	http.HandleFunc("/bounce", func(w http.ResponseWriter, r *http.Request) {
		er := controller.ViaBouncer()
		if er != nil {
			w.Write([]byte("server problem: " + er.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("all is good"))
			w.WriteHeader(http.StatusOK)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
