package payment

import (
	"JobHuntz/features/verification"
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentCore struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	OrderID     string         `gorm:"type:varchar(50)" json:"order_id" form:"order_id"`
	Amount      string         `json:"amount" form:"amount"`
	UserID      uint           `json:"user_id" form:"user_id"`
	BankAccount string         `gorm:"type:enum('bca', 'bri', 'bni'); default:'bca'"`
	VANumber    string         `gorm:"type:varchar(50)"`
	Status      string         `gorm:"type:varchar(50)"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PaymentHandler interface {
	Payment() echo.HandlerFunc
	Notification() echo.HandlerFunc
}

type PaymentService interface {
	GetOrderJobseekerDetail(dbRaw *sql.DB, userID uint) (verification.OrderJobseekerCore, error)
	GetOrderCompanyDetail(dbRaw *sql.DB, userID uint) (verification.OrderCompanyCore, error)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdateStatus(dbRaw *sql.DB, pay PaymentCore) error
	UpdatePayment(request PaymentCore) error
	CallbackMid(dbRaw *sql.DB, input PaymentCore) error
}

type PaymentData interface {
	GetOrderJobseekerDetail(dbRaw *sql.DB, userID uint) (verification.OrderJobseekerCore, error)
	GetOrderCompanyDetail(dbRaw *sql.DB, userID uint) (verification.OrderCompanyCore, error)
	Payment(request PaymentCore) (PaymentCore, error)
	UpdateStatus(dbRaw *sql.DB, pay PaymentCore) error
	UpdatePayment(request PaymentCore) error
	CallbackMid(dbRaw *sql.DB, input PaymentCore) error
}
