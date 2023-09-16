package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/be-tech-test-btg/app/service"
	"github.com/be-tech-test-btg/config"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CustomerControllerImpl struct {
	custService service.CustomerService
}

func NewCustomerController(custService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		custService: custService,
	}
}

func (c *CustomerControllerImpl) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	var response config.Response

	custIdRaw := mux.Vars(r)["id"]

	custId, err := strconv.Atoi(custIdRaw)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Error(err)
		response.Message = "invalid id"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := c.custService.GetCustomerByID(custId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error("Record not found")
			w.WriteHeader(http.StatusNotFound)
			response.Message = "customer not found"
			json.NewEncoder(w).Encode(response)
			return
		}

		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "internal server error"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = result
	response.Message = "data found"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
