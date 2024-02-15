package vacancy

type Core struct {
	ID              uint   `json:"id" form:"id"`
	CompanyID       uint   `json:"company_id" form:"company_id"`
	Name            string `json:"name" form:"name"`
	Job_type        string `json:"job_type" form:"job_type"`
	Salary_range    string `json:"salary_range" form:"salary_range"`
	Category        string `json:"category" form:"category"`
	Address         string `json:"address" form:"address"`
	Job_description string `json:"job_desc" form:"job_desc"`
	Job_requirement string `json:"job_req" form:"job_req"`
	Status          string `json:"status" form:"status"`
}

type CompanyCore struct {
	ID                  uint
	Company_name        string `validate:"required"`
	Full_name           string `validate:"required"`
	Email               string `validate:"required"`
	Password            string `validate:"required"`
	Company_type        string `validate:"required"`
	Company_size        string `validate:"required"`
	Website             string `validate:"required"`
	Description         string
	Status_Verification string
	Banners             string
	Address             string
	Phone               string
	Token               string
}

type JobStatusCore struct {
	Status string `json:"status" form:"status"`
}

type JobDataInterface interface {
	GetJobById(jobID int) (Core, error)
	CreateJob(input Core) error
	DeleteJobById(input []Core, ID int) error
	MyCompanyVacancies(companyID uint) ([]Core, error)
	GetAllJobs() ([]Core, error)
	GetById(id uint) (*CompanyCore, error)
	CountJobsByUserID(userID uint) (int, error)
	UpdateStatus(input JobStatusCore, vacancyID uint) error
}

type JobServiceInterface interface {
	GetJobById(jobID int) (Core, error)
	CreateJob(input Core) error
	DeleteJobById(input []Core, ID int) error
	MyCompanyVacancies(companyID uint) ([]Core, error)
	GetAllJobs() ([]Core, error)
	GetById(id uint) (*CompanyCore, error)
	CountJobsByUserID(userID uint) (int, error)
	UpdateStatus(input JobStatusCore, vacancyID uint) error
}
