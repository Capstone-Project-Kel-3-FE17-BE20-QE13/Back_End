package service_test

import (
	"JobHuntz/features/vacancy"
	"JobHuntz/features/vacancy/service"
	"JobHuntz/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJobService_GetById(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	expectedJob := &vacancy.CompanyCore{ID: 1, Company_name: "Test Job"}
	mockRepo.On("GetById", uint(1)).Return(expectedJob, nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Memanggil method GetById pada JobService
	actualJob, err := jobService.GetById(1)

	// Memeriksa bahwa panggilan ke GetById dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, expectedJob, actualJob)
}

func TestJobService_CreateJob(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	mockRepo.On("CreateJob", mock.Anything).Return(nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Menyiapkan data input untuk CreateJob
	input := vacancy.Core{
		ID:              1,
		CompanyID:       1,
		Name:            "Test Job",
		Job_type:        "remote",
		Salary_range:    "10-15.000.000",
		Category:        "art",
		Job_description: "make a book",
		Job_requirement: "good in writing",
		Status:          "Lowongan dibuka",
	}

	// Memanggil method CreateJob pada JobService
	err := jobService.CreateJob(input)

	// Memeriksa bahwa panggilan ke CreateJob dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error
	assert.NoError(t, err)
}

func TestJobService_GetJobById(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	expectedJob := vacancy.Core{ID: 1, Name: "Test Job"}
	mockRepo.On("GetJobById", int(1)).Return(expectedJob, nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Memanggil method GetJobById pada JobService
	actualJob, err := jobService.GetJobById(1)

	// Memeriksa bahwa panggilan ke GetJobById dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, expectedJob, actualJob)
}

func TestJobService_DeleteJobById(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	mockRepo.On("DeleteJobById", mock.Anything, 1).Return(nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Menyiapkan data input untuk DeleteJobById
	input := []vacancy.Core{{ID: 1, Name: "Test Job"}}

	// Memanggil method DeleteJobById pada JobService
	err := jobService.DeleteJobById(input, 1)

	// Memeriksa bahwa panggilan ke DeleteJobById dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error
	assert.NoError(t, err)
}

func TestJobService_UpdateStatus(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	mockRepo.On("UpdateStatus", mock.Anything, uint(1)).Return(nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Menyiapkan data input untuk UpdateStatus
	input := vacancy.JobStatusCore{Status: "Closed"}

	// Memanggil method UpdateStatus pada JobService
	err := jobService.UpdateStatus(input, 1)

	// Memeriksa bahwa panggilan ke UpdateStatus dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error
	assert.NoError(t, err)
}

func TestJobService_GetAllJobs(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	expectedJobs := []vacancy.Core{
		{ID: 1, Name: "Test Job 1"},
		{ID: 2, Name: "Test Job 2"},
	}
	mockRepo.On("GetAllJobs").Return(expectedJobs, nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Memanggil method GetAllJobs pada JobService
	actualJobs, err := jobService.GetAllJobs()

	// Memeriksa bahwa panggilan ke GetAllJobs dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, expectedJobs, actualJobs)
}

func TestJobService_CountJobsByUserID(t *testing.T) {
	// Membuat objek mock untuk JobDataInterface
	mockRepo := new(mocks.VacancyData)

	// Mengatur perilaku yang diharapkan dari mock
	mockRepo.On("CountJobsByUserID", uint(1)).Return(5, nil)

	// Membuat objek JobService dengan menggunakan mock
	jobService := service.NewJob(mockRepo)

	// Memanggil method CountJobsByUserID pada JobService
	count, err := jobService.CountJobsByUserID(1)

	// Memeriksa bahwa panggilan ke CountJobsByUserID dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, 5, count)
}
