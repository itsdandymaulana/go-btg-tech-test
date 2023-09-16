package service

import (
	"fmt"

	"github.com/be-tech-test-btg/app/dto"
	"github.com/be-tech-test-btg/app/model"
	"github.com/be-tech-test-btg/app/repository"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	custRepo repository.CustomerRepository
	flRepo   repository.FamilyListRepository
	natRepo  repository.NationalityRepository
}

func NewCustomerService(custRepo repository.CustomerRepository, flRepo repository.FamilyListRepository, natRepo repository.NationalityRepository) CustomerService {
	return &CustomerServiceImpl{
		custRepo: custRepo,
		flRepo:   flRepo,
		natRepo:  natRepo,
	}
}

func (s *CustomerServiceImpl) GetCustomerByID(cstId int) (dto.GetCustomerDTO, error) {
	log.Infof("Customer Service: GetCustomerByID [%s]", cstId)
	var family []dto.FamilyDTO
	cust := s.custRepo.GetCustomerByID(uint64(cstId))

	if (cust == model.Customer{}) {
		err := gorm.ErrRecordNotFound
		return dto.GetCustomerDTO{}, err
	}

	familyList := s.flRepo.GetFamilyListByCstID(cust.CstID)

	if len(familyList) > 0 {
		family = []dto.FamilyDTO(familyList)
	}

	nationality := s.natRepo.GetNationalityByCstID(cust.NationalityID)

	customer := dto.GetCustomerDTO{
		CstName:     cust.CstName,
		CstPhoneNum: cust.CstPhoneNum,
		CstDob:      cust.CstDob,
		CstEmail:    cust.CstEmail,
		Nationality: fmt.Sprintf("%s (%s)", nationality.NationalityName, nationality.NationalityCode),
		Family:      family,
	}

	return customer, nil
}
