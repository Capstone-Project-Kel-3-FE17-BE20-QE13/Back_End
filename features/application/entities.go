package application

type Core struct {
	ID                 uint   `json:"id" form:"id"`
	JobseekerID        uint   `json:"jobseeker_id" form:"jobseeker_id"`
	VacancyID          uint   `json:"vacancy_id" form:"vacancy_id"`
	Status_application string `json:"stat_app" form:"stat_app"`
}

type ApplyDataInterface interface {
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
}

type ApplyServiceInterface interface {
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
}
