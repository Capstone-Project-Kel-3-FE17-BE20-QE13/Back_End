package service

import (
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
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

func (service *verificationService) GetDataJobseeker(userID uint) (jobseeker.JobseekerCore, error) {
	// logic validation
	res, err := service.verificationData.GetDataJobseeker(userID)
	return res, err
}

func (service *verificationService) GetDataCompany(userID uint) (company.CompanyCore, error) {
	// logic validation
	res, err := service.verificationData.GetDataCompany(userID)
	return res, err
}

func (service *verificationService) AddOrder(input verification.OrderCore) error {
	// logic validation
	err := service.verificationData.AddOrder(input)
	return err
}

// func (service *verificationService) AddOrderCompany(input verification.OrderCore) error {
// 	// logic validation
// 	err := service.verificationData.AddOrderCompany(input)
// 	return err
// }
