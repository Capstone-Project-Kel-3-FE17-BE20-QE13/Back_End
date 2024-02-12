package service

import (
	"JobHuntz/features/verification"
)

type verificationService struct {
	verificationData verification.VerificationDataInterface
}

// dependency injection
func New(repo verification.VerificationDataInterface) verification.VerificationServiceInterface {
	return &verificationService{
		verificationData: repo,
	}
}

func (service *verificationService) AddOrderJobseeker(input verification.OrderJobseekerCore) error {
	// logic validation
	err := service.verificationData.AddOrderJobseeker(input)
	return err
}

func (service *verificationService) AddOrderCompany(input verification.OrderCompanyCore) error {
	// logic validation
	err := service.verificationData.AddOrderCompany(input)
	return err
}
