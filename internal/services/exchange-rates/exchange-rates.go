package exchange_rates

import (
	"context"
	"e2e4-test-task/internal/adapters"
	"e2e4-test-task/internal/model"
	"e2e4-test-task/internal/storage"
	"encoding/json"
	g "gopkg.in/h2non/gentleman.v2"
	"time"
)

func GetByDate(date time.Time, db *storage.CurrencyStorage) ([]model.Currency, error) {
	var (
		j        string
		err      error
		response *g.Response
		result   []model.Currency
	)
	result, err = db.Get(context.Background(), date)

	if len(result) == 0 {
		response, err = g.New().URL("https://www.cbr.ru/scripts/XML_daily.asp").
			Request().
			SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8").
			SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11").
			SetQuery("date", date.Format("02.01.2006")).
			Send()

		if err != nil {
			return []model.Currency{}, err
		}

		j, err = adapters.Xml2json(response.String())
		if err != nil {
			return []model.Currency{}, err
		}

		var data model.Curs
		err = json.Unmarshal([]byte(j), &data)
		if err != nil {
			return []model.Currency{}, err
		}

		for i := 0; i < len(data.ValCurs.Valute); i++ {
			currency := &data.ValCurs.Valute[i]
			currency.Date = date.Format("02.01.2006")
			err := db.Add(context.Background(), *currency, date)
			if err != nil {
				return nil, err
			}
		}
		result = data.ValCurs.Valute
	}

	return result, nil
}
