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

func (service *verificationService) AddOrder(input verification.OrderCore) error {
	// logic validation
	err := service.verificationData.AddOrder(input)
	return err
}
