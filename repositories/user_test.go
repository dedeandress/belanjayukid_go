package repositories

import (
	"belanjayukid_go/models"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"regexp"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repository UserRepository
	user *models.User
}

func (s *Suite) SetupSuite() {
	var(
		db *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dsn := url.URL{
		User: url.UserPassword("andresd", "Password01"),
		Scheme: "postgres",
		Host: fmt.Sprintf("%s:%d", "localhost", 5432),
		Path: "postgres",
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn.String(),
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector)
	require.NoError(s.T(), err)

	maxLifetime := 10 * time.Second
	maxIdle, maxOpenConnection := 5, 5

	s.repository = &userRepository{db: &DataSource{s.DB, maxIdle, maxOpenConnection, maxLifetime}}
}

func TestInit(t *testing.T){
	suite.Run(t, new(Suite))
}

var (
	id       = uuid.New()
	email    = "andres@gmail.com"
	password = ".eV1N7PZuuB9Il9eST1HdQQupZ6fzehNK"
)

func (s *Suite) Test_user_repository_Insert(){
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("email","password","id") VALUES ($1,$2,$3) RETURNING "id"`)).WithArgs(email, password, id).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id.String()))
	s.mock.ExpectCommit()

	s.mock.ExpectQuery("SELECT \\* FROM \"users\" WHERE \"users\"\\.\"id\" \\= \\$1").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id.String(), email, password))

	_, err := s.repository.Insert(&models.User{ID: &id, Email: email, Password: password})

	require.NoError(s.T(), err)
}

func (s *Suite) Test_user_repository_Get_User_By_ID(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id.String(), email, password))

	_, err := s.repository.GetUserByID(id.String())

	require.NoError(s.T(), err)
}

func (s *Suite) Test_user_repository_Get_User_By_Email(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id.String(), email, password))

	_, err := s.repository.GetUserByEmail(email)

	require.NoError(s.T(), err)
}