package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/yukyoooo/go_next_ddd/application"
)

type EmployeeHandler interface {
	GetEmployee(http.ResponseWriter, *http.Request, httprouter.Params)
	Register(http.ResponseWriter, *http.Request, httprouter.Params)
}

type employeeHandler struct {
	employeeApplicationService application.EmployeeApplicationService
}

func NewEmployeeHandler(employeeApplicationService application.EmployeeApplicationService) EmployeeHandler {
	return &employeeHandler{employeeApplicationService: employeeApplicationService}
}

func (eh *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idString := ps.ByName("id")
	idInt, _ := strconv.Atoi(idString)
	employee, err := eh.employeeApplicationService.GetEmployeeById(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func (eh *employeeHandler) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := eh.employeeApplicationService.Register("taro", "yamada", "aaa@example.com", "password", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success")
}
