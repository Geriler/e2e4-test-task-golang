package storage

import (
	"context"
	"e2e4-test-task/internal/model"
	"e2e4-test-task/internal/storage/migration"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CurrencyStorage struct {
	db *sqlx.DB
}

var cs *CurrencyStorage

func Init(ctx context.Context) (*CurrencyStorage, error) {
	if cs == nil {
		db, err := sqlx.ConnectContext(ctx, "sqlite3", "db.sqlite3")
		if err != nil {
			return nil, err
		}
		cs = &CurrencyStorage{db: db}

		err = migration.Init(cs.db.DB)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func (s *CurrencyStorage) Get(ctx context.Context, date time.Time) ([]model.Currency, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var currencies []model.Currency
	err = conn.SelectContext(ctx, &currencies, "SELECT * FROM currencies WHERE date = $1", date.Format("02.01.2006"))
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func (s *CurrencyStorage) Add(ctx context.Context, model model.Currency, date time.Time) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecContext(
		ctx,
		"INSERT INTO currencies (num_code, char_code, nominal, name, value, date) VALUES ($1, $2, $3, $4, $5, $6)",
		model.NumCode,
		model.CharCode,
		model.Nominal,
		model.Name,
		model.Value,
		date.Format("02.01.2006"),
	)
	if err != nil {
		return err
	}

	return nil
}
