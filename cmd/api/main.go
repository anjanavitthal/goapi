package main

import (
	"fmt"
	"net/http"

	"github.com/anjanavitthal/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println(`Server starts`)
	fmt.Println(`
	 ____  ___       _    ____ ___ 
	/ ___|/ _ \     / \  |  _ \_ _|
   | |  _| | | |   / _ \ | |_) | | 
   | |_| | |_| |  / ___ \|  __/| | 
	\____|\___/  /_/   \_\_|  |___|
								   

   `)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
