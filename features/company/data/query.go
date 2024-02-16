package data

import (
	"JobHuntz/app/middlewares"
	"JobHuntz/features/company"
	"JobHuntz/utils/uploads"
	"errors"
	"mime/multipart"
	"strings"

	"gorm.io/gorm"
)

type CompanyQuery struct {
	db            *gorm.DB
	uploadService uploads.CloudinaryInterface
}

func New(db *gorm.DB, cloud uploads.CloudinaryInterface) company.CompanyDataInterface {
	return &CompanyQuery{
		db:            db,
		uploadService: cloud,
	}
}

func (repo *CompanyQuery) RegisterCompany(input company.CompanyCore) (data *company.CompanyCore, token string, err error) {
	inpuCompany := Company{
		Company_name:        input.Company_name,
		Full_name:           input.Full_name,
		Email:               input.Email,
		Password:            input.Password,
		Company_type:        input.Company_type,
		Company_size:        input.Company_size,
		Website:             input.Website,
		Status_Verification: "Unverified",
	}

	tx := repo.db.Create(&inpuCompany)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, "", errors.New("insert failed, row affected = 0")
	}

	var companyGorm Company
	tx = repo.db.Where("email = ?", input.Email).First(&companyGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	result := companyGorm.ModelRegisterToCore()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}

// LoginCompany implements company.CompanyDataInterface.
func (repo *CompanyQuery) LoginCompany(email string, password string) (data *company.CompanyCore, err error) {
	var companyGorm Company
	tx := repo.db.Where("email = ?", email).First(&companyGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := companyGorm.ModelRegisterToCore()

	return &result, nil
}

// GetById implements company.CompanyDataInterface.
func (repo *CompanyQuery) GetById(id uint) (*company.CompanyCore, error) {
	var companyData Company
	tx := repo.db.First(&companyData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	responData := &company.CompanyCore{
		ID:                  companyData.ID,
		Company_name:        companyData.Company_name,
		Full_name:           companyData.Full_name,
		Email:               companyData.Email,
		Company_type:        companyData.Company_type,
		Company_size:        companyData.Company_size,
		Website:             companyData.Website,
		Description:         companyData.Description,
		Status_Verification: companyData.Status_Verification,
		Banners:             companyData.Banners,
		Address:             companyData.Address,
		Phone:               companyData.Phone,
	}

	return responData, nil
}

// UpdateCompany implements company.CompanyDataInterface.
func (repo *CompanyQuery) UpdateCompany(id int, input company.CompanyCore, file multipart.File, nameFile string) error {
	var folderName string = "img/company"

	imgGorm, _ := repo.GetCompanyImageById(id)

	if imgGorm != "" {
		splitImgSlash := strings.Split(imgGorm, "/")
		publicIdSliceSlash := splitImgSlash[7:10]
		publicIdGormSlash := strings.Join(publicIdSliceSlash, "/")

		splitPublicId := strings.Split(publicIdGormSlash, ".")
		publicIdSliced := splitPublicId[0:(len(splitPublicId) - 1)]
		publicId := strings.Join(publicIdSliced, ".")

		repo.uploadService.Destroy(publicId)
	}

	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}

		input.Banners = imgUrl.SecureURL
	}

	companyGorm := CoreModelCompanyUpdate(input)
	tx := repo.db.Model(&Company{}).Where("id = ?", id).Updates(companyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

func (repo *CompanyQuery) GetCompanyImageById(companyId int) (string, error) {
	var company Company
	if err := repo.db.Table("companies").Where("id = ?", companyId).First(&company).Error; err != nil {
		return "", err
	}

	return company.Banners, nil
}
