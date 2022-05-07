package repositories

import (
	"belanjayukid_go/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"regexp"
)

var (
	categoryID = uuid.New()
	categoryName = "makanan"
)

func (s *Suite) Test_category_repository_Insert(){
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "categories" ("name","id") VALUES ($1,$2) RETURNING "id"`)).WithArgs(categoryName, categoryID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(categoryID.String()))
	s.mock.ExpectCommit()

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE "categories"."id" = $1`)).WithArgs(categoryID).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(categoryID.String(), categoryName))

	_, err := s.categoryRepository.Insert(&models.Category{ID: &categoryID, Name: categoryName})
	require.NoError(s.T(), err)
}

func (s *Suite) Test_category_repository_Get_Category_List(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(categoryID.String(), categoryName))

	_, err := s.categoryRepository.GetCategoryList()
	require.NoError(s.T(), err)
}
