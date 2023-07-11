package main

import (
	"context"
	"e2e4-test-task/internal/model"
	er "e2e4-test-task/internal/services/exchange-rates"
	"e2e4-test-task/internal/storage"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func handleExchangeRates(c echo.Context) error {
	var (
		db   *storage.CurrencyStorage
		date time.Time
		err  error
		res  []model.Currency
	)

	dParam := c.QueryParam("date")
	if dParam == "" {
		date = time.Now()
	} else {
		if date, err = time.Parse("02.01.2006", dParam); err != nil {
			return c.JSON(http.StatusBadRequest, "incorrect date")
		}
	}

	db, err = storage.Init(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	res, err = er.GetByDate(date, db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, res)
}
