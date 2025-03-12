package domain_test

import (
	"database/sql/driver"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kevin07696/produce-service/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	// Create a new SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	// Configure GORM logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Log output
		logger.Config{
			LogLevel: logger.Info, // Log level (Info, Warn, Error, Silent)
			Colorful: true,        // Enable color
		},
	)

	// Replace GORM's underlying database driver with the mock connection
	gDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: gormLogger, // Enable GORM logger
	})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	return gDB, mock
}

func TestTableMigration(t *testing.T) {
	testCases := []struct {
		Name               string
		ExpectedTableCount int
		ExpectedError      error
	}{
		{Name: "TableAlreadyExists", ExpectedTableCount: 1, ExpectedError: nil},
		{Name: "TableDoesNotExist_CreateTableSucceeds", ExpectedTableCount: 0, ExpectedError: nil},
		{Name: "TableDoesNotExist_CreateTableFails", ExpectedTableCount: 0, ExpectedError: gorm.ErrForeignKeyViolated},
	}

	db, mock := SetupMockDB(t)
	repository := domain.NewProductRepository(db, 1)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock table migration
			mock.ExpectQuery(`SELECT count\(\*\) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA\(\) AND table_name = \$1 AND table_type = \$2`).
				WithArgs("products", sqlmock.AnyArg()).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.ExpectedTableCount))

			if tc.ExpectedTableCount == 0 {
				if tc.ExpectedError == nil {
					mock.ExpectExec(`CREATE TABLE "products"`).
						WillReturnResult(sqlmock.NewResult(0, 0))
					mock.ExpectExec(`CREATE INDEX IF NOT EXISTS "idx_products_deleted_at" ON "products" \("deleted_at"\)`).
						WillReturnResult(sqlmock.NewResult(0, 0))

				} else {
					mock.ExpectExec(`CREATE TABLE "products"`).
						WillReturnError(tc.ExpectedError)
				}
			}

			// Run the migration
			err := repository.Migrate()

			// Verify that the expected migration was executed
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %v", err)
			}

			if tc.ExpectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, tc.ExpectedError, err.Error())
			}
		})
	}
}

func TestProductRepository_ReadProductSummaries(t *testing.T) {
	summary := domain.ProductSummary{ID: uuid.New().String(), Name: "Product 1", ThumbnailUrl: "https://example.com/1.jpg", CategoryName: "Category A", CentPrice: 1000, InStock: true, Rating: 1, AmountSold: 100, CreatedAt: time.Now()}

	testCases := []struct {
		Name           string
		Page           int
		Rows           [][]driver.Value
		Arguments      []driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name: "Page0_ReadSummariesSucceeds",
			Page: 0,
			Rows: [][]driver.Value{
				{summary.ID, summary.Name, summary.ThumbnailUrl, summary.CategoryName, summary.CentPrice, summary.InStock, summary.Rating, summary.AmountSold, summary.CreatedAt},
				{uuid.New(), "Product 2", "https://example.com/2.jpg", "Category B", 2000, false, .80, 50, time.Now()},
			},
			Arguments:      []driver.Value{10},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name: "Page1_ReadSummariesSucceeds",
			Page: 1,
			Rows: [][]driver.Value{
				{summary.ID, summary.Name, summary.ThumbnailUrl, summary.CategoryName, summary.CentPrice, summary.InStock, summary.Rating, summary.AmountSold, summary.CreatedAt},
				{uuid.New(), "Product 2", "https://example.com/2.jpg", "Category B", 2000, false, .80, 50, time.Now()},
			},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Page1_ReadFails_ReturnsNotFound",
			Page:           1,
			Rows:           [][]driver.Value{},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "Page1_ReadFails_ReturnsUnknown",
			Page:           1,
			Rows:           [][]driver.Value{},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternal,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db, 10)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			rows := sqlmock.NewRows([]string{"id", "name", "thumbnail_url", "category_name", "cent_price", "in_stock", "rating", "amount_sold", "created_at"}).
				AddRows(tc.Rows...)

			mock.ExpectQuery(`SELECT .* FROM "products" WHERE "products"."deleted_at" IS NULL LIMIT .*`).
				WithArgs(tc.Arguments...).WillReturnRows(rows)

			// Run the query for product summaries.
			summaries, status := repo.ReadProductSummaries(tc.Page)

			// Verify that the query was executed
			assert.Equal(t, domain.StatusOK, status, "Status should be OK")
			assert.Equal(t, len(tc.Rows), len(summaries), "Expected 2 product summaries")
			if len(tc.Rows) > 0 {
				assert.Equal(t, summary, summaries[0], "Product Summary should match")
			}

			// Ensure all expectations were met.
			assert.NoError(t, mock.ExpectationsWereMet(), "All expectations should be met")
		})
	}

}

func TestProductRepository_ReadProductDetail(t *testing.T) {
	detail := domain.ProductDetail{
		ID:           uuid.NewString(),
		Name:         "Product 1",
		CategoryName: "Category A",
		Description:  "Product Description",
		CentPrice:    10000,
		Rating:       0.85,
		Attributes:   datatypes.JSON(`{ "material": "100% Cotton", "care_instructions": "Machine wash cold. Tumble dry low.", "size_chart": { "xs": { "bust": "30-32 inches", "waist": "24-26 inches", "hips": "33-35 inches" } } } `),
		Options:      datatypes.JSON(`{ "options": [ { "name": "size", "options": ["xs", "s", "m", "lg"] } ] } `),
		MainOption:   datatypes.JSON(`{ "name": "pattern", "options": [ { "value": "Striped", "image_urls": ["https://example.com/striped-pattern1.jpg", "https://example.com/striped-pattern2.jpg"] } ] }`),
	}

	testCases := []struct {
		Name           string
		ProductID      datatypes.UUID
		Row            []driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Succeeds",
			ProductID:      datatypes.UUID(uuid.MustParse(detail.ID)),
			Row:            []driver.Value{detail.ID, detail.Name, detail.CategoryName, detail.Description, detail.CentPrice, detail.Rating, detail.Attributes, detail.Options, detail.MainOption, detail.CreatedAt},
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Fails_ReturnsNotFound",
			ProductID:      datatypes.UUID(uuid.New()),
			Row:            []driver.Value{},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "Fails_ReturnsUnknown",
			ProductID:      datatypes.UUID(uuid.New()),
			Row:            []driver.Value{},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternal,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db, 10)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			query := mock.ExpectQuery(`SELECT .* FROM "products" WHERE id = \$1 LIMIT \$2`).
				WithArgs(tc.ProductID, 1)

			if tc.ErrorMock == nil {
				rows := mock.NewRows([]string{"id", "name", "category_name", "description", "cent_price", "rating", "attributes", "options", "main_option", "created_at"}).
					AddRow(tc.Row...)
				query.WillReturnRows(rows)
			} else {
				query.WillReturnError(tc.ErrorMock)
			}

			// Run the query for product response.
			response, status := repo.ReadProductDetail(tc.ProductID)

			// Verify that the query was executed
			if tc.ErrorMock == nil {
				assert.Equal(t, detail, response)
			}
			assert.Equal(t, tc.ExpectedStatus, status)

			// Ensure all expectations were met.
			assert.NoError(t, mock.ExpectationsWereMet(), "All expectations should be met")
		})
	}
}

func TestProductRepository_ReadCategories(t *testing.T) {
	testCases := []struct {
		Name           string
		Page           int
		Rows           [][]driver.Value
		Arguments      []driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Page0_ReadCategoriesSucceeds",
			Page:           0,
			Rows:           [][]driver.Value{{"Category A"}, {"Category B"}, {"Category C"}},
			Arguments:      []driver.Value{10},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Page1_ReadCategoriesSucceeds",
			Page:           1,
			Rows:           [][]driver.Value{{"Category A"}, {"Category B"}, {"Category C"}},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Page1_ReadFails_ReturnsNotFound",
			Page:           1,
			Rows:           [][]driver.Value{},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "Page1_ReadFails_ReturnsUnknown",
			Page:           1,
			Rows:           [][]driver.Value{},
			Arguments:      []driver.Value{10, 10},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternal,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db, 10)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			rows := sqlmock.NewRows([]string{"category_name"}).
				AddRows(tc.Rows...)

			mock.ExpectQuery(`SELECT DISTINCT "category_name" FROM "products" LIMIT .*`).
				WithArgs(tc.Arguments...).WillReturnRows(rows)

			// Run the query for product categories.
			categories, status := repo.ReadCategories(tc.Page)

			// Verify that the query was executed
			assert.Equal(t, domain.StatusOK, status, "Status should be OK")
			assert.Equal(t, len(tc.Rows), len(categories), "Expected 2 product summaries")

			// Ensure all expectations were met.
			assert.NoError(t, mock.ExpectationsWereMet(), "All expectations should be met")
		})
	}
}

