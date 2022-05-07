package repositories

import (
	"belanjayukid_go/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"regexp"
)

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

	_, err := s.userRepository.Insert(&models.User{ID: &id, Email: email, Password: password})

	require.NoError(s.T(), err)
}

func (s *Suite) Test_user_repository_Get_User_By_ID(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.id = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id.String(), email, password))

	_, err := s.userRepository.GetUserByID(id.String())

	require.NoError(s.T(), err)
}

func (s *Suite) Test_user_repository_Get_User_By_Email(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id.String(), email, password))

	_, err := s.userRepository.GetUserByEmail(email)

	require.NoError(s.T(), err)
}