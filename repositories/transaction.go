package repositories

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TransactionRepository interface {
	Insert()(transaction *models.Transaction, err error)
	Update(transactionDetails []models.TransactionDetail) (err error)
	GetTransaction(transactionID string) (*models.Transaction, error)
	GetTransactionList(transactionID *string, status *int) (*[]models.Transaction, error)
	GetTransactionDetailByTransactionID(transactionID string) (*[]models.TransactionDetail, error)
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


func (t *transactionRepository) GetTransaction(transactionID string) (*models.Transaction, error) {
	var transaction *models.Transaction
	res := transactionRepo.db.Scopes(filterTransactionByID(transactionID)).Find(&transaction)
	if res.Error != nil {
		return nil, res.Error
	}

	return transaction, nil
}

func (t *transactionRepository) GetTransactionDetailByTransactionID(transactionID string) (*[]models.TransactionDetail, error) {
	var transactionDetails *[]models.TransactionDetail
	res := transactionRepo.db.Model(models.TransactionDetail{}).Preload("Transaction").Preload("ProductDetail").Scopes(filterTransactionDetailsByTransactionID(transactionID)).Find(&transactionDetails)
	if res.Error != nil {
		return nil, res.Error
	}
	return transactionDetails, nil
}

func (t *transactionRepository) GetTransactionList(transactionID *string, status *int) (*[]models.Transaction, error) {

	scopes := make([]func(db *gorm.DB) *gorm.DB, 0)
	if transactionID != nil {
		scopes = append(scopes, filterTransactionByID(*transactionID))
	}
	if status != nil {
		scopes = append(scopes, filterTransactionByStatus(*status))
	}

	var transactions *[]models.Transaction
	res := transactionRepo.db.Model(models.Transaction{}).Scopes(scopes...).Find(&transactions)
	if res.Error != nil {
		return nil, res.Error
	}

	return transactions, nil
}

func filterTransactionDetailsByTransactionID(transactionID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("transaction_details.transaction_id = ?", transactionID)
}

func filterTransactionByStatus(status int) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("status = ?", status)
}

func filterTransactionByID(transactionID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("id = ?", transactionID)
}