package repositories

import (
	"belanjayukid_go/configs"
	"belanjayukid_go/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"

	_ "gorm.io/driver/postgres"
)

var (
	DB    *DataSource
	DBTrx *DataSource
)

type DataSource struct {
	*gorm.DB
	maxIdleConnection     int
	maxOpenConnection     int
	maxConnectionLifetime time.Duration
}

func InitDBFactory() error {
	url, err := configs.GetConfigRequired(configs.DATABASE_URL)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	sslMode, err := configs.GetConfigRequired(configs.DB_SSL_MODE)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	logModeString, err := configs.GetConfigRequired(configs.DB_LOG_MODE)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	logMode, err := strconv.ParseBool(logModeString)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	if url == "" {
		return fmt.Errorf("Unable to get configuration variable for PostgreSQL, make sure you already set it ")
	}

	DB, err = databaseConnection(url, sslMode, logMode)
	if err != nil {
		return err
	}

	err = migrateDatabase()
	if err != nil {
		return err
	}

	fmt.Println("Database Connection Started")
	return nil
}

func migrateDatabase() error {
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}
	DB.Model(models.User{})
	DB.AutoMigrate(
		models.User{},
		models.Category{},
		models.ProductUnit{},
		models.Product{},
		models.ProductDetail{},
		models.Transaction{},
		models.TransactionDetail{},
	)

	return nil
}

func databaseConnection(url string, sslMode string, logMode bool) (*DataSource, error) {
	dsn := fmt.Sprintf("%s?sslmode=%s", url, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	maxLifetime := 10 * time.Second
	maxIdle, maxOpenConnection := 5, 5
	sqlDB, err :=db.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpenConnection)
	sqlDB.SetConnMaxLifetime(maxLifetime)
	return &DataSource{db, maxIdle, maxOpenConnection, maxLifetime}, nil
}

func BeginTransaction() {
	DBTemp := *DB
	DBTrx = &DataSource{
		DB:                    DBTemp.Begin(),
		maxIdleConnection:     DBTemp.maxIdleConnection,
		maxOpenConnection:     DBTemp.maxOpenConnection,
		maxConnectionLifetime: DBTemp.maxConnectionLifetime,
	}
}

func RollbackTransaction() {
	DBTrx.DB.Rollback()
	DBTrx = nil
}

func CommitTransaction() {
	DBTrx.DB.Commit()
	DBTrx = nil
}
