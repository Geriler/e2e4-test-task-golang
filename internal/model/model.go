package model

type Currency struct {
	NumCode  string `db:"num_code"`
	CharCode string `db:"char_code"`
	Nominal  string `db:"nominal"`
	Name     string `db:"name"`
	Value    string `db:"value"`
	Date     string `db:"date"`
}

type Info struct {
	Valute []Currency
}

type Curs struct {
	ValCurs Info
}
