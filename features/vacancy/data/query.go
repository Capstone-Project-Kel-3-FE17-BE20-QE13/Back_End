package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/vacancy"
	"errors"

	"gorm.io/gorm"
)

type jobQuery struct {
	db *gorm.DB
}

func NewJob(db *gorm.DB) vacancy.JobDataInterface {
	return &jobQuery{
		db: db,
	}
}

func (repo *jobQuery) CountJobsByUserID(userID uint) (int, error) {
	var count int64
	tx := repo.db.Model(&database.Vacancy{}).Where("company_id = ?", userID).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

func (repo *jobQuery) GetById(id uint) (*vacancy.CompanyCore, error) {
	var companyData database.Company
	tx := repo.db.First(&companyData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	responData := &vacancy.CompanyCore{
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

func (repo *jobQuery) CreateJob(input vacancy.Core) error {
	newJobGorm := CoreToModel(input)

	tx := repo.db.Create(&newJobGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *jobQuery) MyCompanyVacancies(companyID uint) ([]vacancy.Core, error) {
	var jobDataGorm []database.Vacancy
	tx := repo.db.Where("company_id = ?", companyID).Find(&jobDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	allVacanciesCore := ModelGormToCore(jobDataGorm)

	return allVacanciesCore, nil
}

func (repo *jobQuery) GetAllJobs() ([]vacancy.Core, error) {
	var jobDataGorm []database.Vacancy
	tx := repo.db.Find(&jobDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allVacanciesCore := ModelGormToCore(jobDataGorm)

	return allVacanciesCore, nil
}

func (repo *jobQuery) GetJobById(id int) (vacancy.Core, error) {
	var singleJobGorm database.Vacancy
	tx := repo.db.First(&singleJobGorm, id)
	if tx.Error != nil {
		return vacancy.Core{}, tx.Error
	}

	singleVacancyCore := ModelToCore(singleJobGorm)

	return singleVacancyCore, nil
}

func (repo *jobQuery) DeleteJobById(input []vacancy.Core, id int) error {
	allVacanciesGorm := CoretoModelGorm(input)

	txDel := repo.db.Delete(&allVacanciesGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}

func (repo *jobQuery) UpdateStatus(input vacancy.JobStatusCore, vacancyID uint) error {
	jobStatusGorm := CoreStatusToModel(input)

	tx := repo.db.Model(&database.Vacancy{}).Where("id = ?", vacancyID).Updates(jobStatusGorm)

	if tx.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
