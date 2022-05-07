package repositories

import (
	"belanjayukid_go/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"regexp"
)

var (
	productUnitID = uuid.New()
	productUnitName = "lusin"
)

func (s *Suite) Test_product_unit_repository_Insert(){
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "product_units" ("name","id") VALUES ($1,$2) RETURNING "id"`)).WithArgs(productUnitName, productUnitID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(productUnitID))
	s.mock.ExpectCommit()

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_units" WHERE "product_units"."id" = $1`)).WithArgs(productUnitID).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(productUnitID.String(), productUnitName))

	_, err := s.productUnitRepository.Insert(&models.ProductUnit{ID: &productUnitID, Name: productUnitName})
	require.NoError(s.T(), err)
}

func (s *Suite) Test_product_unit_repository_Get_Product_Unit_List(){
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_units"`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(productUnitID.String(), productUnitName))

	_, err := s.productUnitRepository.GetProductUnitList()
	require.NoError(s.T(), err)
}
