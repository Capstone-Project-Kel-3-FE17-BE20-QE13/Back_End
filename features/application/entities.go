package application

type Core struct {
	JobseekerID uint
	JobId       uint
	Status      string
}

type ApplyDataInterface interface {
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
}

type ApplyServiceInterface interface {
	CreateApplication(input Core) (uint, error)
	GetAllApplications(JobseekerID uint) ([]Core, error)
}
