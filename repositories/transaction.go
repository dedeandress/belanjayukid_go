package repositories

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
	"time"
)

type TransactionRepository interface {
	Insert()(transaction *models.Transaction, err error)
	Update(transactionDetails []models.TransactionDetail) (err error)
	UpdateTrxStatus(transactionID string, status int) (err error)
	UpdateTrxTotalPrice(transactionID string, totalPrice decimal.Decimal) (err error)
}

type transactionRepository struct {
	db *DataSource
}

var transactionRepo *transactionRepository

func GetTransactionRepository() TransactionRepository {
	if DBTrx != nil {
		transactionRepo = &transactionRepository{db: DBTrx}
	} else {
		transactionRepo = &transactionRepository{db: DB}
	}

	return transactionRepo
}

func (t *transactionRepository) Insert() (transaction *models.Transaction, err error)  {
	transaction = &models.Transaction{}
	if err := transactionRepo.db.Create(models.Transaction{ID: uuid.New(), Date: time.Now(), Status: enums.INIT_TRANSACTION, TotalPrice: decimal.NewFromInt(0)}).Scan(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, err
}

func (t *transactionRepository) Update(transactionDetails []models.TransactionDetail) (err error) {
	err = transactionRepo.db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"number_of_purchases"}),
		}).Create(&transactionDetails).Error
	return err
}


func (t *transactionRepository) UpdateTrxStatus(transactionID string, status int) (err error) {
	err = transactionRepo.db.Model(models.Transaction{}).Where("id = ?", transactionID).Update("status", status).Error
	return err
}


func (t *transactionRepository) UpdateTrxTotalPrice(transactionID string, totalPrice decimal.Decimal) (err error) {
	err = transactionRepo.db.Model(models.Transaction{}).Where("id = ?", transactionID).Update("total_price", totalPrice).Error
	return err
}