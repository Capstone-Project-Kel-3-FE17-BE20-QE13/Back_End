package service_test

import (
	"JobHuntz/features/application"
	"JobHuntz/features/application/service"
	"JobHuntz/features/favorit"
	"JobHuntz/features/jobseeker"
	"JobHuntz/mocks"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataCompany(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.ApplicationData)

	returnData := favorit.DataCompanyCore{}

	t.Run("Success Get Company's Data", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetDataCompany", mock.AnythingOfType("*sql.DB"), mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var dbRaw *sql.DB
		var vacancyID uint = 1

		result, err := service.GetDataCompany(dbRaw, vacancyID)

		assert.NoError(t, err)
		assert.Equal(t, returnData.Company_name, result.Company_name)
		assert.Equal(t, returnData.Position, result.Position)
		repo.AssertExpectations(t)
	})
}

func TestGetMyData(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.ApplicationData)

	returnData := jobseeker.JobseekerCore{}

	t.Run("Success Get My Data", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetMyData", mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var userID uint = 1
		result, err := service.GetMyData(userID)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestCountApplication(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.ApplicationData)

	var returnData uint

	t.Run("Success Count Application", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("CountApplication", mock.AnythingOfType("*sql.DB"), mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var dbRaw *sql.DB
		var userID uint = 1

		result, err := service.CountApplication(dbRaw, userID)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestCreateApplication(t *testing.T) {
	// Pembuatan object mock yang dibutuhkan
	repo := new(mocks.ApplicationData)

	// Menyiapkan data yang akan dikembalikan
	inputData := application.Core{
		ID:                 1,
		JobseekerID:        2,
		VacancyID:          3,
		Position:           "Software Engineer",
		Company_name:       "ABC Company",
		Status_application: "Dikirim",
	}

	// var count uint = 1
	// var status string = "Unverified"

	t.Run("Success Create Application", func(t *testing.T) {
		repo.On("CreateApplication", mock.AnythingOfType("application.Core")).Return(
			func(input application.Core) error {
				return nil
			},
		).Once()

		service := service.New(repo)
		err := service.CreateApplication(inputData, 3, "Verified")

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Cannot Create Application", func(t *testing.T) {
		repo.On("CreateApplication", mock.AnythingOfType("application.Core")).Return(
			func(input application.Core) error {
				return nil
			},
		).Once()

		service := service.New(repo)
		err := service.CreateApplication(inputData, 3, "Unverified")

		assert.Error(t, err)
	})

}

func TestGetAllApplications(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.ApplicationData)

	returnData := []application.Core{}

	t.Run("Success Get My Applications", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetAllApplications", mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var jobseekerID uint = 1

		result, err := service.GetAllApplications(jobseekerID)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestGetAllApplicationsCompany(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.ApplicationData)

	returnData := []application.Core{}

	t.Run("Success Get Applications to My Company", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetAllApplicationsCompany", mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var vacancyID uint = 1

		result, err := service.GetAllApplicationsCompany(vacancyID)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestEditApplication(t *testing.T) {
	// Pembuatan object mock yang dibutuhkan
	repo := new(mocks.ApplicationData)

	// Menyiapkan data yang akan dikembalikan
	inputData := application.Core{
		Status_application: "Interview",
	}

	var id uint = 1

	t.Run("Success Edit Application", func(t *testing.T) {
		repo.On("Edit", mock.AnythingOfType("uint"), mock.AnythingOfType("application.Core")).Return(
			func(id uint, input application.Core) error {
				return nil
			},
		).Once()

		service := service.New(repo)
		err := service.EditApplication(id, inputData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
