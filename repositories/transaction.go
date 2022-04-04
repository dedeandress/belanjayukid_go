package repositories

import (
	"belanjayukid_go/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionRepository interface {
	Insert()(transaction *models.Transaction, err error)
	Update(transactionDetails []models.TransactionDetail) (err error)
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
	if err := transactionRepo.db.Create(models.Transaction{ID: uuid.New(), Date: time.Now(), Status: 0, TotalPrice: decimal.NewFromInt(0)}).Scan(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, err
}

func (t *transactionRepository) Update(transactionDetails []models.TransactionDetail) (err error) {
	err = transactionRepo.db.Create(&transactionDetails).Error
	return err
}