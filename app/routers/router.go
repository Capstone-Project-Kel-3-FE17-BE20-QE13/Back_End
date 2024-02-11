package router

import (
	"JobHuntz/app/middlewares"

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

	favorit := _favoritData.New(db)
	favoritService := _favoritService.New(favorit)
	favoritHandlerAPI := _favoritHandler.New(favoritService)

	// authentication
	e.POST("/register/jobseekers", jobseekerHandlerAPI.RegisterJobseeker)
	e.POST("/login/jobseekers", jobseekerHandlerAPI.LoginJobseeker)
	e.POST("/register/company", companyHandlerAPI.RegisterCompany)
	e.POST("/login/company", companyHandlerAPI.LoginCompany)

	// jobseekers
	e.PUT("/jobseekers", jobseekerHandlerAPI.UpdateJobseeker, middlewares.JWTMiddleware())

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

	// vacancy
	e.GET("/vacancy", vacancyHandlerAPI.GetAllJob)
	e.POST("/vacancy", vacancyHandlerAPI.CreateJobs, middlewares.JWTMiddleware())
	e.DELETE("vacancy/jobs_id", vacancyHandlerAPI.Delete, middlewares.JWTMiddleware())

	// application
	e.POST("/application", applicationHandlerAPI.CreateApply, middlewares.JWTMiddleware())
	e.GET("/applications/:userID", applicationHandlerAPI.GetAllApplications, middlewares.JWTMiddleware())

	// favorit
	e.POST("/favorit", favoritHandlerAPI.CreateFavorit, middlewares.JWTMiddleware())
	e.GET("/favorit", favoritHandlerAPI.GetAllFavorit, middlewares.JWTMiddleware())
	e.DELETE("/favorit/favoritID", favoritHandlerAPI.DeleteFavById, middlewares.JWTMiddleware())
}
