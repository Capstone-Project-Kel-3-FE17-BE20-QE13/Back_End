package handler

import (
	"JobHuntz/features/jobseeker"
)

type LoginResponse struct {
	ID        uint   `json:"id" form:"id"`
	Full_name string `gorm:"not null" json:"full_name" form:"full_name"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Token     string `json:"token" form:"token"`
	Roles     string `json:"roles" form:"roles"`
}

type JobseekerResponse struct {
	ID                  uint   `json:"id" form:"id"`
	Full_name           string `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string `gorm:"not null" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	Password            string `gorm:"not null" json:"password" form:"password"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Birth_date          string `json:"birth_date" form:"birth_date"`
	Gender              string `json:"gender" form:"gender"`
	Resume              string `json:"resume" form:"resume"`
	Status_Verification string `json:"stat_verif" form:"stat_verif"`
	Banners             string `json:"banners" form:"banners"`
	Careers             []CareerResponse
}

type JobseekerResponseById struct {
	ID                  uint   `json:"id" form:"id"`
	Full_name           string `gorm:"not null" json:"full_name" form:"full_name"`
	Username            string `gorm:"not null" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	Address             string `json:"address" form:"address"`
	Phone               string `json:"phone" form:"phone"`
	Birth_date          string `json:"birth_date" form:"birth_date"`
	Gender              string `json:"gender" form:"gender"`
	Resume              string `json:"resume" form:"resume"`
	Status_Verification string `json:"stat_verif" form:"stat_verif"`
	Banners             string `json:"banners" form:"banners"`
	Careers             []CareerResponse
	Educations          []EducationResponse
	Cvs                 []CvResponse
	Licenses            []LicenseResponse
	Skills              []SkillResponse
}

type CvResponse struct {
	ID          uint   `json:"id" form:"id"`
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	CV_file     string `json:"cv_file" form:"cv_file"`
}

type CareerResponse struct {
	ID           uint   `json:"id" form:"id"`
	JobseekerID  uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Position     string `json:"position" form:"position"`
	Company_name string `json:"company_name" form:"company_name"`
	Date_start   string `json:"date_start" form:"date_start"`
	Date_end     string `json:"date_end" form:"date_end"`
}

type EducationResponse struct {
	ID              uint   `json:"id" form:"id"`
	JobseekerID     uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Education_level string `json:"ed_level" form:"ed_level"`
	Major           string `json:"major" form:"major"`
	Graduation_date string `json:"grad_date" form:"grad_date"`
}

type LicenseResponse struct {
	ID             uint   `json:"id" form:"id"`
	JobseekerID    uint   `json:"jobseeker_id" form:"jobseeker_id"`
	License_name   string `json:"license_name" form:"license_name"`
	Published_date string `json:"pub_date" form:"pub_date"`
	Expiry_date    string `json:"exp_date" form:"exp_date"`
	License_file   string `json:"license" form:"license"`
}

type SkillResponse struct {
	ID          uint   `json:"id" form:"id"`
	JobseekerID uint   `json:"jobseeker_id" form:"jobseeker_id"`
	Skill       string `json:"skill" form:"skill"`
	Description string `json:"description" form:"description"`
}

func CoreResponById(data jobseeker.JobseekerCore) JobseekerResponseById {
	career := make([]CareerResponse, len(data.Careers))
	educations := make([]EducationResponse, len(data.Educations))
	cvs := make([]CvResponse, len(data.Cvs))
	licenses := make([]LicenseResponse, len(data.Licenses))
	skils := make([]SkillResponse, len(data.Skills))

	for i, skil := range data.Skills {
		skils[i] = SkillResponse{
			ID:          skil.ID,
			JobseekerID: skil.JobseekerID,
			Skill:       skil.Skill,
			Description: skil.Description,
		}
	}

	for i, license := range data.Licenses {
		licenses[i] = LicenseResponse{
			ID:             license.ID,
			JobseekerID:    license.JobseekerID,
			License_name:   license.License_name,
			Published_date: license.Published_date,
			Expiry_date:    license.Expiry_date,
			License_file:   license.License_file,
		}
	}

	for i, cv := range data.Cvs {
		cvs[i] = CvResponse{
			JobseekerID: cv.JobseekerID,
			CV_file:     cv.CV_file,
		}
	}

	for i, education := range data.Educations {
		educations[i] = EducationResponse{
			ID:              education.ID,
			JobseekerID:     education.JobseekerID,
			Education_level: education.Education_level,
			Major:           education.Major,
			Graduation_date: education.Graduation_date,
		}
	}

	for i, carer := range data.Careers {
		career[i] = CareerResponse{
			ID:           carer.ID,
			JobseekerID:  carer.JobseekerID,
			Position:     carer.Position,
			Company_name: carer.Company_name,
			Date_start:   carer.Date_start,
			Date_end:     carer.Date_end,
		}
	}

	return JobseekerResponseById{
		ID:                  data.ID,
		Full_name:           data.Full_name,
		Username:            data.Username,
		Email:               data.Email,
		Address:             data.Address,
		Phone:               data.Phone,
		Birth_date:          data.Birth_date,
		Gender:              data.Gender,
		Resume:              data.Resume,
		Status_Verification: data.Status_Verification,
		Banners:             data.Banners,
		Careers:             career,
		Educations:          educations,
		Cvs:                 cvs,
		Licenses:            licenses,
		Skills:              skils,
	}
}

func CoreJobseekerToResponse(input jobseeker.JobseekerCore) JobseekerResponse {
	return JobseekerResponse{
		ID:                  input.ID,
		Full_name:           input.Full_name,
		Username:            input.Username,
		Email:               input.Email,
		Password:            input.Password,
		Address:             input.Address,
		Phone:               input.Phone,
		Birth_date:          input.Birth_date,
		Gender:              input.Gender,
		Resume:              input.Resume,
		Status_Verification: input.Status_Verification,
		Banners:             input.Banners,
	}
}

func CoreJobseekerToResponseLogin(input jobseeker.JobseekerCore, token string) LoginResponse {
	return LoginResponse{
		ID:        input.ID,
		Full_name: input.Full_name,
		Email:     input.Email,
		Roles:     "jobseeker",
		Token:     token,
	}
}

func CoreCareerToResponse(input jobseeker.CareerCore) CareerResponse {
	return CareerResponse{
		ID:           input.ID,
		JobseekerID:  input.JobseekerID,
		Position:     input.Position,
		Company_name: input.Company_name,
		Date_start:   input.Date_start,
		Date_end:     input.Date_end,
	}
}

func CoreCareersToResponse(input []jobseeker.CareerCore) []CareerResponse {
	var careerResponses []CareerResponse
	for _, v := range input {
		var careerInput = CareerResponse{
			ID:           v.ID,
			JobseekerID:  v.JobseekerID,
			Position:     v.Position,
			Company_name: v.Company_name,
			Date_start:   v.Date_start,
			Date_end:     v.Date_end,
		}
		careerResponses = append(careerResponses, careerInput)
	}

	return careerResponses
}

func CoreEducationToResponse(input jobseeker.EducationCore) EducationResponse {
	return EducationResponse{
		ID:              input.ID,
		JobseekerID:     input.JobseekerID,
		Education_level: input.Education_level,
		Major:           input.Major,
		Graduation_date: input.Graduation_date,
	}
}

func CoreEdusToResponse(input []jobseeker.EducationCore) []EducationResponse {
	var eduResponses []EducationResponse
	for _, input := range input {
		var eduInput = EducationResponse{
			ID:              input.ID,
			JobseekerID:     input.JobseekerID,
			Education_level: input.Education_level,
			Major:           input.Major,
			Graduation_date: input.Graduation_date,
		}
		eduResponses = append(eduResponses, eduInput)
	}

	return eduResponses
}

func CoreLicenseToResponse(input jobseeker.LicenseCore) LicenseResponse {
	return LicenseResponse{
		ID:             input.ID,
		JobseekerID:    input.JobseekerID,
		License_name:   input.License_name,
		Published_date: input.Published_date,
		Expiry_date:    input.Expiry_date,
		License_file:   input.License_file,
	}
}

func CoreLicensesToResponse(input []jobseeker.LicenseCore) []LicenseResponse {
	var licenseResponses []LicenseResponse
	for _, input := range input {
		var licenseInput = LicenseResponse{
			ID:             input.ID,
			JobseekerID:    input.JobseekerID,
			License_name:   input.License_name,
			Published_date: input.Published_date,
			Expiry_date:    input.Expiry_date,
			License_file:   input.License_file,
		}
		licenseResponses = append(licenseResponses, licenseInput)
	}

	return licenseResponses
}

func CoreSkillToResponse(input jobseeker.SkillCore) SkillResponse {
	return SkillResponse{
		ID:          input.ID,
		JobseekerID: input.JobseekerID,
		Skill:       input.Skill,
		Description: input.Description,
	}
}

func CoreSkillsToResponse(data []jobseeker.SkillCore) []SkillResponse {
	var skillsResponse []SkillResponse
	for _, input := range data {
		var skillInput = SkillResponse{
			ID:          input.ID,
			JobseekerID: input.JobseekerID,
			Skill:       input.Skill,
			Description: input.Description,
		}
		skillsResponse = append(skillsResponse, skillInput)
	}

	return skillsResponse
}
