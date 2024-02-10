package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/company"
	"JobHuntz/utils/responses"
	"errors"
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

func (service *companyService) LoginCompany(email string, password string) (company.CompanyCore, string, error) {
	if email == "" {
		return company.CompanyCore{}, "", errors.New("email is required")
	} else if password == "" {
		return company.CompanyCore{}, "", errors.New("password is required")
	}

	ressLogin, err := service.companyData.LoginCompany(email)
	if err != nil {
		return company.CompanyCore{}, "", errors.New(err.Error() + "login error, cannot retrieve data")
	}

	cekPass := responses.ComparePassword(password, ressLogin.Password)
	if !cekPass {
		return company.CompanyCore{}, "", errors.New("login failed, wrong password")
	}

	token, err := middlewares.CreateToken(ressLogin.ID)
	if err != nil {
		return company.CompanyCore{}, "", errors.New(err.Error() + "cannot create token")
	}
	return ressLogin, token, nil
}
