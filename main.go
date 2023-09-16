package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/be-tech-test-btg/app/controller"
	"github.com/be-tech-test-btg/app/repository"
	"github.com/be-tech-test-btg/app/service"
	"github.com/be-tech-test-btg/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	cfg            *config.Config                   = config.LoadEnv()
	db             *gorm.DB                         = config.ConnectDB(cfg)
	custRepo       repository.CustomerRepository    = repository.NewCustomerRepository(db)
	flRepo         repository.FamilyListRepository  = repository.NewFamilyListRepository(db)
	natRepo        repository.NationalityRepository = repository.NewNationalityRepository(db)
	custService    service.CustomerService          = service.NewCustomerService(custRepo, flRepo, natRepo)
	custController controller.CustomerController    = controller.NewCustomerController(custService)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/customers/{id}", custController.GetCustomerByID)

	fmt.Println("Start server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
