package service_test

import (
	"JobHuntz/features/favorit"
	"JobHuntz/features/favorit/service"
	"JobHuntz/mocks"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFavService_GetDataCompany(t *testing.T) {
	// Membuat objek mock untuk FavDataInterface
	mockRepo := new(mocks.FavoritData)

	// Mengatur perilaku yang diharapkan dari mock
	expectedData := favorit.DataCompanyCore{Position: "Mangaka", Company_name: "Test Company"}
	mockRepo.On("GetDataCompany", mock.Anything, uint(1)).Return(expectedData, nil)

	// Membuat objek FavService dengan menggunakan mock
	favService := service.New(mockRepo)

	// Memanggil method GetDataCompany pada FavService
	db := &sql.DB{} // Placeholder for a real database connection
	actualData, err := favService.GetDataCompany(db, 1)

	// Memeriksa bahwa panggilan ke GetDataCompany dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)
}

func TestFavService_CreateFavorit(t *testing.T) {
	// Membuat objek mock untuk FavDataInterface
	mockRepo := new(mocks.FavoritData)

	// Mengatur perilaku yang diharapkan dari mock
	expectedID := uint(1)
	mockRepo.On("CreateFavorit", mock.Anything).Return(expectedID, nil)

	// Membuat objek FavService dengan menggunakan mock
	favService := service.New(mockRepo)

	// Menyiapkan data input untuk CreateFavorit
	input := favorit.Core{ID: 1, JobseekerID: 1, VacancyID: 2, Position: "pelaut", Company_name: "Perusahaan Laut"}

	// Memanggil method CreateFavorit pada FavService
	actualID, err := favService.CreateFavorit(input)

	// Memeriksa bahwa panggilan ke CreateFavorit dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
	assert.NoError(t, err)
	assert.Equal(t, expectedID, actualID)
}

func TestFavService_DeleteFavById(t *testing.T) {
	// Membuat objek mock untuk FavDataInterface
	mockRepo := new(mocks.FavoritData)

	// Mengatur perilaku yang diharapkan dari mock
	mockRepo.On("DeleteFavById", mock.Anything, 1).Return(nil)

	// Membuat objek FavService dengan menggunakan mock
	favService := service.New(mockRepo)

	// Menyiapkan data input untuk DeleteFavById
	input := []favorit.Core{{ID: 1, JobseekerID: 1, VacancyID: 2, Position: "pelaut", Company_name: "Perusahaan Laut"}}

	// Memanggil method DeleteFavById pada FavService
	err := favService.DeleteFavById(input, 1)

	// Memeriksa bahwa panggilan ke DeleteFavById dilakukan dengan benar
	mockRepo.AssertExpectations(t)

	// Memastikan tidak ada error
	assert.NoError(t, err)
}

// func TestFavService_GetAllFavorit(t *testing.T) {
// 	// Membuat objek mock untuk FavDataInterface
// 	mockRepo := new(mocks.FavoritData)

// 	userID := userID{id}
// 	// Mengatur perilaku yang diharapkan dari mock
// 	expectedData := []favorit.Core{
// 		{ID: 1, JobseekerID: 1, VacancyID: 2, Position: "pelaut", Company_name: "Perusahaan Laut"},
// 		{ID: 2, JobseekerID: 1, VacancyID: 2, Position: "pemancing", Company_name: "Perusahaan Laut"},
// 	}
// 	mockRepo.On("GetAllFavorit").Return(expectedData, nil)

// 	// Membuat objek FavService dengan menggunakan mock
// 	favService := service.New(mockRepo)

// 	// Memanggil method GetAllFavorit pada FavService
// 	actualData, err := favService.GetAllFavorit(userID)

// 	// Memeriksa bahwa panggilan ke GetAllFavorit dilakukan dengan benar
// 	mockRepo.AssertExpectations(t)

// 	// Memastikan tidak ada error dan hasil yang diharapkan sesuai
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedData, actualData)
// }
