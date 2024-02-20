package service

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/company"
	"JobHuntz/utils/encrypts"
	"errors"
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type companyService struct {
	companyData company.CompanyDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

func New(repo company.CompanyDataInterface, hash encrypts.HashInterface) company.CompanyServiceInterface {
	return &companyService{
		companyData: repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// RegisterCompany implements company.CompanyServiceInterface.
func (service *companyService) RegisterCompany(input company.CompanyCore) (data *company.CompanyCore, token string, err error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return nil, "", errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return nil, "", errors.New("rror hashing password")
		}
		input.Password = hashedPass
	}

	data, generatedToken, err := service.companyData.RegisterCompany(input)
	return data, generatedToken, err
}

// LoginCompany implements company.CompanyServiceInterface.
func (service *companyService) LoginCompany(email string, password string) (data *company.CompanyCore, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.companyData.LoginCompany(email, password)
	if err != nil {
		return nil, "", errors.New("email atau password salah")
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

// GetById implements company.CompanyServiceInterface.
func (service *companyService) GetById(id uint) (*company.CompanyCore, error) {
	result, err := service.companyData.GetById(id)
	return result, err
}

// UpdateCompany implements company.CompanyServiceInterface.
func (service *companyService) UpdateCompany(id int, input company.CompanyCore, file multipart.File, nameFile string) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.companyData.UpdateCompany(id, input, file, nameFile)
	return err
}
