package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.POST("/employee", employeeHandler.Register)
	router.GET("/employee/:id", employeeHandler.GetEmployee)
	router.GET("/", statusHandler)
	router.GET("/username", usernameHandler)

	/* サーバー起動 */
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Config.Port // ローカル環境ではデフォルトのポートを設定
	}
	log.Printf("Listening on %v", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
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
