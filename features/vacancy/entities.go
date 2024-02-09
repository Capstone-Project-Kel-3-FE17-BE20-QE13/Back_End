package vacancy

import (
	"JobHuntz/features/category"
)

type Core struct {
	ID          uint `json:"id"`
	Name        string
	TypeJob     string
	Salary      string
	CategoryId  int
	Category    category.Core
	JobDesc     string
	Requirement string
	CreatedBy   uint
	CompanyId   uint
	// Company      companies.Domain
	// Applications []applications.Domain

}

type JobDataInterface interface {
	//GetCompanyID(ID uint) (uint, error)
	GetJobById(jobID int) (Core, error)
	CreateJob(input Core) error
	// DeleteAllJobs(ctx context.Context) error
	DeleteJobById(input []Core, ID int) error
	GetAllJobs() ([]Core, error)

	//SearchJobs(Name string) ([]Core, error)
	//FilterJobByCategory(ctx context.Context, categoryId int) ([]Core, error)
}

type JobServiceInterface interface {
	//GetCompanyID(ID uint) (uint, error)
	GetJobById(jobID int) (Core, error)
	CreateJob(input Core) error
	// DeleteAllJobs(ctx context.Context) error
	DeleteJobById(input []Core, ID int) error
	GetAllJobs() ([]Core, error)

	//SearchJobs(Name string) ([]Core, error)
	//FilterJobByCategory(ctx context.Context, categoryId int) ([]Job, error)
}
