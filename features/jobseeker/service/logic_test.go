package service

import (
	"JobHuntz/features/jobseeker"
	"JobHuntz/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetId(t *testing.T) {
	repo := new(mocks.JobseekerData)
	// hash := new(hashMock.HashService)
	srv := New(repo)
	returnData := &jobseeker.JobseekerCore{
		ID:         1,
		Full_name:  "A",
		Username:   "b",
		Email:      "c",
		Address:    "d",
		Phone:      "e",
		Birth_date: "f",
		Gender:     "g",
		Resume:     "h",
	}

	repo.On("GetByIdJobSeeker", uint(1)).Return(returnData, nil)
	result, err := srv.GetByIdJobSeeker(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRegister(t *testing.T) {
	repo := new(mocks.JobseekerData)
	service := JobseekerService{jobseekerData: repo}

	input := jobseeker.JobseekerCore{
		Full_name: "A",
		Username:  "B",
		Email:     "C",
		Password:  "D",
	}

	repo.On("Register", input).Return(nil)
	err := service.Register(input)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

// func TestLogin(t *testing.T) {
// 	mockJobseekerData := new(mocks.JobseekerData)
// 	// mockJobseekerData := new(MockJobseekerData)

// 	// Create an instance of JobseekerService with the mock
// 	service := &JobseekerService{jobseekerData: mockJobseekerData}

// 	// Mock data for testing
// 	mockJobseeker := jobseeker.JobseekerCore{
// 		ID:       1,
// 		Email:    "test@example.com",
// 		Password: "hashed_password", // Assuming the password is hashed
// 		// Add other fields as needed for testing
// 	}

// 	// Mock the behavior of JobseekerData's Login method
// 	mockJobseekerData.On("Login", "test@example.com").Return(mockJobseeker, nil)
// 	mockJobseekerData.On("Login", "").Return(jobseeker.JobseekerCore{}, errors.New("email is required"))

// 	// Test case: Successful login
// 	jobseekerData, token, err := service.Login("test@example.com", "password")
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockJobseeker, jobseekerData)
// 	assert.NotEmpty(t, token)
// }

func TestLogin_WrongPassword(t *testing.T) {
	mockJobseekerData := new(mocks.JobseekerData)
	service := &JobseekerService{jobseekerData: mockJobseekerData}

	expectedJobseeker := jobseeker.JobseekerCore{ID: 1, Email: "test@example.com"}

	mockJobseekerData.On("Login", "test@example.com").Return(expectedJobseeker, nil)

	_, _, err := service.Login("test@example.com", "wrong_password")

	assert.EqualError(t, err, "login failed, wrong password")
}

func TestLogin_EmptyEmail(t *testing.T) {
	service := &JobseekerService{}

	_, _, err := service.Login("", "password")

	assert.EqualError(t, err, "email is required")
}

func TestLogin_EmptyPassword(t *testing.T) {
	service := &JobseekerService{}

	_, _, err := service.Login("test@example.com", "")

	assert.EqualError(t, err, "password is required")
}
