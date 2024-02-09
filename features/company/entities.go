package company

type CompanyCore struct {
	Company_name        string
	Full_name           string
	Email               string
	Password            string
	Address             string
	Phone               string
	Company_type        string
	Company_size        string
	Website             string
	Description         string
	Status_Verification string
	Banners             string
}

type CompanyServiceInterface interface {
	RegisterCompany(input CompanyCore) error
}

type CompanyDataInterface interface {
	RegisterCompany(input CompanyCore) error
}
