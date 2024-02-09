package service

import (
	"JobHuntz/features/company"
)

type companyService struct {
	companyData company.CompanyDataInterface
}

func New(repo company.CompanyDataInterface) company.CompanyServiceInterface {
	return &companyService{
		companyData: repo,
	}
}

func (service *companyService) RegisterCompany(input company.CompanyCore) error {
	err := service.companyData.RegisterCompany(input)
	return err
}
