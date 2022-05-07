package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	userRepository UserRepository
	categoryRepository CategoryRepository
	productUnitRepository ProductUnitRepository
}

func (s *Suite) SetupSuite() {
	var(
		db *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "localhost:5432/postgres",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector)
	require.NoError(s.T(), err)

	maxLifetime := 10 * time.Second
	maxIdle, maxOpenConnection := 5, 5

	s.userRepository = &userRepository{db: &DataSource{s.DB, maxIdle, maxOpenConnection, maxLifetime}}
	s.categoryRepository = &categoryRepository{db: &DataSource{s.DB, maxIdle, maxOpenConnection, maxLifetime}}
	s.productUnitRepository = &productUnitRepository{db: &DataSource{s.DB, maxIdle, maxOpenConnection, maxLifetime}}
}

func TestInit(t *testing.T){
	suite.Run(t, new(Suite))
}
