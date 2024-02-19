package service_test

import (
	"JobHuntz/features/payment"
	"JobHuntz/features/payment/service"
	"JobHuntz/features/verification"
	"JobHuntz/mocks"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPaymentService_GetOrderJobseekerDetail(t *testing.T) {
	mockRepo := &mocks.PaymentData{}
	service := service.New(mockRepo, nil)

	expectedData := verification.OrderJobseekerCore{}
	mockRepo.On("GetOrderJobseekerDetail", mock.Anything, uint(1)).Return(expectedData, nil)

	data, err := service.GetOrderJobseekerDetail(&sql.DB{}, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
	mockRepo.AssertExpectations(t)
}

func TestPaymentService_GetOrderCompanyDetail(t *testing.T) {
	mockRepo := &mocks.PaymentData{}
	service := service.New(mockRepo, nil)

	expectedData := verification.OrderCompanyCore{}
	mockRepo.On("GetOrderCompanyDetail", mock.Anything, uint(1)).Return(expectedData, nil)

	data, err := service.GetOrderCompanyDetail(&sql.DB{}, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
	mockRepo.AssertExpectations(t)
}

// Similarly, write tests for other methods...

// func TestPaymentService_Payment_ValidInput(t *testing.T) {
// 	mockRepo := &mocks.PaymentData{}
// 	service := service.New(mockRepo, nil)

// 	input := payment.PaymentCore{}
// 	mockRepo.On("Payment", input).Return(input, nil)

// 	result, err := service.Payment(input)

// 	assert.NoError(t, err)
// 	assert.Equal(t, input, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestPaymentService_Payment_InvalidInput(t *testing.T) {
// 	mockRepo := &mocks.PaymentData{}
// 	service := service.New(mockRepo, nil)

// 	input := payment.PaymentCore{}
// 	mockRepo.On("Payment", input).Return(payment.PaymentCore{}, errors.New("error"))

// 	result, err := service.Payment(input)

// 	assert.Error(t, err)
// 	assert.Equal(t, payment.PaymentCore{}, result)
// 	mockRepo.AssertExpectations(t)
// }

func TestPaymentService_UpdateStatus(t *testing.T) {
	mockRepo := &mocks.PaymentData{}
	service := service.New(mockRepo, nil)

	input := payment.PaymentCore{}
	mockRepo.On("UpdateStatus", mock.Anything, input).Return(nil)

	err := service.UpdateStatus(&sql.DB{}, input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// func TestPaymentService_CallbackMid_ValidInput(t *testing.T) {
// 	mockRepo := &mocks.PaymentData{}
// 	service := service.New(mockRepo, nil)

// 	input := payment.PaymentCore{OrderID: "123"}
// 	mockRepo.On("CallbackMid", mock.Anything, input).Return(nil)

// 	err := service.CallbackMid(&sql.DB{}, input)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestPaymentService_CallbackMid_InvalidInput(t *testing.T) {
// 	mockRepo := &mocks.PaymentData{}
// 	service := service.New(mockRepo, nil)

// 	input := payment.PaymentCore{OrderID: ""}
// 	mockRepo.On("CallbackMid", mock.Anything, input).Return(errors.New("error"))

// 	err := service.CallbackMid(&sql.DB{}, input)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

func TestPaymentService_UpdatePayment_ValidInput(t *testing.T) {
	mockRepo := &mocks.PaymentData{}
	service := service.New(mockRepo, nil)

	input := payment.PaymentCore{}
	mockRepo.On("UpdatePayment", input).Return(nil)

	err := service.UpdatePayment(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPaymentService_UpdatePayment_InvalidInput(t *testing.T) {
	mockRepo := &mocks.PaymentData{}
	service := service.New(mockRepo, nil)

	input := payment.PaymentCore{}
	mockRepo.On("UpdatePayment", input).Return(errors.New("error"))

	err := service.UpdatePayment(input)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
