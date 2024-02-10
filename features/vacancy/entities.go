package vacancy

type Core struct {
	ID              uint   `json:"id" form:"id"`
	CompanyID       uint   `json:"company_id" form:"company_id"`
	Name            string `json:"name" form:"name"`
	Job_type        string `json:"job_type" form:"job_type"`
	Salary_range    string `json:"salary_range" form:"salary_range"`
	Category        string `json:"category" form:"category"`
	Job_description string `json:"job_desc" form:"job_desc"`
	Job_requirement string `json:"job_req" form:"job_req"`
	Created_by      uint   `json:"created_by" form:"created_by"`
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
