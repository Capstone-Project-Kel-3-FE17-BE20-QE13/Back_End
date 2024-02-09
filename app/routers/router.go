package router

import (
	_jobseekerData "JobHuntz/features/jobseeker/data"
	_jobseekerHandler "JobHuntz/features/jobseeker/handler"
	_jobseekerService "JobHuntz/features/jobseeker/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	jobseekerData := _jobseekerData.New(db)
	jobseekerService := _jobseekerService.New(jobseekerData)
	jobseekerHandlerAPI := _jobseekerHandler.New(jobseekerService)

	// productData := _dataProduct.New(db)
	// productService := _productService.New(productData)
	// productHandlerAPI := _productHandler.New(productService)

	// StoreRepo := _StoreRepo.New(db)
	// StoreService := _StoreService.New(StoreRepo)
	// StoreHandler := _StoreHandler.New(StoreService)

	// cartData := _CartData.New(db)
	// cartService := _CartService.New(cartData)
	// cartHandlerAPI := _CartHandler.New(cartService)

	// itemData := _ItemData.New(db)
	// itemService := _ItemService.New(itemData)
	// itemHandlerAPI := _ItemHandler.New(itemService)

	// orderData := _OrderData.New(db)
	// orderService := _OrderService.New(orderData)
	// orderHandlerAPI := _OrderHandler.New(orderService)

	// adminRepo := _adminRepo.New(db)
	// adminService := _adminService.New(adminRepo)
	// adminHandlerAPI := _adminHandler.New(adminService)

	// paymentData := _paymentdata.New(db)
	// validate := validator.New()
	// paymentService := _paymentservice.New(paymentData, validate)
	// paymentHandler := _paymenthandler.New(paymentService)

	// user
	e.POST("/jobseekers", jobseekerHandlerAPI.RegisterJobseeker)
	e.POST("/login", jobseekerHandlerAPI.Login)
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
}
