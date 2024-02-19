package service_test

import (
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
	"JobHuntz/features/verification"
	"JobHuntz/features/verification/service"
	"JobHuntz/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataJobseeker(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.OrderData)

	returnData := jobseeker.JobseekerCore{}

	t.Run("Success Get Jobseeker's Data", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetDataJobseeker", mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var userID uint = 1
		result, err := service.GetDataJobseeker(userID)

		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, result.ID)
		assert.Equal(t, returnData.Username, result.Username)
		assert.Equal(t, returnData.Email, result.Email)
		repo.AssertExpectations(t)
	})
}

func TestGetDataCompany(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.OrderData)

	returnData := company.CompanyCore{}

	t.Run("Success Get Company's Data", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("GetDataCompany", mock.AnythingOfType("uint")).Return(returnData, nil).Once()

		//create object service
		service := service.New(repo)

		//return dari service layer
		var userID uint = 1
		result, err := service.GetDataCompany(userID)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestAddOrderJobseeker(t *testing.T) {
	// Pembuatan object mock yang dibutuhkan
	repo := new(mocks.OrderData)

	// Menyiapkan data yang akan dikembalikan
	inputData := verification.OrderJobseekerCore{}

	t.Run("Success Add Order Jobseeker", func(t *testing.T) {
		repo.On("AddOrderJobseeker", mock.Anything).Return(
			func(input verification.OrderJobseekerCore) error {
				return nil
			},
		).Once()

		service := service.New(repo)
		err := service.AddOrderJobseeker(inputData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestAddOrderCompany(t *testing.T) {
	// Pembuatan object mock yang dibutuhkan
	repo := new(mocks.OrderData)

	// Menyiapkan data yang akan dikembalikan
	inputData := verification.OrderCompanyCore{}

	t.Run("Success AddOrderCompany", func(t *testing.T) {
		repo.On("AddOrderCompany", mock.Anything).Return(
			func(input verification.OrderCompanyCore) error {
				return nil
			},
		).Once()

		service := service.New(repo)
		err := service.AddOrderCompany(inputData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
