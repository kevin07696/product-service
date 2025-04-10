package domain_test

import (
	"context"
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
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: gormLogger, // Enable GORM logger
	})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	return gormDB, mock
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
	repository := domain.NewProductRepository(db)

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
		Rows           [][]driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name: "ReadSummariesSucceeds",
			Rows: [][]driver.Value{
				{summary.ID, summary.Name, summary.ThumbnailUrl, summary.CategoryName, summary.CentPrice, summary.InStock, summary.Rating, summary.AmountSold, summary.CreatedAt},
				{uuid.New(), "Product 2", "https://example.com/2.jpg", "Category B", 2000, false, .80, 50, time.Now()},
			},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "ReadFails_ReturnsNotFound",
			Rows:           [][]driver.Value{},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "ReadFails_ReturnsUnknown",
			Rows:           [][]driver.Value{},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternalError,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			query := mock.ExpectQuery(`SELECT .* FROM "products" WHERE category_name <> \$1`).WithArgs("Archived")

			if tc.ErrorMock == nil {
				rows := sqlmock.NewRows([]string{"id", "name", "thumbnail_url", "category_name", "cent_price", "in_stock", "rating", "amount_sold", "created_at"}).
					AddRows(tc.Rows...)
				query.WillReturnRows(rows)
			} else {
				query.WillReturnError(tc.ErrorMock)
			}
			// Run the query for product summaries.
			summaries, status := repo.ReadProductSummaries(context.TODO())

			// Verify that the query was executed
			assert.Equal(t, tc.ExpectedStatus, status)
			assert.Equal(t, len(tc.Rows), len(summaries))
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
		ProductID      uuid.UUID
		Row            []driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Succeeds",
			ProductID:      uuid.MustParse(detail.ID),
			Row:            []driver.Value{detail.ID, detail.Name, detail.CategoryName, detail.Description, detail.CentPrice, detail.Rating, detail.Attributes, detail.Options, detail.MainOption, detail.CreatedAt},
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Fails_ReturnsNotFound",
			ProductID:      uuid.New(),
			Row:            []driver.Value{},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "Fails_ReturnsUnknown",
			ProductID:      uuid.New(),
			Row:            []driver.Value{},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternalError,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			query := mock.ExpectQuery(`SELECT .* FROM "products" WHERE id = \$1 AND category_name <> \$2 LIMIT \$3`).
				WithArgs(tc.ProductID, "Archived", 1)

			if tc.ErrorMock == nil {
				rows := mock.NewRows([]string{"id", "name", "category_name", "description", "cent_price", "rating", "attributes", "options", "main_option", "created_at"}).
					AddRow(tc.Row...)
				query.WillReturnRows(rows)
			} else {
				query.WillReturnError(tc.ErrorMock)
			}

			// Run the query for product response.
			response, status := repo.ReadProductDetail(context.TODO(), tc.ProductID)

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
		Rows           [][]driver.Value
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Succeeds",
			Rows:           [][]driver.Value{{"Category A"}, {"Category B"}, {"Category C"}},
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Fails_ReturnsNotFound",
			Rows:           [][]driver.Value{},
			ErrorMock:      gorm.ErrRecordNotFound,
			ExpectedStatus: domain.StatusNotFound,
		},
		{
			Name:           "Fails_ReturnsUnknown",
			Rows:           [][]driver.Value{},
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternalError,
		},
	}

	db, mock := SetupMockDB(t)
	repo := domain.NewProductRepository(db)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Mock the query for product summaries.
			query := mock.ExpectQuery(`SELECT DISTINCT category_name FROM "products" WHERE category_name <> \$1`).WithArgs("Archived")

			if tc.ErrorMock == nil {
				rows := sqlmock.NewRows([]string{"category_name"}).
					AddRows(tc.Rows...)
				query.WillReturnRows(rows)
			} else {
				query.WillReturnError(tc.ErrorMock)
			}

			// Run the query for product categories.
			categories, status := repo.ReadCategories(context.TODO())

			// Verify that the query was executed
			assert.Equal(t, tc.ExpectedStatus, status)
			assert.Equal(t, len(tc.Rows), len(categories))

			// Ensure all expectations were met.
			assert.NoError(t, mock.ExpectationsWereMet(), "All expectations should be met")
		})
	}
}

func TestWriteProduct(t *testing.T) {
	product := domain.Product{
		Model: domain.Model{
			ID: (uuid.New()),
		},
		Name:          "Classic Jeans",
		ThumbnailUrl:  "https://my-bucket.s3.us-east-1.amazonaws.com/images/classic-jeans.jpg",
		Description:   "These are really nice jeans.",
		CentPrice:     5000,
		InStock:       true,
		Rating:        0.75,
		AmountSold:    6,
		GalleryOption: datatypes.JSON(`{ name: pattern, options: [ { value: Striped, image_urls: [https://example.com/striped-pattern1.jpg, https://example.com/striped-pattern2.jpg] } ] }`),
		Options:       datatypes.JSON(`{ options: [ { name: size, options: [xs, s, m, lg] } ] }`),
		Attributes:    datatypes.JSON(`{ attributes: { material: 100% Cotton, care_instructions: Machine wash cold. Tumble dry low., size_chart: { xs: { bust: 30-32 inches, waist: 24-26 inches, hips: 33-35 inches } } } }`),
		CategoryName:  "pants",
	}

	testCases := []struct {
		Name         string
		ErrorMock    error
		ExpectStatus domain.StatusCode
	}{
		{
			Name:         "Success_ReturnsStatusOK",
			ErrorMock:    nil,
			ExpectStatus: domain.StatusOK,
		},
		{
			Name:         "Fails_ReturnsDuplicateKey",
			ErrorMock:    gorm.ErrDuplicatedKey,
			ExpectStatus: domain.StatusAlreadyExists,
		},
		{
			Name:         "FailsUnexpectedly_ReturnsInternal",
			ErrorMock:    errors.New("Unexpected Error"),
			ExpectStatus: domain.StatusInternalError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// setup
			db, mock := SetupMockDB(t)

			// mock query
			mock.ExpectBegin()

			// Fix: Use sqlmock.AnyArg() for the ID parameter to match any UUID
			query := mock.ExpectExec(`INSERT INTO "products" \("id","created_at","updated_at","deleted_at","name","thumbnail_url","category_name","description","cent_price","amount_sold","in_stock","rating","options","main_option","attributes"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9,\$10,\$11,\$12,\$13,\$14,\$15\)`).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), product.Name, product.ThumbnailUrl, product.CategoryName, product.Description, product.CentPrice, product.AmountSold, product.InStock, product.Rating, product.Options, product.GalleryOption, product.Attributes)

			if tc.ErrorMock == nil {
				query.WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				query.WillReturnError(tc.ErrorMock)
				mock.ExpectRollback()
			}

			// call function
			repository := domain.NewProductRepository(db)
			status := repository.WriteProduct(context.TODO(), &product)

			// assert function expectations
			assert.Equal(t, tc.ExpectStatus, status)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	product := domain.Product{
		Model: domain.Model{
			ID: (uuid.New()),
		},
		Name:       "Product 1",
		Attributes: datatypes.JSON(`{ attributes: { material: 100% Cotton, care_instructions: Machine wash cold. Tumble dry low., size_chart: { xs: { bust: 30-32 inches, waist: 24-26 inches, hips: 33-35 inches } } } }`),
	}

	testCases := []struct {
		Name              string
		ProductId         string
		ProductUpdates    domain.Product
		ErrorMock         error
		ExpectedExecution string
		ExpectedArguments []driver.Value
		ExpectedStatus    domain.StatusCode
	}{
		{
			Name:              "Success_ReturnStatusOK",
			ProductId:         "fbb6e4ab-8569-4c7f-bd8d-d754bf44bb30",
			ProductUpdates:    domain.Product{Attributes: product.Attributes},
			ExpectedExecution: `UPDATE "products" SET "updated_at"=\$1,"attributes"=\$2 WHERE id = \$3 AND "products"."deleted_at" IS NULL`,
			ExpectedArguments: []driver.Value{sqlmock.AnyArg(), product.Attributes, "fbb6e4ab-8569-4c7f-bd8d-d754bf44bb30"},
			ErrorMock:         nil,
			ExpectedStatus:    domain.StatusOK,
		},
		{
			Name:              "Fails_ReturnDuplicateKey",
			ProductId:         "e838ab0e-d398-42a1-8acf-f3b8f6330812",
			ProductUpdates:    domain.Product{Name: product.Name},
			ExpectedExecution: `UPDATE "products" SET "updated_at"=\$1,"name"=\$2 WHERE id = \$3 AND "products"."deleted_at" IS NULL`,
			ExpectedArguments: []driver.Value{sqlmock.AnyArg(), product.Name, "e838ab0e-d398-42a1-8acf-f3b8f6330812"},
			ErrorMock:         gorm.ErrDuplicatedKey,
			ExpectedStatus:    domain.StatusAlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			db, mock := SetupMockDB(t)

			mock.ExpectBegin()

			exec := mock.ExpectExec(tc.ExpectedExecution).WithArgs(tc.ExpectedArguments...)
			if tc.ErrorMock == nil {
				exec.WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			} else {
				exec.WillReturnError(tc.ErrorMock)
				mock.ExpectRollback()
			}

			repository := domain.NewProductRepository(db)
			status := repository.UpdateProduct(context.TODO(), uuid.MustParse(tc.ProductId), tc.ProductUpdates)

			assert.NoError(t, mock.ExpectationsWereMet())
			assert.Equal(t, tc.ExpectedStatus, status)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	testCases := []struct {
		Name           string
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Success",
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Fail_ReturnsInternal",
			ErrorMock:      errors.New(""),
			ExpectedStatus: domain.StatusInternalError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			db, mock := SetupMockDB(t)

			productId := (uuid.New())

			mock.ExpectBegin()

			exec := mock.ExpectExec(`UPDATE "products" SET "deleted_at"=\$1 WHERE id = \$2 AND "products"."deleted_at" IS NULL`).
				WithArgs(sqlmock.AnyArg(), productId)

			if tc.ErrorMock == nil {
				exec.WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				exec.WillReturnError(tc.ErrorMock)
				mock.ExpectRollback()
			}

			repository := domain.NewProductRepository(db)
			status := repository.DeleteProduct(context.TODO(), productId)

			assert.Equal(t, tc.ExpectedStatus, status)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRecoverProduct(t *testing.T) {
	testCases := []struct {
		Name           string
		ProductId      string
		ErrorMock      error
		ExpectedStatus domain.StatusCode
	}{
		{
			Name:           "Success_ReturnStatusOK",
			ProductId:      "354c6abe-89c4-46dd-b51c-d1b9c20dac4b",
			ErrorMock:      nil,
			ExpectedStatus: domain.StatusOK,
		},
		{
			Name:           "Fails_ReturnUnknown",
			ProductId:      "354c6abe-89c4-46dd-b51c-d1b9c20dac4b",
			ErrorMock:      errors.New("Unknown Error"),
			ExpectedStatus: domain.StatusInternalError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			db, mock := SetupMockDB(t)

			mock.ExpectBegin()

			exec := mock.ExpectExec(`UPDATE "products" SET "deleted_at"=\$1 WHERE id = \$2`).
				WithArgs(nil, "354c6abe-89c4-46dd-b51c-d1b9c20dac4b")
			if tc.ErrorMock == nil {
				exec.WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			} else {
				exec.WillReturnError(tc.ErrorMock)
				mock.ExpectRollback()
			}

			repository := domain.NewProductRepository(db)
			status := repository.RecoverProduct(context.TODO(), (uuid.MustParse("354c6abe-89c4-46dd-b51c-d1b9c20dac4b")))

			assert.NoError(t, mock.ExpectationsWereMet())
			assert.Equal(t, tc.ExpectedStatus, status)
		})
	}
}
