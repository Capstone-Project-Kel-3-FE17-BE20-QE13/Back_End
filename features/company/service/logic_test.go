package service

import (
	"JobHuntz/features/company"
	"JobHuntz/mocks"
	hashMock "JobHuntz/utils/encrypts/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetId(t *testing.T) {
	repo := new(mocks.CompanyData)
	hash := new(hashMock.HashMock)
	srv := New(repo, hash)
	returnData := &company.CompanyCore{
		ID:                  1,
		Company_name:        "A",
		Full_name:           "B",
		Email:               "C",
		Company_type:        "D",
		Company_size:        "E",
		Website:             "F",
		Description:         "G",
		Status_Verification: "H",
		Banners:             "I",
		Address:             "K",
		Phone:               "K",
	}

	repo.On("GetById", uint(1)).Return(returnData, nil)
	result, err := srv.GetById(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestRegisterCompany(t *testing.T) {
	repo := new(mocks.CompanyData)
	hash := new(hashMock.HashMock)
	inputData := company.CompanyCore{
		Company_name: "Hanapi bagas",
		Full_name:    "Hello",
		Email:        "test@gmail.com",
		Password:     "qwerty",
		Company_type: "IT",
		Company_size: "150",
		Website:      "linku.com",
	}

	t.Run("Success Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("RegisterCompany", mock.Anything).Return(
			func(input company.CompanyCore) (*company.CompanyCore, string, error) {
				return &company.CompanyCore{}, "token", nil
			},
		).Once()
		srv := New(repo, hash)
		res, token, err := srv.RegisterCompany(inputData)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})

	t.Run("Error Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("RegisterCompany", mock.Anything).Return(
			func(input company.CompanyCore) (*company.CompanyCore, string, error) {
				return &company.CompanyCore{}, "token", nil
			},
		).Once()

		srv := New(repo, hash)

		res, token, err := srv.RegisterCompany(inputData)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})
}

func TestLoginCompany(t *testing.T) {
	repo := new(mocks.CompanyData)
	hash := new(hashMock.HashMock)
	email := "alta@mail.com"
	password := "qwerty123"

	t.Run("Login Failure - Email atau password salah", func(t *testing.T) {
		repo.On("LoginCompany", email, password).Return(nil, errors.New("Email atau password salah")).Once()

		srv := New(repo, hash)

		resultUser, token, err := srv.LoginCompany(email, password)

		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, token)
		assert.Equal(t, "Email atau password salah", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Login Failure - Password Mismatch", func(t *testing.T) {
		expectedUser := &company.CompanyCore{
			Email:    email,
			Password: "hashed_password"}

		repo.On("LoginCompany", email, password).Return(expectedUser, nil).Once()
		hash.On("CheckPasswordHash", "hashed_password", password).Return(false).Once()

		srv := New(repo, hash)

		resultUser, token, err := srv.LoginCompany(email, password)

		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, token)
		assert.Equal(t, "password tidak sesuai.", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestUpdateCompany(t *testing.T) {
	repo := new(mocks.CompanyData)
	hash := new(hashMock.HashMock)
	srv := New(repo, hash)
	mockFile := new(multipart.File)

	input := company.CompanyCore{
		ID:           1,
		Company_name: "Data dumy",
		Full_name:    "data dumy",
		Email:        "datadumy@gmail.com",
		Company_type: "IT",
		Company_size: "125",
		Website:      "data umy.com",
		Description:  "data dumy",
		Banners:      "www.image.com/company",
		Address:      "data dumy",
		Phone:        "1234567890",
	}

	t.Run("Succes Update Company", func(t *testing.T) {
		repo.On("UpdateCompany", 1, input, *mockFile, "filename", mock.Anything).Return(nil).Once()
		err := srv.UpdateCompany(1, input, *mockFile, "filename")
		assert.NoError(t, err)
		repo.AssertExpectations(t)

		userIdInvalid := 0
		err = srv.UpdateCompany(userIdInvalid, input, *mockFile, "filename")
		expectedErr := errors.New("invalid id")
		assert.EqualError(t, err, expectedErr.Error())
	})
}
