package queries

import (
	"booking-car/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestTransactionQuery_BrowseTypeAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,` +
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 ` +
		`ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", "", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseTypeOnGoing(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `on_going`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,` +
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 ` +
		`AND t.transaction_type=$4 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0, transactionType).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseTypeSuccess(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `success`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,` +
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 ` +
		`AND t.transaction_type=$4 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0, transactionType).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseTypeCanceled(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `canceled`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,` +
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 ` +
		`AND t.transaction_type=$4 ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0, transactionType).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseWithUserIdTypeAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,`+
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL `+
		`WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.user_id=$2 ORDER BY created_at desc LIMIT $3 OFFSET $4`
	mock.ExpectQuery(statement).WithArgs(search, userId,10,0).WillReturnRows(rows)
	res, err := repository.BrowseByUserId("", "created_at", "desc", userId, "", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseWithUserIdTypeOnGoing(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `on_going`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,`+
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL `+
		`WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.user_id=$2 AND t.transaction_type=$5 ORDER BY created_at desc LIMIT $3 OFFSET $4`
	mock.ExpectQuery(statement).WithArgs(search, userId,10,0,transactionType).WillReturnRows(rows)
	res, err := repository.BrowseByUserId("", "created_at", "desc", userId, transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseWithUserIdTypeSuccess(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `success`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,`+
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL `+
		`WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.user_id=$2 AND t.transaction_type=$5 ORDER BY created_at desc LIMIT $3 OFFSET $4`
	mock.ExpectQuery(statement).WithArgs(search, userId,10,0,transactionType).WillReturnRows(rows)
	res, err := repository.BrowseByUserId("", "created_at", "desc", userId, transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_BrowseWithUserIdTypeOnCanceled(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	transactionType := `canceled`
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,`+
		`u.last_name,u.email,u.phone_number FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL `+
		`WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.user_id=$2 AND t.transaction_type=$5 ORDER BY created_at desc LIMIT $3 OFFSET $4`
	mock.ExpectQuery(statement).WithArgs(search, userId,10,0,transactionType).WillReturnRows(rows)
	res, err := repository.BrowseByUserId("", "created_at", "desc", userId, transactionType, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestNewTransactionQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	userId := gofakeit.UUID()
	model := models.NewTransactionModel().SetId(gofakeit.UUID()).SetUserId(userId).SetTransactionNumber(strconv.Itoa(gofakeit.Number(1, 10))).
		SetTotalAmount(gofakeit.Price(100000000, 130000000)).SetPaymentReceived(gofakeit.Price(100000000, 130000000)).
		SetCreatedAt(now).SetUpdatedAt(now).SetPaidAt(now).SetCanceledAt(now)
	model.User = models.NewUserModel().SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPhoneNumber(gofakeit.Phone())
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"t.id", "t.user_id", "t.transaction_type", "t.transaction_number", "t.total_amount", "t.payment_received",
		"t.created_at", "t.updated_at", "t.paid_at", "t.canceled_at", "u.first_name", "u.last_name", "u.email", "u.phone_number", "array_to_string"}).AddRow(model.Id(), model.UserId(), model.TransactionType(),
		model.TransactionNumber(), model.TotalAmount(), model.PaymentReceived(), model.CreatedAt(), model.UpdatedAt(), model.PaidAt(), model.CanceledAt(), model.User.FirstName(),
		model.User.LastName(), model.User.Email(), model.User.PhoneNumber(), model.TransactionDetail())
	statement := `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,u.first_name,u.last_name,u.email,u.phone_number ,ARRAY_TO_STRING(ARRAY_AGG(td.id ||':'|| td.car_id ||':'|| td.car_brand ||':'|| td.car_type ||':'|| td.car_color ||':'|| td.production_year ||':'|| td.price ||':'|| td.quantity ||':'|| td.sub_total),',') FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL INNER JOIN transaction_details td ON td.transaction_id = t.id AND td.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.id=$1 GROUP BY t.id,u.id`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)
	res, err := repository.ReadBy("t.id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountWithUserIdTypeAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	userId := gofakeit.UUID()
	transactionType := ""
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.user_id='` + userId + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", userId, transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeWithUserIdOnSuccess(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	userId := gofakeit.UUID()
	transactionType := "success"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `' AND t.user_id='` + userId + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", userId, transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeWithUserIdOnGoing(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	userId := gofakeit.UUID()
	transactionType := "on_going"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `' AND t.user_id='` + userId + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", userId, transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountWithUserIdTypeCancel(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	userId := gofakeit.UUID()
	transactionType := "canceled"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `' AND t.user_id='` + userId + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", userId, transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	transactionType := ""
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", "", transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeOnGoing(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	transactionType := "on_going"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", "", transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeSuccess(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	transactionType := "success"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", "", transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountTypeCanceled(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	transactionType := "canceled"
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL WHERE t.deleted_at IS NULL AND t.transaction_number LIKE $1 AND t.transaction_type='` + transactionType + `'`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("", "", transactionType)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestTransactionQuery_CountAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	repository := TransactionQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(t.id) FROM transactions t`
	mock.ExpectQuery(statement).WithArgs().WillReturnRows(rows)
	res, err := repository.CountAll()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
