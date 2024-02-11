package router

import (
	_companyData "JobHuntz/features/company/data"
	_companyHandler "JobHuntz/features/company/handler"
	_companyService "JobHuntz/features/company/service"
	"JobHuntz/utils/encrypts"
	"JobHuntz/utils/uploads"


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
	hashService := encrypts.NewHashService()
	uploadService := uploads.NewCloudService()

	jobseekerData := _jobseekerData.New(db)
	jobseekerService := _jobseekerService.New(jobseekerData)
	jobseekerHandlerAPI := _jobseekerHandler.New(jobseekerService)

	company := _companyData.New(db, uploadService)
	companyService := _companyService.New(company, hashService)
	companyHandlerAPI := _companyHandler.New(companyService)


	// company
	e.POST("/register/company", companyHandlerAPI.RegisterCompany)
	e.POST("/login/company", companyHandlerAPI.LoginCompany)
	e.GET("/company", companyHandlerAPI.GetById, middlewares.JWTMiddleware())
	e.PUT("/company", companyHandlerAPI.UpdateCompany, middlewares.JWTMiddleware())

	//e.POST("/jobseeker/career", userHandlerAPI.CreateCareer, middlewares.JWTMiddleware())
	// e.GET("/users", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	// e.PUT("/users", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	// e.DELETE("/users", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// // store
	// e.GET("/stores", StoreHandler.GetAllStore, middlewares.JWTMiddleware())
	// // e.GET("/stores/:store_id", StoreHandler.GetStoreById, middlewares.JWTMiddleware()) // error
	// e.POST("/stores", StoreHandler.CreateStore, middlewares.JWTMiddleware())
	// e.PUT("/stores/:store_id", StoreHandler.UpdateStoreById, middlewares.JWTMiddleware())
	// e.DELETE("/stores/:store_id", StoreHandler.DeleteStoreById, middlewares.JWTMiddleware())

	// // product
	// e.POST("/products", productHandlerAPI.CreateProduct, middlewares.JWTMiddleware())
	// e.PUT("/products/:product_id", productHandlerAPI.UpdateProduct, middlewares.JWTMiddleware())
	// e.DELETE("/products/:product_id", productHandlerAPI.Delete, middlewares.JWTMiddleware())
	// e.GET("/all-products", productHandlerAPI.GetAllProducts, middlewares.JWTMiddleware())
	// e.GET("/products/:product_id", productHandlerAPI.GetSingleProduct)
	// e.GET("my-products", productHandlerAPI.GetStoreProduct, middlewares.JWTMiddleware())

	// // shopping cart item
	// e.POST("/shopping-cart", itemHandlerAPI.CreateItem, middlewares.JWTMiddleware())
	// e.PUT("/shopping-cart", itemHandlerAPI.UpdateItem, middlewares.JWTMiddleware())
	// e.DELETE("/shopping-cart", itemHandlerAPI.DeleteItem, middlewares.JWTMiddleware())
	// e.GET("/shopping-cart", itemHandlerAPI.GetItems, middlewares.JWTMiddleware())

	// // order
	// e.POST("/orders", orderHandlerAPI.CreateOrderItem, middlewares.JWTMiddleware())
	// e.GET("/orders", orderHandlerAPI.GetDetailOrder, middlewares.JWTMiddleware())
	// //e.DELETE("/orders-cancel", orderHandlerAPI.CancelOrder, middlewares.JWTMiddleware())
	// e.GET("/orders-history", orderHandlerAPI.OrderHistories, middlewares.JWTMiddleware())

	// // admin
	// e.POST("/admin-login", adminHandlerAPI.Login)
	// e.GET("/admin", adminHandlerAPI.GetAllUser)
	// e.GET("/alluser", userHandlerAPI.GetAllUser)

	// // payment
	// e.POST("/payments", paymentHandler.Payment(), middlewares.JWTMiddleware())
	// e.POST("/payments/callback", paymentHandler.Notification())

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
