package data

import (
	"JobHuntz/app/database"
	"JobHuntz/features/application"
	"JobHuntz/features/favorit"
	"JobHuntz/features/jobseeker"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ApplyQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) application.ApplyDataInterface {
	return &ApplyQuery{
		db: db,
	}
}

func (repo *ApplyQuery) GetDataCompany(dbRaw *sql.DB, vacancyID uint) (favorit.DataCompanyCore, error) {
	// simpan ke DB
	var dataCompany favorit.DataCompanyCore

	query := `SELECT vacancies.name, companies.company_name 
	FROM companies JOIN vacancies ON companies.id = vacancies.company_id
	WHERE vacancies.id = ?`

	tx := dbRaw.QueryRow(query, vacancyID)
	if err := tx.Scan(&dataCompany.Position, &dataCompany.Company_name); err != nil {
		//
	}

	return dataCompany, nil
}

func (repo *ApplyQuery) GetMyData(userID uint) (jobseeker.JobseekerCore, error) {
	// simpan ke DB
	var myDatas database.Jobseeker

	if err := repo.db.First(&myDatas, "id = ?", userID).Error; err != nil {
		return jobseeker.JobseekerCore{}, errors.New(err.Error())
	}

	myDatasCore := ModelGormJobseekerToCore(myDatas)

	fmt.Println("isi dataku: ", myDatas)

	return myDatasCore, nil
}

func (repo *ApplyQuery) CountApplication(dbRaw *sql.DB, userID uint) (uint, error) {
	var count uint

	query := `select count(status_application) from applications where jobseeker_id = ?;`

	tx := dbRaw.QueryRow(query, userID)
	if err := tx.Scan(&count); err != nil {
		//
	}

	fmt.Println("isi countku: ", count)

	return count, nil
}

func (repo *ApplyQuery) CreateApplication(input application.Core) error {
	// simpan ke DB
	newApplyGorm := CoreToModel(input)

	tx := repo.db.Create(&newApplyGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *ApplyQuery) GetAllApplications(jobseekerID uint) ([]application.Core, error) {
	var applicationsDataGorm []database.Application
	tx := repo.db.Where("jobseeker_id = ?", jobseekerID).Find(&applicationsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allApplicationsCore := ModelGormToCore(applicationsDataGorm)

	return allApplicationsCore, nil
}

// GetAllApplicationsCompany implements application.ApplyDataInterface.
func (repo *ApplyQuery) GetAllApplicationsCompany(vacancyID uint) ([]application.Core, error) {
	var applicationsDataGorm []database.Application

	tx := repo.db.Where("vacancy_id = ?", vacancyID).Find(&applicationsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allApplicationsCore := ModelGormToCore(applicationsDataGorm)

	return allApplicationsCore, nil
}

// func (repo *ApplyQuery) GetAllApplicationsCompany(dbRaw *sql.DB, vacancyID_int int) ([]application.ListApplicantsCore, error) {
// 	var listApplicants []application.ListApplicantsCore

// 	query := `SELECT applications.id, applications.jobseeker_id, jobseekers.full_name, jobseekers.username, jobseekers.email,
// 	applications.vacancy_id, applications.position, applications.company_name, applications.status_application
// 	FROM applications
// 	JOIN jobseekers ON applications.jobseeker_id = jobseekers.id
// 	WHERE applications.vacancy_id = ?`

// 	rows, err := dbRaw.Query(query, vacancyID_int)
// 	if err != nil {
// 		return nil, errors.New("failed to read data applicants" + err.Error())
// 	}

// 	for rows.Next() {
// 		var dataApplicant application.ListApplicantsCore
// 		if err := rows.Scan(&dataApplicant.ID, &dataApplicant.JobseekerID, &dataApplicant.Full_name, &dataApplicant.Username, &dataApplicant.Email, &dataApplicant.VacancyID, &dataApplicant.Position, &dataApplicant.Company_name, &dataApplicant.Status_application); err != nil {
// 			return nil, errors.New("failed to read each data applicant" + err.Error())
// 		}
// 		listApplicants = append(listApplicants, dataApplicant)
// 	}

// 	return listApplicants, nil
// }

func (repo *ApplyQuery) Edit(id uint, input application.Core) error {
	dataApplication := database.Application{
		Status_application: input.Status_application,
	}

	tx := repo.db.Model(&database.Application{}).Where("id = ?", id).Updates(&dataApplication)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("edit failed, row affected = 0")
	}

	return nil
}
