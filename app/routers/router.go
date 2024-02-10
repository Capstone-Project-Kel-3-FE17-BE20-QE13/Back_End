package router

import (
	_companyData "JobHuntz/features/company/data"
	_companyHandler "JobHuntz/features/company/handler"
	_companyService "JobHuntz/features/company/service"

	"JobHuntz/app/middlewares"

	_jobseekerData "JobHuntz/features/jobseeker/data"
	_jobseekerHandler "JobHuntz/features/jobseeker/handler"
	_jobseekerService "JobHuntz/features/jobseeker/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_vacancyData "JobHuntz/features/vacancy/data"
	_vacancyHandler "JobHuntz/features/vacancy/handler"
	_vacancyService "JobHuntz/features/vacancy/service"

	_applyData "JobHuntz/features/application/data"
	_applyHandler "JobHuntz/features/application/handler"
	_applyService "JobHuntz/features/application/service"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	jobseekerData := _jobseekerData.New(db)
	jobseekerService := _jobseekerService.New(jobseekerData)
	jobseekerHandlerAPI := _jobseekerHandler.New(jobseekerService)

	company := _companyData.New(db)
	companyService := _companyService.New(company)
	companyHandlerAPI := _companyHandler.New(companyService)

	vacancy := _vacancyData.NewJob(db)
	vacancyService := _vacancyService.NewJob(vacancy)
	vacancyHandlerAPI := _vacancyHandler.NewJob(vacancyService)

	application := _applyData.New(db)
	applicationService := _applyService.New(application)
	applicationHandlerAPI := _applyHandler.New(applicationService)

	// company
	e.POST("/company", companyHandlerAPI.RegisterCompany)

	// authentication
	e.POST("/register/jobseekers", jobseekerHandlerAPI.RegisterJobseeker)
	e.POST("/login/jobseekers", jobseekerHandlerAPI.LoginJobseeker)

	// jobseekers
	e.PUT("/jobseekers", jobseekerHandlerAPI.UpdateJobseeker, middlewares.JWTMiddleware())

	// curriculum vitae
	e.POST("/cv", jobseekerHandlerAPI.CreateCV, middlewares.JWTMiddleware())
	e.GET("/cv", jobseekerHandlerAPI.GetCV, middlewares.JWTMiddleware())

	e.GET("/vacancy", vacancyHandlerAPI.GetAllJob)
	e.POST("/vacancy", vacancyHandlerAPI.CreateJobs, middlewares.JWTMiddleware())
	e.DELETE("vacancy/jobs_id", vacancyHandlerAPI.Delete, middlewares.JWTMiddleware())

	e.POST("/application", applicationHandlerAPI.CreateApply, middlewares.JWTMiddleware())
	e.GET("/applications/:userID", applicationHandlerAPI.GetAllApplications, middlewares.JWTMiddleware())

}
