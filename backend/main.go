package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yukyoooo/go_next_ddd/application"
	"github.com/yukyoooo/go_next_ddd/config"
	"github.com/yukyoooo/go_next_ddd/domain/model"
	"github.com/yukyoooo/go_next_ddd/domain/model/employee"
	"github.com/yukyoooo/go_next_ddd/infrastructure/repository"
	rest "github.com/yukyoooo/go_next_ddd/interface/handler"
)

func main() {
	fmt.Println(model.Db)

	/* 依存関係を定義 */
	employeeRepository := repository.NewEmployeeRepository(model.Db)
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeApplicationService := application.NewEmployeeApplicationService(employeeRepository, *employeeService)
	employeeHandler := rest.NewEmployeeHandler(*employeeApplicationService)

	/* ルーティング */
	router := httprouter.New()
	router.POST("/api/employee", employeeHandler.Register)
	router.GET("/api/employee/:id", employeeHandler.GetEmployee)
	router.GET("/status", statusHandler)
	router.GET("/username", usernameHandler)

	http.ListenAndServe(config.Config.Port, &Server{router})
	log.Fatal(http.ListenAndServe(config.Config.Port, router))

}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "https://example.com")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}

func statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var resp = []byte(`{"status": "ok"}`)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.Write(resp)
}

func usernameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var resp = []byte(`{"username": "yukyoooo"}`)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.Write(resp)
}
