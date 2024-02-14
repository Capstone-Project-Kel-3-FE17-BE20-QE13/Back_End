package data

import (
	"database/sql"
	"errors"
	"fmt"

	"JobHuntz/app/database"
	"JobHuntz/app/middlewares"
	"JobHuntz/features/company"
	"JobHuntz/features/jobseeker"
	"JobHuntz/features/payment"
	"JobHuntz/features/verification"

	"gorm.io/gorm"
)

var log = middlewares.Log()

type paymentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.PaymentData {
	return &paymentQuery{
		db: db,
	}
}

func (pq *paymentQuery) GetOrderDetail(dbRaw *sql.DB, userID uint) (verification.OrderCore, error) {
	// var order_id string
	// var price float64

	var dataOrder verification.OrderCore

	query := `SELECT * FROM orders WHERE orders.company_id = ? OR orders.jobseeker_id = ?`

	rowID := dbRaw.QueryRow(query, userID, userID)

	if err := rowID.Scan(&dataOrder.ID, &dataOrder.JobseekerID, &dataOrder.CompanyID, &dataOrder.Price, &dataOrder.Status_order); err != nil {
		if err == sql.ErrNoRows {
			// Handle jika tidak ada data ditemukan
			return verification.OrderCore{}, errors.New("no data found")
		}
		log.Fatal("cannot scan data: ")
	}
	return dataOrder, nil
}

func (pq *paymentQuery) Payment(request payment.PaymentCore) (payment.PaymentCore, error) {
	paymentData := chargeMidtrans(request)

	result := pq.db.Create(&paymentData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("payment record not found")
		return payment.PaymentCore{}, errors.New("payment record not found")
	}

	if result.RowsAffected == 0 {
		log.Warn("no charge payment has been created")
		return payment.PaymentCore{}, errors.New("row affected : 0")
	}

	if result.Error != nil {
		log.Error("error while charging payment")
		return payment.PaymentCore{}, errors.New("internal server error")
	}
	fmt.Printf("log ppayment data : %v\n", paymentData)
	fmt.Printf("log ppayment model: %v\n", paymentModels(paymentData))
	return paymentModels(paymentData), nil
}

func (pq *paymentQuery) UpdateStatus(dbRaw *sql.DB, pay payment.PaymentCore) error {
	// Buat pernyataan SQL UPDATE
	query := "UPDATE payments SET status = ? WHERE transaction_id = ?"

	// Eksekusi pernyataan SQL
	_, err := dbRaw.Exec(query, pay.Status, pay.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update implements user.UserDataInterface.
func (pq *paymentQuery) CallbackMid(dbRaw *sql.DB, input payment.PaymentCore) error {
	dataGorm := CoreToModel(input)
	tx := pq.db.Model(&database.Payment{}).Where("order_id = ?", input.OrderID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	// Jika tidak ada data yang diupdate, kembalikan error
	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	} else if tx.RowsAffected != 0 {
		// Jika ada data yang diupdate, lanjutkan proses
		fmt.Println("row affected: ", tx.RowsAffected)

		// data jobseeker
		query1 := `select jobseekers.id from jobseekers
		join orders on jobseekers.id = orders.jobseeker_id
		where orders.id = ?`

		rowID1 := dbRaw.QueryRow(query1, input.OrderID)

		var dataJobseeker jobseeker.JobseekerCore
		if err := rowID1.Scan(&dataJobseeker.ID); err != nil {
			if err == sql.ErrNoRows {
				//
			}
		}

		fmt.Println("isi data jobseeker: ", dataJobseeker)

		// data company
		query2 := `select companies.id from companies
		join orders on companies.id = orders.company_id
		where orders.id = ?`

		rowID2 := dbRaw.QueryRow(query2, input.OrderID)

		var dataCompany company.CompanyCore
		if err := rowID2.Scan(&dataCompany.ID); err != nil {
			if err == sql.ErrNoRows {
				//
			}
		}

		fmt.Println("isi data company: ", dataCompany)

		// -------------------------------------
		if dataJobseeker.ID != 0 {
			query1 := "UPDATE jobseekers SET status_verification = 'Verified' WHERE id = ?"
			_, err := dbRaw.Exec(query1, dataJobseeker.ID)
			if err != nil {
				return err
			}
		} else {
			query2 := "UPDATE companies SET status_verification = 'Verified' WHERE id = ?"
			_, err := dbRaw.Exec(query2, dataCompany.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// UpdatePayment implements payment.PaymentData
func (pq *paymentQuery) UpdatePayment(request payment.PaymentCore) error {
	req := paymentEntities(request)
	log.Sugar().Infof("callback midtrans status: %s, order ID: %s, transaction ID: %s",
		req.Status, req.OrderID, req.ID)
	query := pq.db.Table("payments").
		Where("id = ? AND order_id = ?", request.ID, request.OrderID).
		Updates(map[string]interface{}{
			"status": request.Status,
		})
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return errors.New("user profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no payment record has been updated")
		return errors.New("no payment record has been updated")
	}

	if query.Error != nil {
		log.Error("error while updating payment status")
		return errors.New("internal server error")
	}

	return nil
}
