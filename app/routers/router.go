package router

import (
	"JobHuntz/utils/encrypts"
	"JobHuntz/utils/uploads"

	"JobHuntz/app/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_jobseekerData "JobHuntz/features/jobseeker/data"
	_jobseekerHandler "JobHuntz/features/jobseeker/handler"
	_jobseekerService "JobHuntz/features/jobseeker/service"

	_companyData "JobHuntz/features/company/data"
	_companyHandler "JobHuntz/features/company/handler"
	_companyService "JobHuntz/features/company/service"

	_vacancyData "JobHuntz/features/vacancy/data"
	_vacancyHandler "JobHuntz/features/vacancy/handler"
	_vacancyService "JobHuntz/features/vacancy/service"

	_applyData "JobHuntz/features/application/data"
	_applyHandler "JobHuntz/features/application/handler"
	_applyService "JobHuntz/features/application/service"

	_favoritData "JobHuntz/features/favorit/data"
	_favoritHandler "JobHuntz/features/favorit/handler"
	_favoritService "JobHuntz/features/favorit/service"

	_verifData "JobHuntz/features/verification/data"
	_verifHandler "JobHuntz/features/verification/handler"
	_verifService "JobHuntz/features/verification/service"

	_paymentdata "JobHuntz/features/payment/data"
	_paymenthandler "JobHuntz/features/payment/handler"
	_paymentservice "JobHuntz/features/payment/service"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()
	uploadService := uploads.NewCloudService()

	jobseekerData := _jobseekerData.New(db)
	jobseekerService := _jobseekerService.New(jobseekerData)
	jobseekerHandlerAPI := _jobseekerHandler.New(jobseekerService)

	company := _companyData.New(db, uploadService)
	companyService := _companyService.New(company, hashService)
	companyHandlerAPI := _companyHandler.New(companyService)

	vacancy := _vacancyData.NewJob(db)
	vacancyService := _vacancyService.NewJob(vacancy)
	vacancyHandlerAPI := _vacancyHandler.NewJob(vacancyService)

	application := _applyData.New(db)
	applicationService := _applyService.New(application)
	applicationHandlerAPI := _applyHandler.New(applicationService)

	favorit := _favoritData.New(db)
	favoritService := _favoritService.New(favorit)
	favoritHandlerAPI := _favoritHandler.New(favoritService)

	verif := _verifData.New(db)
	verifService := _verifService.New(verif)
	verifHandlerAPI := _verifHandler.New(verifService)

	paymentData := _paymentdata.New(db)
	validate := validator.New()
	paymentService := _paymentservice.New(paymentData, validate)
	paymentHandler := _paymenthandler.New(paymentService)

	// company
	e.POST("/register/company", companyHandlerAPI.RegisterCompany)
	e.POST("/login/company", companyHandlerAPI.LoginCompany)
	e.GET("/company", companyHandlerAPI.GetById, middlewares.JWTMiddleware())
	e.PUT("/company", companyHandlerAPI.UpdateCompany, middlewares.JWTMiddleware())
	e.GET("/company-getjobseeker", jobseekerHandlerAPI.GetjobseekerByCompany, middlewares.JWTMiddleware())

	// jobseekers
	e.POST("/register/jobseekers", jobseekerHandlerAPI.RegisterJobseeker)
	e.POST("/login/jobseekers", jobseekerHandlerAPI.LoginJobseeker)
	e.PUT("/jobseekers", jobseekerHandlerAPI.UpdateJobseeker, middlewares.JWTMiddleware())
	e.GET("/jobseekers", jobseekerHandlerAPI.GetByIdJobSeeker, middlewares.JWTMiddleware())

	// curriculum vitae
	e.POST("/cv", jobseekerHandlerAPI.CreateCV, middlewares.JWTMiddleware())
	e.GET("/cv", jobseekerHandlerAPI.GetCV, middlewares.JWTMiddleware())
	e.PUT("/cv", jobseekerHandlerAPI.UpdateCV, middlewares.JWTMiddleware())
	e.DELETE("/cv", jobseekerHandlerAPI.DeleteCV, middlewares.JWTMiddleware())

	// career
	e.POST("/career", jobseekerHandlerAPI.CreateCareer, middlewares.JWTMiddleware())
	e.GET("/career/:career_id", jobseekerHandlerAPI.GetSingleCareer, middlewares.JWTMiddleware())
	e.GET("/all-careers", jobseekerHandlerAPI.GetAllCareers, middlewares.JWTMiddleware())
	e.PUT("/career/:career_id", jobseekerHandlerAPI.UpdateCareer, middlewares.JWTMiddleware())
	e.DELETE("/career/:career_id", jobseekerHandlerAPI.DeleteCareer, middlewares.JWTMiddleware())

	// education
	e.POST("/education", jobseekerHandlerAPI.CreateEducation, middlewares.JWTMiddleware())
	e.GET("/education/:education_id", jobseekerHandlerAPI.GetSingleEducation, middlewares.JWTMiddleware())
	e.GET("/all-educations", jobseekerHandlerAPI.GetAllEducations, middlewares.JWTMiddleware())
	e.PUT("/education/:education_id", jobseekerHandlerAPI.UpdateEducation, middlewares.JWTMiddleware())
	e.DELETE("/education/:education_id", jobseekerHandlerAPI.DeleteEducation, middlewares.JWTMiddleware())

	// license
	e.POST("/license", jobseekerHandlerAPI.CreateLicense, middlewares.JWTMiddleware())
	e.GET("/license/:license_id", jobseekerHandlerAPI.GetSingleLicense, middlewares.JWTMiddleware())
	e.GET("/all-licenses", jobseekerHandlerAPI.GetAllLicenses, middlewares.JWTMiddleware())
	e.PUT("/license/:license_id", jobseekerHandlerAPI.UpdateLicense, middlewares.JWTMiddleware())
	e.DELETE("/license/:license_id", jobseekerHandlerAPI.DeleteLicense, middlewares.JWTMiddleware())

	// Skill
	e.POST("/skill", jobseekerHandlerAPI.CreateSkill, middlewares.JWTMiddleware())
	e.GET("/skill/:skill_id", jobseekerHandlerAPI.GetSingleSkill, middlewares.JWTMiddleware())
	e.GET("/all-skills", jobseekerHandlerAPI.GetAllSkills, middlewares.JWTMiddleware())
	e.PUT("/skill/:skill_id", jobseekerHandlerAPI.UpdateSkill, middlewares.JWTMiddleware())
	e.DELETE("/skill/:skill_id", jobseekerHandlerAPI.DeleteSkill, middlewares.JWTMiddleware())

	// vacancy
	e.GET("/all-vacancies", vacancyHandlerAPI.GetAllJob)
	e.GET("/vacancy/:vacancy_id", vacancyHandlerAPI.GetJobById, middlewares.JWTMiddleware())
	e.POST("/vacancy", vacancyHandlerAPI.CreateJobs, middlewares.JWTMiddleware())
	e.GET("/mycompany-vacancies", vacancyHandlerAPI.GetVacanciesMadeByCompany, middlewares.JWTMiddleware())
	e.DELETE("vacancy/:vacancy_id", vacancyHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/vacancy/:vacancy_id", vacancyHandlerAPI.UpdateVacancyStatus, middlewares.JWTMiddleware())

	// application
	e.POST("/application", applicationHandlerAPI.CreateApply, middlewares.JWTMiddleware())
	e.GET("/applications-jobseeker", applicationHandlerAPI.AppHistoryJobseeker, middlewares.JWTMiddleware())
	e.GET("/applications-company", applicationHandlerAPI.AppHistoryCompany, middlewares.JWTMiddleware())
	e.PUT("/application/:id", applicationHandlerAPI.EditApplicationStatus, middlewares.JWTMiddleware())

	// favorit
	e.POST("/favorit", favoritHandlerAPI.CreateFavorit, middlewares.JWTMiddleware())
	e.GET("/favorit", favoritHandlerAPI.GetAllFavorit, middlewares.JWTMiddleware())
	e.DELETE("/favorit/:favorit_id", favoritHandlerAPI.DeleteFavById, middlewares.JWTMiddleware())

	// order
	e.POST("/order/jobseeker", verifHandlerAPI.CreateOrderJobseeker, middlewares.JWTMiddleware())
	e.POST("/order-company", verifHandlerAPI.CreateOrderCompany, middlewares.JWTMiddleware())

	// payment
	e.POST("/payments", paymentHandler.Payment(), middlewares.JWTMiddleware())
	e.POST("/payments/callback", paymentHandler.Notification())
}
