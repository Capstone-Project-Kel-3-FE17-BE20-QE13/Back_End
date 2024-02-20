package service

import (
	"JobHuntz/features/jobseeker"
	"JobHuntz/mocks"
	"database/sql"
	"errors"
	"mime/multipart"

	"testing"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/stretchr/testify/assert"
)

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

func TestGetId(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.JobseekerCore{
		ID:                  1,
		Full_name:           "A",
		Username:            "B",
		Email:               "C",
		Password:            "D",
		Address:             "E",
		Phone:               "F",
		Birth_date:          "G",
		Gender:              "H",
		Resume:              "I",
		Status_Verification: "J",
		Banners:             "K",
		Roles:               "L",
	}
	repo.On("GetByIdJobSeeker", uint(1)).Return(returnData, nil)
	result, err := srv.GetByIdJobSeeker(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestJobsekerByCompany(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.JobseekerCore{
		ID:                  1,
		Full_name:           "A",
		Username:            "B",
		Email:               "C",
		Password:            "D",
		Address:             "E",
		Phone:               "F",
		Birth_date:          "G",
		Gender:              "H",
		Resume:              "I",
		Status_Verification: "J",
		Banners:             "K",
		Roles:               "L",
	}
	repo.On("GetByIdJobSeeker", uint(1)).Return(returnData, nil)
	repo.On("GetjobseekerByCompany", uint(1)).Return(returnData, nil)

	result, err := srv.GetjobseekerByCompany(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestEduById(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.EducationCore{
		ID:              1,
		JobseekerID:     1,
		Education_level: "A",
		Major:           "B",
		Graduation_date: "C",
	}

	repo.On("GetEduByID", uint(1)).Return(*returnData, nil)
	result, err := srv.GetEduByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLicenseById(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.LicenseCore{
		ID:             1,
		JobseekerID:    1,
		License_name:   "A",
		Published_date: "B",
		Expiry_date:    "C",
		License_file:   "D",
	}

	repo.On("GetLicenseByID", uint(1)).Return(*returnData, nil)
	result, err := srv.GetLicenseByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSkillByid(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.SkillCore{
		ID:          1,
		JobseekerID: 1,
		Skill:       "A",
		Description: "B",
	}

	repo.On("GetSkillByID", uint(1)).Return(*returnData, nil)
	result, err := srv.GetSkillByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateCv(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.CVCore{
		ID:          1,
		JobseekerID: 1,
		CV_file:     "test.pdf",
	}

	repo.On("AddCV", *returnData).Return(errors.New("your cv is already exist"))

	err := srv.AddCV(*returnData, 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "your cv is already exist")
}

func TestCreatAddCareer(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.CareerCore{
		ID:           1,
		JobseekerID:  1,
		Position:     "A",
		Company_name: "B",
		Date_start:   "C",
		Date_end:     "D",
	}

	repo.On("AddCareer", *returnData).Return(errors.New("your cv is already exist"))
	err := srv.AddCareer(*returnData)
	assert.Error(t, err)
	assert.EqualError(t, err, "your cv is already exist")
}

func TestCreateEducation(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.EducationCore{
		ID:              1,
		JobseekerID:     2,
		Education_level: "A",
		Major:           "B",
		Graduation_date: "C",
	}

	repo.On("AddEducation", *returnData).Return(errors.New("your cv is already exist"))
	err := srv.AddEducation(*returnData)
	assert.Error(t, err)
	assert.EqualError(t, err, "your cv is already exist")
}

func TestCreateLicense(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.LicenseCore{
		ID:             1,
		JobseekerID:    1,
		License_name:   "A",
		Published_date: "B",
		Expiry_date:    "C",
		License_file:   "D",
	}

	repo.On("AddLicense", *returnData).Return(errors.New("your cv is already exist"))
	err := srv.AddLicense(*returnData)
	assert.Error(t, err)
	assert.EqualError(t, err, "your cv is already exist")
}

func TestCreateSkill(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.SkillCore{
		ID:          1,
		JobseekerID: 1,
		Skill:       "A",
		Description: "B",
	}

	repo.On("AddSkill", *returnData).Return(errors.New("your cv is already exist"))
	err := srv.AddSkill(*returnData)
	assert.Error(t, err)
	assert.EqualError(t, err, "your cv is already exist")
}

func TestAllEmail(t *testing.T) {
	repo := new(mocks.JobseekerData)
	email := "test@example.com"
	expectedJobseekerCore := jobseeker.JobseekerCore{
		Email: email,
	}
	expectedError := errors.New("expected error")

	repo.On("AllEmails", email).Return(expectedJobseekerCore, nil)
	repo.On("AllEmails", "nonexistent@example.com").Return(jobseeker.JobseekerCore{}, expectedError)

	resultJobseekerCore, err := repo.AllEmails(email)
	assert.NoError(t, err)
	assert.Equal(t, expectedJobseekerCore, resultJobseekerCore)

	resultJobseekerCore, err = repo.AllEmails("nonexistent@example.com")
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, jobseeker.JobseekerCore{}, resultJobseekerCore)

	repo.AssertExpectations(t)
}

func TestAllUserName(t *testing.T) {
	mockJobseekerData := new(mocks.JobseekerData)
	username := "testuser"
	expectedJobseekerCore := jobseeker.JobseekerCore{
		Username: username,
	}
	expectedError := errors.New("expected error")

	mockJobseekerData.On("AllUsernames", username).Return(expectedJobseekerCore, nil)
	mockJobseekerData.On("AllUsernames", "nonexistentuser").Return(jobseeker.JobseekerCore{}, expectedError)

	resultJobseekerCore, err := mockJobseekerData.AllUsernames(username)
	assert.NoError(t, err)
	assert.Equal(t, expectedJobseekerCore, resultJobseekerCore)

	resultJobseekerCore, err = mockJobseekerData.AllUsernames("nonexistentuser")
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, jobseeker.JobseekerCore{}, resultJobseekerCore)

	mockJobseekerData.AssertExpectations(t)
}

func TestCountCv(t *testing.T) {
	mockJobseekerData := new(mocks.JobseekerData)
	service := New(mockJobseekerData)
	dbMock := &sql.DB{}
	seekerID := uint(123)
	expectedCount := uint(5)
	expectedError := errors.New("dummy error")

	mockJobseekerData.On("CountCV", dbMock, seekerID).Return(expectedCount, nil).Once()
	mockJobseekerData.On("CountCV", dbMock, seekerID).Return(uint(0), expectedError).Once()

	count, err := service.CountCV(dbMock, seekerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, count)

	count, err = service.CountCV(dbMock, seekerID)
	assert.Error(t, err)
	assert.Equal(t, uint(0), count)
	assert.Equal(t, expectedError, err)

	mockJobseekerData.AssertExpectations(t)
}

func TestGetCareerByID(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	returnData := &jobseeker.CareerCore{
		ID:           1,
		JobseekerID:  1,
		Position:     "A",
		Company_name: "B",
		Date_start:   "C",
		Date_end:     "D",
	}
	repo.On("GetCareerByID", uint(1)).Return(*returnData, nil)
	result, err := srv.GetCareerByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCareerList(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	userId := uint(1)
	returnData := []jobseeker.CareerCore{
		{
			ID:           1,
			JobseekerID:  userId,
			Position:     "A",
			Company_name: "B",
			Date_start:   "C",
			Date_end:     "D",
		},
	}

	repo.On("GetCareerList", userId).Return(returnData, nil).Once()

	result, err := srv.GetCareerList(userId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, returnData, result)

	repo.AssertCalled(t, "GetCareerList", userId)
}

func TestEduList(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	userId := uint(1)
	returnData := []jobseeker.EducationCore{
		{
			ID:              1,
			JobseekerID:     userId,
			Education_level: "A",
			Major:           "B",
			Graduation_date: "C",
		},
	}

	repo.On("GetEduList", userId).Return(returnData, nil).Once()

	result, err := srv.GetEduList(userId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, returnData, result)

	repo.AssertCalled(t, "GetEduList", userId)
}

func TestLicenseList(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	userId := uint(1)
	returnData := []jobseeker.LicenseCore{
		{
			ID:             1,
			JobseekerID:    userId,
			License_name:   "A",
			Published_date: "B",
			Expiry_date:    "C",
			License_file:   "D",
		},
	}

	repo.On("GetLicenseList", userId).Return(returnData, nil).Once()

	result, err := srv.GetLicenseList(userId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, returnData, result)

	repo.AssertCalled(t, "GetLicenseList", userId)
}

func TestSkillList(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	userId := uint(1)
	returnData := []jobseeker.SkillCore{
		{
			ID:          1,
			JobseekerID: userId,
			Skill:       "A",
			Description: "B",
		},
	}

	repo.On("GetSkillList", userId).Return(returnData, nil).Once()

	result, err := srv.GetSkillList(userId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, returnData, result)

	repo.AssertCalled(t, "GetSkillList", userId)
}

func TestPDF(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	inputFileHeader := &multipart.FileHeader{
		Filename: "A",
		Header:   map[string][]string{},
		Size:     1,
	}
	expectedOutput := &s3manager.UploadOutput{
		Location:  "A",
		VersionID: new(string),
		UploadID:  "B",
		ETag:      new(string),
	}
	repo.On("PDF", inputFileHeader).Return(expectedOutput, nil)
	output, err := srv.PDF(inputFileHeader)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedOutput, output)
	assert.Nil(t, err)
}

func TestPhoto(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	inputFileHeader := &multipart.FileHeader{
		Filename: "A",
		Header:   map[string][]string{},
		Size:     1,
	}
	expectedOutput := &uploader.UploadResult{
		Response: repo,
	}
	repo.On("Photo", inputFileHeader).Return(expectedOutput, nil)
	output, err := srv.Photo(inputFileHeader)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedOutput, output)
	assert.Nil(t, err)
}

func TestReadCv(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedCV := jobseeker.CVCore{
		ID:          1,
		JobseekerID: seekerID,
		CV_file:     "A",
	}

	repo.On("ReadCV", seekerID).Return(expectedCV, nil)
	cv, err := srv.ReadCV(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedCV, cv)
	assert.Nil(t, err)
}

func TestRegister(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	input := jobseeker.JobseekerRegistCore{
		Full_name: "A",
		Username:  "B",
		Email:     "C",
		Password:  "D",
	}

	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("Register", input).Return(expectedErr)

	err := srv.Register(input)
	assert.Equal(t, expectedErr, err)
}

func TestRemoveCv(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("RemoveCV", seekerID).Return(expectedErr)
	err := srv.RemoveCV(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestRemoveCareer(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("RemoveCareer", seekerID).Return(expectedErr)
	err := srv.RemoveCareer(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestRemoveEducation(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("RemoveEducation", seekerID).Return(expectedErr)
	err := srv.RemoveEducation(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestRemoveLicense(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("RemoveLicense", seekerID).Return(expectedErr)
	err := srv.RemoveLicense(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestRemoveSkill(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("RemoveSkill", seekerID).Return(expectedErr)
	err := srv.RemoveSkill(seekerID)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestUpdateCv(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	input := jobseeker.CVCore{
		ID:          1,
		JobseekerID: 1,
		CV_file:     "A",
	}
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("UpdateCV", input).Return(expectedErr)
	err := srv.UpdateCV(input)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestUpdateCareer(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	input := jobseeker.CareerCore{
		ID:           1,
		JobseekerID:  seekerID,
		Position:     "A",
		Company_name: "B",
		Date_start:   "C",
		Date_end:     "D",
	}

	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("UpdateCareer", seekerID, input).Return(expectedErr)
	err := srv.UpdateCareer(seekerID, input)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestUpdateEducation(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	data := jobseeker.EducationCore{
		ID:              1,
		JobseekerID:     seekerID,
		Education_level: "A",
		Major:           "B",
		Graduation_date: "C",
	}
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("UpdateEducation", seekerID, data).Return(expectedErr)
	err := srv.UpdateEducation(seekerID, data)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

func TestUpdateLicense(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	data := jobseeker.LicenseCore{
		ID:             1,
		JobseekerID:    seekerID,
		License_name:   "A",
		Published_date: "B",
		Expiry_date:    "C",
		License_file:   "D",
	}
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("UpdateLicense", seekerID, data).Return(expectedErr)
	err := srv.UpdateLicense(seekerID, data)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}

// func TestUpdateProfile(t *testing.T) {
// 	repo := new(mocks.JobseekerData)
// 	srv := New(repo)
// 	seekerID := uint(1)
// 	data := jobseeker.JobseekerUpdateCore{
// 		Full_name:  "A",
// 		Username:   "B",
// 		Address:    "C",
// 		Phone:      "D",
// 		Birth_date: "E",
// 		Gender:     "F",
// 		Resume:     "G",
// 		Banners:    "H",
// 	}
// 	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
// 	repo.On("UpdateLicense", seekerID, data).Return(expectedErr)
// 	err := srv.UpdateProfile(seekerID, data)
// 	repo.AssertExpectations(t)
// 	assert.Equal(t, expectedErr, err)
// }

func TestUpdateSkill(t *testing.T) {
	repo := new(mocks.JobseekerData)
	srv := New(repo)
	seekerID := uint(1)
	data := jobseeker.SkillCore{
		ID:          1,
		JobseekerID: seekerID,
		Skill:       "A",
		Description: "B",
	}
	expectedErr := errors.New("Error saat melakukan registrasi jobseeker")
	repo.On("UpdateSkill", seekerID, data).Return(expectedErr)
	err := srv.UpdateSkill(seekerID, data)
	repo.AssertExpectations(t)
	assert.Equal(t, expectedErr, err)
}
